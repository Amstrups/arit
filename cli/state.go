package cli

import (
	"arit/cli/parser"
	"arit/modules"
	"fmt"
)

type IDENT string

type State struct {
	Vars map[IDENT]any
	modules.Module
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
	/*
			switch xt := x.(type) {
			case int:
		    return fmt.Sprintf("%s: %d", xt)
		  case []int:
		    return
			default
			default:
				return fmt.Sprintf("unknown type with value %v", xt)

			}
	*/

}

func (s *State) Parse(cmds []parser.Command) error {
	if len(cmds) > 1 {
		panic("pipe is not supported in arit yet")
	}

	cmd := cmds[0]

	mod, ok := s.Submodules[cmd.Module]

	if !ok {
		fmt.Printf("Could not locate a module with name %s", cmd.Module)
		return nil
	}

	value, err := mod.Parse(cmd)
	if err != nil {
		return err
	}

	fmt.Println(value)

	return nil
}
