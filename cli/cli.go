package cli

import (
	"arit/modules"
	"arit/modules/submodules"
	"errors"
)

const helpMsg = "standard help message"

type MODE int8

const (
	CLI MODE = iota
	SHELL
	SERVER
	TUI
)

var AvaliableModules = map[string]*modules.Submodule{}

func addSubmodule(sub *modules.Submodule) {
	for _, k := range sub.Keys {
		AvaliableModules[k] = sub
	}
}

func init() {
	addSubmodule(&submodules.Random)
	addSubmodule(&submodules.Prime)
}

func Parse(args []string) error {
	ste := &State{
		Vars:    map[string]string{},
		Modules: AvaliableModules,
		History: []string{},
	}

	if len(args) == 0 {
		return shell(ste)
	}

	switch args[0] {
	case "shell", "sh":
		return shell(ste)
	case "ui":
		return ui(ste)
	case "server":
		return errors.New("server NYI")
	default:
		return ste.Parse(args)
	}
}
