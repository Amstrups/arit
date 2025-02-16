package cli

import (
	"arit/modules"
	"errors"
)

const helpMsg = "standard help message"

func Parse(args []string) error {
	if len(args) == 0 {
		return errors.New("Cannot process empty list of args")
	}

	ste := &State{
		Vars:   map[IDENT]any{},
		Module: modules.Full(),
	}

	switch args[0] {
	case "shell", "sh":
		return shell(ste)
	case "ui":
		return ui()
	case "server":
		return errors.New("server NYI")
	default:
		return ste.Parse(args)
	}
}
