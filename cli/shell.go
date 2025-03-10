package cli

import (
	"bufio"
	"fmt"
	"os"
)

func shell(ste *State) error {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\033[38;2;120;166;248m>>> ")
		fmt.Print("\033[0m")
		scanner.Scan()

		input := scanner.Text()

		if input == "ls" {

		}

		if input == "exit" {
			return nil
		}

		ste.ParseRaw(input)
	}
}
