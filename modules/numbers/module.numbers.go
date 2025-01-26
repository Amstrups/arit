package numbers

import (
	"errors"
)


func Help() string {
	return "topic \"help rand\" NYI"
}

func Eval(args []string) (any, error) {

	switch args[0] {
	case "cap", "Capitilization":
		if len(args[1:]) != 1 {
			return nil, errors.New("command \"arit rand capitilization\" requires exactly 1 argument")
		}
		return capitilization(trim(args[1])), nil
	case "num", "number":
		return number(), nil
	default:
		return nil, errors.New("arg: " + args[0] + " is not supported")
	}
}

