package modules

import (
	"arit/modules/util"
	"fmt"
)

type (
	F func([]string) (any, error)

	Submodule struct {
		Name, Help string
		Keys       []string
		Funcs      map[string]*Function
	}

	Function struct {
		Name, Help string
		N          int
		F
	}
)

func (sub *Submodule) Run(args []string) (any, error) {
	if len(args) == 0 {
		f, ok := sub.Funcs[util.DEFAULT_KEY]
		if !ok {
			return nil, fmt.Errorf("module %s does not contain default func", sub.Name)
		}

		if f.N > 0 {
			return nil, fmt.Errorf("default module for %s require args", sub.Name)
		}

		return f.F([]string{})
	}

	f, f_ok := sub.Funcs[args[0]]
	if !f_ok {
		return nil, fmt.Errorf("%s %s: unknown command", args[0], args[1])
	}

	return f.F(args[1:])
}
