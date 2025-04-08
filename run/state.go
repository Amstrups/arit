package run

import (
	"arit/modules"
	"arit/run/cursor"
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"
)

type (
	StdFunction struct {
		key  string
		help string
		f    func(*State, ...string) (any, error)
	}

	State struct {
		Vars    map[string]string
		Modules map[string]*modules.Submodule
		Aliases map[string]string
		History [][]byte
		current []byte
		*cursor.Cursor
	}
)

var std map[string]*StdFunction

const stdModuleName = "Std"

func init() {
	get := &StdFunction{
		key:  "get",
		help: "Retrieve value(s) from storage",
		f: func(ste *State, args ...string) (any, error) {
			return ste.Get(args)
		},
	}

	set := &StdFunction{
		key:  "set",
		help: "Add value(s) to storage",
		f: func(ste *State, args ...string) (any, error) {
			return nil, ste.Store(args)
		},
	}

	export := &StdFunction{
		key:  "export",
		help: "Prints all values in storage",
		f: func(ste *State, args ...string) (any, error) {
			for k, v := range ste.Vars {
				fmt.Printf("%s=%s\n", k, v)
			}
			return nil, nil
		},
	}

	helpCmd := &StdFunction{
		key:  "help",
		help: "Prints this",
		f: func(ste *State, args ...string) (any, error) {
			return nil, ste.PrintHelp(args)
		},
	}

	pwd := &StdFunction{
		key:  "pwd",
		help: "Get current working directory",
		f: func(*State, ...string) (any, error) {
			return os.Getenv("PWD"), nil
		},
	}

	moon := &StdFunction{
		key:  "moon",
		help: "Moon",
		f: func(*State, ...string) (any, error) {
			moon, err := os.ReadFile("bin/moon")
			return string(moon), err
		},
	}

	std = map[string]*StdFunction{
		helpCmd.key: helpCmd,
		set.key:     set,
		get.key:     get,
		export.key:  export,
		pwd.key:     pwd,
		moon.key:    moon,
	}
}

func (s *State) getSubmodule(arg string) (*modules.Submodule, error) {
	name, ok := s.Aliases[arg]
	if !ok {
		return nil, fmt.Errorf("help %s: unknown command", arg)
	}

	sub, ok := s.Modules[name]
	if !ok {
		return nil, fmt.Errorf("help %s: unknown command", arg)
	}

	return sub, nil
}

const column_whitespace = 3

func (s *State) PrintHelp(args []string) error {
	if len(args) > 1 {
		return fmt.Errorf("help command accepts {0,1} argument(s)")
	}

	if len(args) == 0 { // TODO: Fix help menu for std
		s.InsertAtNewline2(cursor.F_MODULE_HELP_HEADER, "Commands", "Module", "Help/description")

		for _, f := range std {
			s.InsertAtNewline2(cursor.F_MODULE_HELP, f.key, stdModuleName, f.help)
		}

		for _, sub := range s.Modules {
			s.InsertAtNewline2(cursor.F_MODULE_HELP, strings.Join(sub.Keys, "/"), sub.Name, sub.Help)
		}

		s.Render()
		return nil
	}

	sub, err := s.getSubmodule(args[0])
	if err != nil {
		return err
	}

	var n_len, c_len = cursor.FUNCTION_BYTES_LEN, cursor.COMMAND_BYTES_LEN

	for k, v := range sub.Funcs {
		n_len = max(n_len, len(v.Name))
		c_len = max(c_len, len(k))
	}

	size := n_len + c_len + 2*column_whitespace
	clean := bytes.Repeat([]byte(" "), size)
	b := make([]byte, size)
	copy(b, clean)
	copy(b, []byte("Function"))
	copy(b[n_len+column_whitespace:], []byte("Command"))
	fmt.Printf("\033[1G\033[1m%s%s\033[0m\n", b, "Description")

	for k, v := range sub.Funcs {
		copy(b, clean)
		copy(b, []byte(v.Name))
		copy(b[n_len+column_whitespace:], []byte(k))
		fmt.Printf("\033[1G%s", b)
		fmt.Print(v.Help + "\n\033[1G")
	}

	return nil
}

