package run

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/term"
)

const (
	PASTE_HEAD = "^[[200~"
	PASTE_TAIL = "^[[201~"
)

func shell(ste *State, tui bool) error {
	var before func() error = repl_pre
	var after func(string, error) error = repl_post
	var exit func() error = func() error { fmt.Print(" goodbye"); return nil }

	logger := NewLogger()

	if tui {
		// change pre and post
	}

	var f uintptr
	var t_state *term.State
	defer func() {
		exit()
		fmt.Print("\033[?2004l")
		term.Restore(int(f), t_state)
		logger.Close()
	}()

	f = os.Stdin.Fd()
	//t_state, _ = term.MakeRaw(int(f))
	scanner := bufio.NewReader(os.Stdin)

	fmt.Print("\033[?2004h")

	/*
		c := make(chan os.Signal, 10)
		signal.Notify(c, os.Interrupt)
	*/

	y := len(ste.History)

	for {
		if err := before(); err != nil {
			return err
		}
		val := make([]byte, 128)
		b := make([]byte, 7)

		var ln, x int
		//fmt.Fprintf(os.Stderr, "Bytes: %v", b)

	read:
		for {
			n, err := scanner.Read(b)
			if err != nil {
				return err
			}

			logger.Write([]byte("abc"))
			logger.Write(b)

			switch n {
			case 1: // ascii
				switch b[0] {
				case 13: // cook on enter
					fmt.Println("\n", string(val))
					break read
				case 127: // delete on backspace
					if ln == 0 || x == 0 {
						continue
					}

					if x < ln {
						for i := range val[x:ln] {
							val[x+i-1] = val[x+i]
						}
					}

					x--
					ln--

					val[ln] = 0
					fmt.Print("\033[1D\033[s")
					fmt.Printf("%v \033[u", string(val[x:]))
				case 120: // exit on 'x'
					return nil
				default:
					if x < ln {
						copy(val[x+1:ln+1], val[x:ln])
					}
					val[x] = b[0]

					fmt.Printf("\033[s%v\033[u\033[1C", string(val[x:]))

					x++
					ln++
				}
			case 2: // weird symbols
				fmt.Printf("the fuck is '%s' (%d, %d)\n", b[0:1], b[0], b[1])
				break read
			case 3: // arrow keys
				switch b[2] {
				case 'A': // up
					if y < len(ste.History) {
						ste.History[y] = val[:ln]
					}
					if y > 0 {
						y--
						ln = len(ste.History[y])
						val = make([]byte, 128)
						copy(val[:ln], ste.History[y])
						x = max(x, ln)
					}

					fmt.Printf("\033[s%v\033[u\033[1C", string(val))

				case 'B': // down
					if y < len(ste.History) {
						ste.History[y] = val[:ln]
						y++
						ln = len(ste.History[y])
						val = make([]byte, 128)
						copy(val[0:ln], ste.History[y])
						x = max(x, ln)
					}

				case 'C': // right
					if x < ln {
						x++
						fmt.Print("\033[1C")
					}

				case 'D': //left
					if x > 0 {
						x--
						fmt.Print("\033[1D")
					}
					continue

				default:
					fmt.Printf("Unknown: %d %d %d => '%s'\n", int(b[0]), int(b[1]), int(b[2]), string(b))
					break read

				}
			case 7: // paste brackets
				if string(b) != PASTE_HEAD {
					return fmt.Errorf("unknown head at read: %s", string(b))
				}
				fmt.Println()

			}
		}

		fmt.Print("\033[0G")

		ste.History = append(ste.History, val[:ln])
		y++
		// fmt.Printf("\"%s\"\n\033[0G", string(b))

		if err := after("", nil); err != nil { //TODO: proper use of 'after'
			return err
		}

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
