package run

import (
	"arit/modules"
	"arit/modules/submodules"
	"errors"
)

const helpMsg = "standard help message"

var AvaliableModules = map[string]*modules.Submodule{}
var Aliases = map[string]string{}

func addSubmodule(sub *modules.Submodule) {
	for _, k := range sub.Keys {
		Aliases[k] = sub.Name
	}

	AvaliableModules[sub.Name] = sub
}

func init() {
	addSubmodule(&submodules.Random)
	addSubmodule(&submodules.Prime)
	addSubmodule(&submodules.Economics)
}

func Parse(args []string) error {
	ste := &State{
		Vars:    map[string]string{},
		Modules: AvaliableModules,
		Aliases: Aliases,
		History: [][]byte{},
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
