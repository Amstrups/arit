package cli

import (
	"arit/cli/parser"
	"arit/modules"
	"errors"
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

const helpMsg = "standard help message"

func Parse(args []string) error {
	if len(args) == 0 {
		return errors.New("Cannot process empty list of args")
	}

	parser := parser.CmdParser{}
  fmt.Println(args)
	parser.ParseArgs(args)
  spew.Dump(parser.Commands)
	return nil


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
