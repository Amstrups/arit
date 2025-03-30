package cli

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/term"
)

func shell(ste *State) error {
	var f uintptr
	//var t_state *term.State
	defer func() {
		//term.Restore(int(f), t_state)
	}()

	f = os.Stdin.Fd()
	//t_state, _ := term.MakeRaw(int(f))
	term.MakeRaw(int(f))
	scanner := bufio.NewReader(os.Stdin)

	//c := make(chan os.Signal, 10)

	for {
		fmt.Print("\033[38;2;120;166;248m>>> ")
		fmt.Print("\033[0m")
		b := make([]byte, 3)
		_, err := scanner.Read(b)
		if err != nil {
			return err
		}

		if string(b[0]) == "x" {
			return nil

		}
		//input := scanner.Scan()()
		if b[0] == 27 { // Ascii escape

		}

		fmt.Printf("\"%s\"\n\033[0G", string(b))
		continue

		/*
			if input == "ls" {

			}

			if input == "exit" {
				return nil
			}

			err := ste.ParseRaw(input)
			if err != nil {
				fmt.Println(err)
			}
		*/
	}
}
