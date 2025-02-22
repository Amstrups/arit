package cli

import (
	"arit/cli/parser"
	"arit/modules"
	"errors"
)

const helpMsg = "standard help message"

func Parse(args []string) error {
	if len(args) == 0 {
		return errors.New("Cannot process empty list of args")
	}

	parser := parser.CmdParser{}

	cmds := parser.ParseArgs(args)

	ste := &State{
		Vars:   map[IDENT]any{},
		Module: modules.Full(),
	}

	if len(cmds) == 0 {
		panic("cannot process emtpy arguments list")
	}

	switch cmds[0].Module {
	case "shell", "sh":
		return shell(ste)
	case "ui":
		return ui()
	case "server":
		return errors.New("server NYI")
	default:
		return ste.Parse(cmds)
	}
}
