package cli

import (
	"arit/modules/prime"
	"arit/modules/random"
	"errors"
	"fmt"
	"strings"
)

func Parse(args []string) error {
	if len(args) == 0 {
		return errors.New("Cannot process empty list of args")
	}

	switch args[0] {
	case "help":
		msg := ""
		if len(args) > 1 {
			msg = help(args[1:])
		} else {
			msg = "standard help messsage"
		}
		fmt.Println(msg)
		return nil
	case "shell", "sh":
		return shell()
	case "ui":
		return ui()
	case "server":
		return errors.New("server NYI")
	default:
		return module(args)
	}
}

func help(args []string) string {
	if len(args) > 1 {
		return "unknown topic: " + strings.Join(args, " ")
	}

	mod := args[0]

	switch mod {
	case "r", "rand", "random":
		return random.Help()
	case "p", "prime":
		return prime.Help()
	default:
		return "unknown module: " + mod
	}

}

func module(args []string) error {
	mod := args[0]
	modArgs := args[1:]

	switch mod {
	case "r", "rand", "random":
		val, err := random.Eval(modArgs)
		if err != nil {
			return err
		}
		fmt.Printf("%v\n", val)
		return nil
	case "p", "prime":
		val, err := prime.Eval(modArgs)
		if err != nil {
			return err
		}
		fmt.Printf("%v\n", val)
		return nil
	default:

		return errors.New("arg: " + args[0] + " is not supported")
	}
}
