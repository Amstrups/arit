package run

import (
	"fmt"
)

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

func (s *State) EchoStored(width int) {
	for k := range s.Vars {
		fmt.Println(s.VarToString(k, width))
	}
}

func (s *State) VarToString(ident string, w int) string {
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
