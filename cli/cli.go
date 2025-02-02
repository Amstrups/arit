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

	switch args[0] {
	case "shell", "sh":
		return shell()
	case "ui":
		return ui()
	case "server":
		return errors.New("server NYI")
	default:
		m, err := modules.New(args)
		if err != nil {
			return err
		}

		return m.Parse()
	}
}
