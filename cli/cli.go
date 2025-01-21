package cli

import (
	"arit/ui"
	"errors"
	"fmt"
)

func help(input []string) string {
	if len(input) == 0 {
		return "Cannot process empty list of args"
	}
	switch {
	default:
		return "unknown topic: " + input[0]
	}

}

func Parse(input []string) error {
	if len(input) == 0 {
		return errors.New("Cannot process empty list of args")
	}

	switch input[0] {
	case "help":
		msg := ""
		if len(input) > 1 {
			msg = help(input[1:])
		} else {
			msg = "standard help messsage"
		}
		fmt.Println(msg)
		return nil

	case "shell", "sh":
		return shell()
	case "ui":
		return ui.Run()
	case "rand":
		return nil
	default:
		return errors.New("arg: " + input[0] + " is not supported yet")
	}
}