var unclosedQuoteError = errors.New("unclosed quote")

func parse(args string) ([][]string, error) {
	commands := [][]string{}
	arg := []string{}
	val := ""

	in_quote := false
	ps := false

	for _, x := range args {
		if in_quote {
			if x == '"' {
				in_quote = false
				arg = append(arg, val)
				val = ""
				continue
			}

			val += string(x)
			continue
		}

		if x == '|' {
			ps = true
			continue
		} else if x == '>' && ps {
			_arg := make([]string, len(arg))
			copy(_arg, arg)

			commands = append(commands, _arg)

			arg = []string{}
			val = ""
			ps = false

			continue
		}

		switch x {
		case cursor.SPACE:
			if val != "" {
				arg = append(arg, val)
			}
			val = ""
			continue
		case '"':
			in_quote = true
		default:
			val += string(x)
		}

	}
	if in_quote {
		return commands, unclosedQuoteError

	}

	if val != "" {
		arg = append(arg, val)
	}
	commands = append(commands, arg)

	return commands, nil
}

func (s *State) ParseRaw(args string) error {
	commands, err := parse(args)

	var msg string = "args: "

	for _, e := range commands {
		msg += fmt.Sprintf("%v", e)
	}

	if err != nil {
		return err
	}

	return JoinErrors(s.Parse(commands...), errors.New(msg))
}

func (s *State) Parse(commands ...[]string) error {
	var piped []any
	for _, args := range commands {
		if len(piped) > 0 {
			str_piped := make([]string, len(piped))
			for i, p := range piped {
				str_piped[i] = fmt.Sprintf("%v", p)
			}
			args = append(args, str_piped...)
		}

		if len(args) == 0 {
			return fmt.Errorf("cannot parse empty list of args")
		}

		// check for std lib
		stdFunc, ok := std[args[0]]
		if ok {
			val, err := stdFunc.f(s, args[1:]...)
			if err != nil {
				return err
			}

			switch t := val.(type) {
			case error:
				return t
			case []any:
				piped = t
			default:
				piped = []any{t}
			}
			continue
		}

		// check for modules
		name, ok := s.Aliases[args[0]]
		if !ok {
			v, err := s.Get(args)
			if err != nil {
				return fmt.Errorf("%s: unknown command", args[0])
			}

			piped = []any{v}
			continue
		}

		sub, ok := s.Modules[name]
		if !ok {
			panic(fmt.Sprintf("alias %s does not point to moduel", name))
		}

		value, err := sub.Run(args[1:])
		if err != nil {
			return err
		}

		if piped, ok = value.([]any); !ok {
			piped = []any{value}
		}
		value = piped
	}

	for _, v := range piped {
		if v != nil {
			s.InsertAnyAtNewline(v)
		}
	}

	return nil
}

func (s *State) insert(x, ln int, bytes ...byte) (n int) {
	inserts := len(bytes)

	if ln+inserts > len(s.current) {
		magnitude := ((len(s.current)+inserts)/128 + 1)
		_current := make([]byte, 128*magnitude)
		copy(_current, s.current)
		s.current = _current
	}
	if x < ln {
		copy(s.current[x+inserts:ln+inserts], s.current[x:ln])
	}
	copy(s.current[x:], bytes)

	return inserts
}

func (s *State) remove(x, ln, k int) (n int) {
	if ln == 0 || x == 0 || k == 0 {
		return
	}

	if k > x {
		k = x
	}

	if x < ln {
		copy(s.current[x-k:ln-k], s.current[x:ln])
	}

	copy(s.current[ln-(k-1):], make([]byte, k))

	return k
}
