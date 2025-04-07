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

func shell(ste *State, dimx, dimy int) error {
	//logger := NewLogger()
	//defer logger.Close()

	f := os.Stdin.Fd()
	t_state, _ := term.MakeRaw(int(f))
	defer term.Restore(int(f), t_state)

	w, h, err := term.GetSize(int(f))
	if err != nil {
		panic(err)
	}

	log_y := 1

	scanner := bufio.NewReader(os.Stdin)

	fmt.Print("\033[?2004h")
	defer fmt.Print("\033[?2004l")

	y := len(ste.History)

	var b []byte
	var ln, x int

	insert := func(xs ...byte) {
		n := ste.insert(x, ln, xs...)
		fmt.Printf("\033[s%s\033[u\033[%dC", ste.current[x:], n)

		x += n
		ln += n
	}

	log := func(x string) {
		fmt.Print("\033[s") // save pos, new line, move left
		fmt.Printf("\033[%d;%dH%-60s\n\033[60D%s", log_y, w-62, x, SPACES[:62])

		fmt.Print("\033[u") // restore pos
		log_y = max((log_y+1)%h, 1)
	}

	log_input := func(n int) {
		str := fmt.Sprintf("\"%s\"", b[:n])

		if n == 6 {
			str = "PASTE"
		}

		if n == 1 {
			switch b[0] {
			case 3:
				str = "INTERRUPT"
				log_y--
			case 13:
				str = "NEWLINE"
				log_y--
			case 27:
				str = "ESCAPE"
			case 127:
				str = "BACKSPACE"
			}
		}

		log(fmt.Sprintf("%3v= %-10s - read: %d, rem:%d", b, str, n, scanner.Buffered()))
	}
	newline := func() {
		fmt.Printf("\n\033[%dG", dimx)
	}

	test_input := func(input string) {
		ste.current = []byte(input)
		ln = len(ste.current)
		fmt.Print("\033[1G")
		fmt.Print(SHELL_PREFIX)
		fmt.Printf("[RUNNING_TEST_INPUT] %s", ste.current)
		newline()
	}

cooking:
	for {
		fmt.Print(SHELL_PREFIX)

		ste.current = make([]byte, 128)
		b = make([]byte, 6)

		ln, x = 0, 0

	read:
		for {
			b = make([]byte, 6)
			n, err := scanner.Read(b)
			if err != nil {
				return err
			}

			log_input(n)

			switch n {
			case 1: // ascii
				switch b[0] {
				case 3: // <C-c>
					newline()
					continue cooking
				case 4: // <C-d>, kill process
					return nil

				case 20: // <C-t>, easy-access test
					test_input("random number |> store x")
					break read

				case 21: // <C-u>, easy-access test2
					test_input("rand cap \"hello world\"")
					break read

				case 25: // <C-y>, easy-access test3
					test_input("rand closed 1 10 |> store y")
					break read

				case 12: // <C-l>
				case 27: // ESC
					continue

				case 13: // cook on enter
					newline()
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
				scanner.Read(b_)
				scanner.Discard(6)
				insert(b_...)
			}
		}

		if ln == 0 {
			continue
		}

		ste.History = append(ste.History, ste.current[:ln])
		y++

		err := ste.ParseRaw(string(ste.current[:ln]))
		if err != nil {
			fmt.Printf("\033[1G%s\n", err.Error())
		}
	}
}
