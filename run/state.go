package run

import (
	"arit/modules"
	"fmt"
	"os"
	"strings"
)

const LS_FORMAT = "\033[1G%-20s %-10s %s\n"

var LS_HEADER_FORMAT = "\033[1m" + LS_FORMAT + "\033[0m"

type State struct {
	Vars    map[string]string
	Modules map[string]*modules.Submodule
	Aliases map[string]string
	History [][]byte
	current []byte
}

func PWD() error {
	fmt.Println(os.Getenv("PWD"))
	return nil
}

func (s *State) PrintHelp(args []string) error {
	if len(args) > 1 {
		return fmt.Errorf("help command accepts 1 argument")
	}

	if len(args) == 0 {
		fmt.Printf(LS_HEADER_FORMAT, "Commands", "Module", "Help/description")
		fmt.Printf(LS_FORMAT, "store/set", "Std", "Add value to storage")
		fmt.Printf(LS_FORMAT, "get", "Std", "Retrieves values from storage")
		fmt.Printf(LS_FORMAT, "export", "Std", "Prints all values in storage")
		for _, sub := range s.Modules {
			fmt.Printf(LS_FORMAT, strings.Join(sub.Keys, "/"), sub.Name, sub.Help)
		}
		return nil
	}

	name, ok := s.Aliases[args[0]]
	if !ok {
		return fmt.Errorf("help %s: unknown command", args[0])
	}

	return s.Modules[name].PrintHelp()
}

func (s *State) Store(args []string) error {
	switch len(args) {
	case 0, 1:
		return fmt.Errorf("store command expect ident and value args")
	case 2:
		s.Vars[args[0]] = args[1]
		return nil
	default:
		return fmt.Errorf("unknown args passed to get command: %v", args)
	}
}

func (s *State) Get(args []string) (string, error) {
	if len(args) > 1 {
		return "", fmt.Errorf("get: unknown args passed to command: %v", args)
	}

	if len(args) == 0 {
		return "", fmt.Errorf("get: expected 1 argument, got 0")
	}

	v, ok := s.Vars[args[0]]
	if !ok {
		return "", fmt.Errorf("get: unknown var %s", args[0])
	}
	return v, nil
}

func (s *State) std(args []string) (ok bool, err error) {
	if len(args) == 0 {
		return false, fmt.Errorf("cannot parse empty list of args")
	}

	switch args[0] {
	case "pwd":
		return true, PWD()

	case "help":
		return true, s.PrintHelp(args[1:])

	case "store", "set":
		return true, s.Store(args[1:])

	case "get":
		val, err := s.Get(args[1:])
		if err == nil {
			fmt.Println(val)
		}
		return true, err

	case "export":
		if args[0] == "export" {
			for k, v := range s.Vars {
				fmt.Printf("%s=%s\n", k, v)
			}
			return true, nil
		}

	default:
		return false, nil
	}

	return false, nil
}

func (s *State) ParseRaw(args string) error {
	return s.Parse(strings.Split(strings.Trim(args, " "), " "))
}

func (s *State) Parse(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("cannot parse empty list of args")
	}

	used, err := s.std(args)
	if used {
		return err
	}

	name, ok := s.Aliases[args[0]]
	if !ok {
		v, err := s.Get(args)
		if err != nil {
			return fmt.Errorf("%s: unknown command", args[0])
		}

		fmt.Println(v)
		return nil
	}

	sub, ok := s.Modules[name]
	if !ok {
		panic(fmt.Sprintf("alias %s does not point to moduel", name))
	}

	value, err := sub.Run(args[1:])
	if err != nil {
		return err
	}

	fmt.Println(value)

	return nil
}

func (s *State) EchoStored(width int) {
	for k := range s.Vars {
		fmt.Println(s.ToString(k, width))
	}
}

func (s *State) ToString(ident string, w int) string {
	x, ok := s.Vars[ident]
	if !ok {
		panic(fmt.Errorf("trying to access unknown var %s in state", ident))
	}

	str := fmt.Sprintf("%s: %v", ident, x)

	if len(str) > w {
		return str[:w-2] + ".."
	}

	return fmt.Sprintf("%s: %v", ident, x)
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
