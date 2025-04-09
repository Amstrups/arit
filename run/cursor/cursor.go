package cursor

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
TODO: consider switching to chan for prints,
and make flush consume channel
*/

type Cursor struct {
	X, Y int // X:vertical, Y:horizontal
	x, y int // x:vertical, y:horizontal
	h, w int

	bg, fg string
	buffer *bufio.Writer
	screen *bufio.Writer

	// Style // TODO: combine tui and shell cursor logic

	Prefix string
	stx    int

	lines  int
	offset int
	mode   int // 0: print, 1: input

}

var l *bufio.Writer

func New(x, y, w, h int, f_prefix, prefix string) *Cursor {
	log_file, err := os.Create("logs/log")
	if err != nil {
		panic(err)
	}

	prefixln := len(prefix)

	state := &Cursor{
		screen: bufio.NewWriter(os.Stdout),
		x:      x,
		y:      y,
		w:      w,
		Prefix: fmt.Sprintf(f_prefix, prefix),
		stx:    x + prefixln,
		offset: prefixln,
		buffer: bufio.NewWriter(log_file),
	}
	return state
}

func (s *Cursor) Focus() {
	fmt.Fprintf(s.screen, F_POSITION, s.stx+s.X, s.y+s.Y)
}

func (c *Cursor) ResetLine() {
	fmt.Fprint(c.screen, CLEAR_LEFT)
	c.Pprefix()
}

func (c *Cursor) Pprefix() {
	fmt.Fprintf(c.screen, F_ABSOLUTE_COLUMN, c.x)
	fmt.Fprint(c.screen, c.Prefix)
	c.mode = 1
	c.lines = 0
}

func (s *Cursor) Moveright() {
	fmt.Fprint(s.screen, "\033[1C")
	s.Y = min(s.Y+1, s.w)
}

func (s *Cursor) RelativeColumn(n int) {
	fmt.Fprintf(s.screen, F_ABSOLUTE_COLUMN, s.stx+n)
}

func (s *Cursor) Moveleft() {
	fmt.Fprint(s.screen, "\033[1D")
	s.Y = max(1, s.Y-1)
}

func (s *Cursor) Insert(str string) {
	a, b := 0, len(str)

	offset := s.mode * s.offset

	for (b-a)+s.Y+offset > s.w {
		avaliable := a + s.w - s.X - offset

		s.mode = 0
		offset = 0

		n, _ := fmt.Fprintf(s.screen, NEWLINE_START_CONCAT, str[a:min(avaliable, b)])
		fmt.Fprintf(s.buffer, "n:%d\n", n)
		s.buffer.Flush()

		a += n

		s.Y = 0
		s.X++

		s.lines++
	}

	n, _ := fmt.Fprintf(s.screen, "%s", str[a:])
	fmt.Fprintf(s.buffer, "n:%d\n", n)
	s.buffer.Flush()
	s.Y += n
}

func (s *Cursor) Insertf(format string, args ...any) {
	s.Insert(fmt.Sprintf(format, args...))
}

func (s *Cursor) InsertBytes(str []byte) {
	s.Insert(fmt.Sprintf("%s", str))
}

func (s *Cursor) InsertAtNewline(str string) {
	s.Insert(str)
	s.Newline()
}

func (s *Cursor) InsertAtNewlinef(format string, args ...any) {
	s.InsertAtNewline(fmt.Sprintf(format, args...))
}

func (s *Cursor) InsertAnyAtNewline(v any) {
	s.InsertAtNewline(fmt.Sprintf(NEWLINE_START_CONCAT_V, v))
}

func (s *Cursor) Newline() {
	fmt.Fprintf(s.screen, NEWLINE_START)
	s.X++
	s.Y = 0
}

func (s *Cursor) Render() {
	//fmt.Fprintf(s.buffer, "%#v", s.screen.)
	s.screen.Flush()
}

func (s *Cursor) Save() {
	fmt.Fprint(s.screen, SAVE)
}

func (s *Cursor) Restore() {
	fmt.Fprint(s.screen, RESTORE)
}

func (s *Cursor) Space() {
	fmt.Fprintf(s.screen, STR_SPACE)
}

func (s *Cursor) Spaces(n int) {
	fmt.Fprint(s.screen, strings.Repeat(STR_SPACE, n))
}
