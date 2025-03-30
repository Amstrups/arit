package cli

import (
	"arit/modules"
	"fmt"
	"strings"
)

type State struct {
	Vars    map[string]string
	Modules map[string]*modules.Submodule
	History []string
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

func (s *State) Fetch(args []string) (string, error) {
	if len(args) > 1 {
		return "", fmt.Errorf("unknown args passed to get command: %v", args)
	}

	v, ok := s.Vars[args[0]]
	if !ok {
		return "", fmt.Errorf("unknown var %s", args[0])
	}
	return v, nil
}

func (s *State) ParseRaw(args string) error {
	return s.Parse(strings.Split(strings.Trim(args, " "), " "))
}

func (s *State) Parse(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("cannot parse empty list of args")
	}

	if args[0] == "store" || args[0] == "set" {
		return s.Store(args[1:])
	}

	sub, ok := s.Modules[args[0]]
	if !ok {
		v, err := s.Fetch(args)
		if err != nil {
			return fmt.Errorf("%s: unknown command", args[0])
		}

		fmt.Println(v)
		return nil
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
