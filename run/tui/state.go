package tui

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Style struct {
	Vert                    string
	TopHori, BottomHori     string
	Topleft, Topright       string
	Bottomleft, Bottomright string
}

func Fancy() Style {
	return Style{
		Vert:        BORDER_HORI,
		TopHori:     BORDER_VERT,
		BottomHori:  BORDER_VERT,
		Topleft:     BORDER_TL,
		Topright:    BORDER_TR,
		Bottomleft:  BORDER_BL,
		Bottomright: BORDER_BR,
	}
}

func Default() Style {
	return Style{
		TopHori:     BORDER_HORI,
		BottomHori:  SPACE,
		Vert:        SPACE,
		Topleft:     CORNER_TL,
		Topright:    CORNER_TR,
		Bottomleft:  CORNER_BL,
		Bottomright: CORNER_BR,
	}
}

func Space() Style {
	return Style{
		Vert:        SPACE,
		TopHori:     SPACE,
		BottomHori:  SPACE,
		Topleft:     SPACE,
		Topright:    SPACE,
		Bottomleft:  SPACE,
		Bottomright: SPACE,
	}
}

type State struct {
	X, Y   int
	bg, fg string
	screen *bufio.Writer
	Style
}

func New() *State {
	state := &State{
		screen: bufio.NewWriter(os.Stdout),
		Style:  Space(),
	}
	return state
}

func (s *State) Set(x, y int) {
	s.X = x
	s.Y = y
}

func (s *State) Newline(offset int) {
	s.Moveto(s.X+1, offset)
}

func (s *State) Moveto(x int, y int) {
	s.Set(x, y)
	fmt.Fprintf(s.screen, "\033[%d;%dH", x, y)
}

func (s *State) MoveToInner(sq Square, offset int) {
	s.Moveto(sq.X+offset, sq.Y+offset+1)
}

func (s *State) Moveright() {
	s.Moveto(s.X, s.Y+1)
}

func (s *State) Clear() {
	s.Set(1, 1)
	fmt.Fprint(s.screen, "\033[2J")
	fmt.Fprint(s.screen, "\033[0;0H")
}

func (s *State) Insert(str string) {
	fmt.Fprint(s.screen, str)
}

func (s *State) Insertln(str string) {
	fmt.Fprintln(s.screen, str)
}

func (s *State) Render() {
	s.screen.Flush()
}

func top_bottom(s Style, inner_n int, title string) ([]string, string) {
	bottom := s.Bottomleft + strings.Repeat(s.BottomHori, inner_n) + s.Bottomright

	prefix := 2
	spaces := 2
	used := len(title) + prefix + spaces

	tops := []string{
		s.Topleft + strings.Repeat(" ", inner_n) + s.Topright,
		fmt.Sprintf(" %s %s %s ", strings.Repeat(s.TopHori, prefix), title, strings.Repeat(s.TopHori, inner_n-used)),
		strings.Repeat(" ", inner_n+2),
	}

	return tops, bottom
}

func (s *State) DrawBlank(sq Square) {
	line := strings.Repeat(" ", sq.W) + "\n"
	lines := strings.Repeat(line, sq.H)
	s.Insert(lines)
}

func (s *State) DrawSquare(sq *Square, title string) {
	s.Insert(BGDarker)
	s.Moveto(sq.X, sq.Y)
	tops, bottom := top_bottom(s.Style, sq.W-2, title)
	std_mid := s.Vert + strings.Repeat(" ", sq.W-2) + s.Vert
	for _, t := range tops {
		s.Insert(t)
		s.Newline(sq.Y)
	}

	s.Insert(BG2)

	for i := len(tops); i < sq.H-1; i++ {
		//s.Insert(colors[count])
		s.Insert(std_mid)
		s.Newline(sq.Y)
	}

	s.Insert(bottom)

	sq.A = 2
	sq.B = len(tops)
	sq.D = 1
}

func (s *State) DrawLines(h int) {
	x := 1

	for x < h-1 {
		s.Moveto(x, 1)
		s.Insert(BG1 + "                    ")
		s.Moveto(x+1, 1)
		s.Insert(BG2 + "                    ")
		x += 2
	}

}
