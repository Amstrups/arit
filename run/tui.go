package run

import (
	"arit/run/tui"
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"golang.org/x/term"
)

// returns (width,height)
func getSize() (int, int) {

	w, h, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		panic(err)
	}

	return w, h
}

func draw(ste *tui.State) (tui.Square, tui.Square) {
	w, h := getSize()

	ste.Insert("\033[48;5;146m")
	background := tui.Square{
		X: 1,
		Y: 1,
		W: w,
		H: h,
	}

	ste.DrawBlank(background)
	ste.Insert("\033[48;5;23m")

	sq1 := tui.Square{
		X: 2,
		Y: 2,
		W: int(w/2) - 1,
		H: h - 2,
	}

	sq2 := sq1
	sq2.Y += w - sq1.W - 2
	//ste.Insert("\033[38;5;146m")

	ste.DrawSquare(&sq1, "not used")
	ste.MoveToInner(sq1, 1)
	ste.Insert("History:")
	ste.DrawSquare(&sq2, "not used")

	ste.MoveToInner(sq2, 1)
	ste.Render()

	return sq1, sq2
}

func altbuff_up() []byte {
	cmd := exec.Command("tput", "smcup")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		panic("switch to alt buff failed")
	}

	return out
}

func altbuff_down() []byte {
	cmd := exec.Command("tput", "rmcup")
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	return out
}

func listen(state *tui.State) {
	c := make(chan os.Signal, 10)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGWINCH) // resize signal
	for msg := range c {
		switch msg {
		case os.Interrupt:
			time.Sleep(time.Millisecond * 200)
			os.Stdout.Write(altbuff_down())
			os.Exit(1)
		default:
			state.Newline(state.Y)
			state.Insert("true")
			state.Newline(state.Y)
			state.Insert(">>> ")
			state.Moveto(state.X, state.Y+2)
		}
	}
}

func read(ui *tui.State, state *State) {
	file, err := os.Create("logs/log.txt")
	if err != nil {
		fmt.Println(err)
		panic(0)
	}

	handle := func(string) {
		ui.Newline(ui.Y - 4)
		ui.Insert("\033[3m")
		ui.Insert("true")
		ui.Newline(ui.Y)
		ui.Insert("\033[23m")
		ui.Insert(">>> ")
		ui.Moveto(ui.X, ui.Y+4)
		ui.Render()

	}

	defer file.Close()

	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		input := scanner.Text()
		state.ParseRaw(input)
		handle(input)
		file.WriteString(input + "\n")
	}

}

func squares() (tui.Square, tui.Square, tui.Square) {
	intPercentage := func(x, p int) int {
		return ((x * 100) - ((100 - p) * x)) / 100

	}
	w, h := getSize()
	h_, w_ := h-3, w-3

	h50 := intPercentage(h_, 65)
	w30 := max(intPercentage(w_, 30), 40)

	topleft := tui.Square{
		X: 2,
		Y: 2,
		H: h_ + 1,
		W: w30,
	}

	topright := tui.Square{
		X: 2,
		Y: w30 + 4,
		H: h_ + 1,
		W: w_ - w30 - 1,
	}

	bottom := tui.Square{
		X: h50 + 3,
		Y: 2,
		H: h_ - h50,
		W: w_ + 1,
	}

	return topleft, topright, bottom
}

func ui(state *State) error {
	os.Stdout.Write(altbuff_up())

	ste := tui.New()
	ste.Clear()

	ste.Style = tui.Default()

	// Draw global background
	ste.Insert(tui.BG1)
	w, h := getSize()
	e := strings.Repeat(tui.SPACE, w)
	for range h {
		ste.Insert(e)
	}

	// Change to squares background color and draw
	topleft, topright, _ := squares()
	ste.Insert(tui.FG2)
	ste.Insert(tui.BGDarker)

	ste.DrawSquare(&topleft, "Variables")
	ste.DrawSquare(&topright, "Shell")
	//ste.DrawSquare(&bottom, "Terminal?")

	ste.Moveto(topright.X+topright.B, topright.Y+topright.A-1)
	prefix := ">>>"
	ste.Insert(prefix)
	ste.Moveto(topright.X+topright.B, topright.Y+topright.A+len(prefix))

	ste.Render()

	go listen(ste)

	read(ste, state)

	return nil
}

func ui_setup() (*tui.State, error) {
	os.Stdout.Write(altbuff_up())
	ste := tui.New()
	ste.Clear()

	ste.Style = tui.Default()

	// Draw global background
	ste.Insert(tui.BG1)
	w, h := getSize()
	e := strings.Repeat(tui.SPACE, w)
	for range h {
		ste.Insert(e)
	}

	// Change to squares background color and draw
	topleft, topright, _ := squares()
	ste.Insert(tui.FG2)
	ste.Insert(tui.BGDarker)

	ste.DrawSquare(&topleft, "Variables")
	ste.DrawSquare(&topright, "Shell")
	//ste.DrawSquare(&bottom, "Terminal?")

	ste.Moveto(topright.X+topright.B, topright.Y+topright.A-1)
	prefix := ">>>"
	ste.Insert(prefix)
	ste.Moveto(topright.X+topright.B, topright.Y+topright.A+len(prefix))

	ste.Render()

	return ste, nil
}
