package run

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"

	"golang.org/x/term"
)

func shell(ste *State, tui bool) error {
	var before func() error = repl_pre
	var after func(string, error) error = repl_post
	var exit func() error = func() error { fmt.Print(" goodbye"); return nil }

	if tui {
		// change pre and post
	}

	var f uintptr
	var t_state *term.State
	defer func() {
		exit()
		term.Restore(int(f), t_state)
	}()

	f = os.Stdin.Fd()
	t_state, _ = term.MakeRaw(int(f))
	scanner := bufio.NewReader(os.Stdin)

	c := make(chan os.Signal, 10)
	signal.Notify(c, os.Interrupt)

	for {
		if err := before(); err != nil {
			return err
		}
		val := []byte{}
		b := make([]byte, 3)
	read:
		for {
			n, err := scanner.Read(b)
			//fmt.Print(b[0])
			if err != nil {
				return err
			}
			switch n {
			case 1:
				switch b[0] {
				case 13: // cook on enter
					fmt.Println()
					break read
				case 127: // delete on backspace
					ln := len(val)
					if ln == 0 {
						continue
					}

					fmt.Print("\033[1D \033[1D")
					val = val[:len(val)-1]

				case 120: // exit on 'x'
					return nil
				default:
					fmt.Print(string(b[0]))
					val = append(val, b[0])
				}
			case 2:
				fmt.Printf("the fuck is %s\n", b[0:1])
			case 3:
				fmt.Printf("Ints: %d %d %d\n", int(b[0]), int(b[1]), int(b[2]))
				break read
			}
		}

		fmt.Print("\033[0G")
		//fmt.Printf("\"%s\"\n\033[0G", string(b))

		if err := after("", nil); err != nil { //TODO: proper use of 'after'
			return err
		}

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

func repl_pre() error {
	fmt.Print("\033[38;2;120;166;248m>>> ")
	fmt.Print("\033[0m")
	return nil
}

func repl_post(val string, err error) error {
	return nil
}
