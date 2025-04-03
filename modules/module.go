package modules

import (
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

const HELP_FORMAT = "\033[1G%-10s %s\n"

var HELP_HEADER_FORMAT = "\033[1m" + HELP_FORMAT + "\033[0m"

func (sub *Submodule) PrintHelp() error {
	fmt.Printf(HELP_HEADER_FORMAT, "Function", "Description")
	fmt.Printf(HELP_FORMAT, "Foo", "Baa")

	return nil
}
