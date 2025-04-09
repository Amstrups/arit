package run

import (
	"arit/run/cursor"
	"bufio"
	"fmt"
	"os"

	"golang.org/x/term"
)

func shell(ste *State) error {

	f := os.Stdin.Fd()
	t_state, _ := term.MakeRaw(int(f))
	defer term.Restore(int(f), t_state)

	/*
		w, _, err := term.GetSize(int(f))
		if err != nil {
			panic(err)
		}
	*/

	scanner := bufio.NewReader(os.Stdin)

	ste.Insert("\033[?2004h")
	defer ste.Insert("\033[?2004l")

	y := len(ste.History)

	var b []byte
	var ln, x int

	insert := func(xs ...byte) {
		n := ste.insert(x, ln, xs...)
		ln += n

		ste.Save()
		ste.InsertBytes(ste.current[x:ln])
		ste.Restore()
		ste.Insertf(cursor.F_MOVE_RIGHT, n)
		ste.Render()

		x += n

	}

	set_current_to_history := func(clear bool) {
		ste.RelativeColumn(ln)
		ste.ResetLine()
		ste.current = make([]byte, 128)
		ln = 0
		x = 0

		if clear {
			return
		}

		ln = copy(ste.current, ste.History[y])
		x = ln

		ste.InsertBytes(ste.current)
		ste.Render()
	}

	test_input := func(input string) {
		ste.current = []byte(input)
		ln = len(ste.current)
		ste.ResetLine()
		ste.Insertf("[RUNNING_TEST_INPUT] %s", ste.current)
		ste.Newline()
		ste.Render()
	}

cooking:
	for {
		ste.Pprefix()
		ste.Render()

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

			ste.log.LogInput(n, b, scanner.Buffered())

			switch n {
			case 1: // ascii
				switch b[0] {
				case cursor.CtrlC: // <C-c>
					ste.Newline()
					ste.Render()
					continue cooking
				case cursor.CtrlD: // <C-d>, kill process
					return nil

				case cursor.CtrlT: // <C-t>, easy-access test
					test_input("random number |> store x")
					break read

				case cursor.CtrlU: // <C-u>, easy-access test2
					test_input("rand cap \"hello world\"")
					break read

				case cursor.CtrlY: // <C-y>, easy-access test3
					test_input("rand closed 1 10 |> store y")
					break read

				case cursor.CtrlL: // <C-l>
				case cursor.ESC: // ESC
					continue

				case cursor.CR: // cook on enter
					ste.Newline()
					ste.Render()
					break read

				case cursor.BACKSPACE: // delete on backspace
					if x == 0 {
						continue
					}

					ste.remove(x, ln, 1)

					x--
					ln--

					ste.Moveleft()
					ste.Save()
					ste.Insert(string(ste.current[x:ln]))
					ste.Space()
					ste.Restore()
					ste.Render()

				default:
					insert(b[0])
				}

			case 2: // weird symbols
				fmt.Printf("the fuck is '%s' (%d, %d)\n", b[0:2], b[0], b[1])
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

					y--
					set_current_to_history(false)

				case 'B': // down
					if y >= len(ste.History) {
						continue
					}

					ste.Insert(cursor.CLEAR_LEFT)

					y++

					set_current_to_history(y == len(ste.History))

				case 'C': // right
					if x < ln {
						x++
						ste.Cursor.Moveright()
						ste.Render()
					}

				case 'D': //left
					if x > 0 {
						x--
						ste.Cursor.Moveleft()
						ste.Render()
					}
					continue

				default:
					fmt.Printf("Unknown: %v => '%s'\n", b[:n], b)
					break read
				}

			case 6: // paste brackets
				if string(b) != cursor.PASTE_WRAPPER_HEAD {

					ste.Insertf("expected %v\n", []byte(cursor.PASTE_WRAPPER_HEAD))
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
			switch err2 := err.(type) {
			case joinedError:
				for _, e := range err2.errors {
					ste.InsertAtNewline(e.Error())
				}

			default:
				ste.InsertAtNewline(err.Error())
			}

		}
		ste.Render()
	}
}
