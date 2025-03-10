package cli

import (
	"arit/modules"
	"fmt"
	"strings"
)

type IDENT string

type State struct {
	Vars    map[IDENT]any
	Modules map[string]*modules.Submodule
}

func (s *State) EchoStored(width int) {
	for k := range s.Vars {
		fmt.Println(s.ToString(k, width))
	}
}

func (s *State) ToString(ident IDENT, w int) string {
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

func (s *State) ParseRaw(args string) error {
	return s.Parse(strings.Split(args, " "))
}

func (s *State) Parse(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("cannot parse empty list of args")
	}

	sub, ok := s.Modules[args[0]]
	if !ok {
		return fmt.Errorf("%s: unknown command", args[0])
	}

	value, err := sub.Run(args[1:])
	if err != nil {
		return err
	}

	fmt.Println(value)

	return nil
}
