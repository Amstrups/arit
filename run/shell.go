package run

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	"golang.org/x/term"
)

const (
	PASTE_HEAD = "\033[200~"
	PASTE_TAIL = "\033[201~"

	SHELL_PREFIX = "\033[1G\033[38;2;120;166;248m>>> \033[0m"
)

var SPACES = bytes.Repeat([]byte("        "), 16)

func shell(ste *State) error {
	//logger := NewLogger()
	//defer logger.Close()

	f := os.Stdin.Fd()
	t_state, _ := term.MakeRaw(int(f))
	defer term.Restore(int(f), t_state)

	scanner := bufio.NewReader(os.Stdin)

	fmt.Print("\033[?2004h")
	defer fmt.Print("\033[?2004l")

	y := len(ste.History)

	var b []byte
	var ln, x int

	/*
			log := func(msg string) {
				fmt.Printf("\033[s\n\033[0G%s", msg)
				fmt.Print("\n\033[0G")
				before()
				fmt.Printf("%s\033[u", ste.current)
			}

		log_line := func(msg string) {
			fmt.Print(msg)
			fmt.Print("\n\033[1G")
		}
	*/

	insert := func(xs ...byte) {
		n := ste.insert(x, ln, xs...)
		fmt.Printf("\033[s%s\033[u\033[%dC", ste.current[x:], n)

		x += n
		ln += n
	}

	for {
		fmt.Print(SHELL_PREFIX)

		ste.current = make([]byte, 128)
		b = make([]byte, 6)

		ln, x = 0, 0

	read:
		for {
			n, err := scanner.Read(b)
			if err != nil {
				return err
			}

			switch n {
			case 1: // ascii
				switch b[0] {
				case 13: // cook on enter
					fmt.Print("\n\033[1G")
					break read
				case 127: // delete on backspace
					if x == 0 {
						continue
					}

					ste.remove(x, ln, 1)

					x--
					ln--

					fmt.Print("\033[1D\033[s")
					fmt.Printf("%s \033[u", ste.current[x:ln])
				case 'Q': // exit on 'Q'
					return nil
				default:
					insert(b[0])
				}

			case 2: // weird symbols
				fmt.Printf("the fuck is '%s' (%d, %d)\n", b[0:1], b[0], b[1])
				break read

			case 3: // arrow keys
				switch b[2] {
				case 'A': // up
					if y < len(ste.History) {
						ste.History[y] = ste.current[:ln]
					}

					if y == 0 {
						continue
					}

					fmt.Printf("%s\033[s%s\033[u", SHELL_PREFIX, SPACES[:ln])

					y--

					ln = len(ste.History[y])
					ste.current = make([]byte, 128)
					copy(ste.current[:ln], ste.History[y])
					x = max(x, ln)

					fmt.Print(SHELL_PREFIX)
					fmt.Printf("%s", ste.current[:ln])

					if x != ln {
						fmt.Print("\033u")
					}

				case 'B': // down
					if y >= len(ste.History) {
						continue
					}

					y++

					ste.current = make([]byte, 128)
					fmt.Printf("%s\033[s%s\033[u", SHELL_PREFIX, SPACES[:ln])

					if y == len(ste.History) {
						ln = 0
						x = 0
						continue
					}

					fmt.Print("\033[s")

					ln = len(ste.History[y])
					copy(ste.current[:ln], ste.History[y])
					x = max(x, ln)
					fmt.Print(SHELL_PREFIX)
					fmt.Printf("%s", ste.current[:ln])
					if x != ln {
						fmt.Print("\033u")
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
					fmt.Printf("Unknown: %v => '%s'\n", b[:n], b)
					break read
				}

			case 6: // paste brackets
				if string(b) != PASTE_HEAD {

					fmt.Printf("expected %v\n", []byte(PASTE_HEAD))
					return fmt.Errorf("unknown head at read: %v", b)
				}

				to_read := scanner.Buffered()
				if to_read <= 6 {
					return fmt.Errorf("expecting paste to contain tail but buffered was %d long", to_read)
				}

				b_ := make([]byte, to_read-6)
				insert(b_...)
			}
		}

		//	fmt.Print("\034[1G")

		ste.History = append(ste.History, ste.current[:ln])
		y++

		err := ste.ParseRaw(string(ste.current[:ln]))
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
