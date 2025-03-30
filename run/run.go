package run

import (
	"arit/modules"
	"arit/modules/submodules"
	"errors"
	"fmt"
)

const helpMsg = "standard help message"

var AvaliableModules = map[string]*modules.Submodule{}

func addSubmodule(sub *modules.Submodule) {
	for _, k := range sub.Keys {
		AvaliableModules[k] = sub
	}
}

func init() {
	addSubmodule(&submodules.Random)
	addSubmodule(&submodules.Prime)
	addSubmodule(&submodules.Economics)
}

func Parse(args []string) error {
	fmt.Println(args)
	ste := &State{
		Vars:    map[string]string{},
		Modules: AvaliableModules,
		History: []string{},
	}

	if len(args) == 0 {
		return shell(ste, false)
	}

	switch args[0] {
	case "shell", "sh":
		return shell(ste, false)
	case "ui":
		return ui(ste)
	case "server":
		return errors.New("server NYI")
	default:
		return ste.Parse(args)
	}
}
