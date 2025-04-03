package modules

import (
	"bytes"
	"fmt"
)

type (
	F func([]string) (any, error)

	Submodule struct {
		Name, Help string
		Keys       []string
		Funcs      map[string]*Function
		Default    *Function
	}

	Function struct {
		Name, Help string
		N          int
		F
	}
)

func (sub *Submodule) Run(args []string) (any, error) {
	try_default := func(args []string) (any, error) {
		f := sub.Default
		if f == nil {
			return nil, fmt.Errorf("module %s does not contain default func", sub.Name)
		}

		if f.N != len(args) {
			return nil, fmt.Errorf("default module for %s require %d args but got %d",
				sub.Name, f.N, len(args))
		}

		return f.F(args)
	}
	if len(args) == 0 {
		return try_default([]string{})
	}

	f, f_ok := sub.Funcs[args[0]]
	if !f_ok {
		return try_default(args[0:])
	}

	return f.F(args[1:])
}

const HELP_FORMAT = "\033[1G%-s %-s %s\n"

var function_bytes = []byte("Function")
var command_bytes = []byte("Command")

const column_whitespace = 3

func (sub *Submodule) PrintHelp() error {
	n_len, c_len := len(function_bytes), len(command_bytes)
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
