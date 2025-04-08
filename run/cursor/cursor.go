package cursor

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
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
	offset int
}

func New(x, y, w, h int, prefix string) *Cursor {
	state := &Cursor{
		screen: bufio.NewWriter(&bytes.Buffer{}),
		buffer: bufio.NewWriter(os.Stdout),
	}
	return state
}

func (s *Cursor) Focus() {
	fmt.Fprintf(s.screen, F_POSITION, s.x+s.X, s.y+s.Y)
}

func (s *Cursor) Set(x, y int) {
	s.X = x
	s.Y = y
}

func (s *Cursor) Moveto(x int, y int) {
	s.Set(x, y)
	fmt.Fprintf(s.screen, "\033[%d;%dH", x, y)
}

func (s *Cursor) Moveright() {
	fmt.Fprint(s.screen, "\033[1C")
	s.Y = min(s.Y+1, s.w)
}

func (s *Cursor) Moveleft() {
	fmt.Fprint(s.screen, "\033[1D")
	s.Y = max(1, s.Y-1)
}

func (s *Cursor) Clear() {
	s.Set(1, 1)
	fmt.Fprint(s.screen, "\033[2J")
	fmt.Fprint(s.screen, "\033[0;0H")
}

func (s *Cursor) Insert(str string) {
	fmt.Fprint(s.screen, str)
}

func (s *Cursor) InsertBytes(str []byte) {
	fmt.Fprintf(s.screen, "%s", str)
}

func (s *Cursor) Insertln(str string) {
	fmt.Fprintln(s.screen, str)
}

func (s *Cursor) InsertAtNewline(str string) {
	fmt.Fprintf(s.screen, NEWLINE_START_CONCAT, str)
	s.X++
}

func (s *Cursor) InsertAtNewline2(format string, args ...any) {
	fmt.Fprintf(s.screen, NEWLINE_START_CONCAT, fmt.Sprintf(format, args...))
	s.X++
}

func (s *Cursor) InsertAnyAtNewline(v any) {
	fmt.Fprintf(s.screen, NEWLINE_START_CONCAT_V, v)
	s.X++
}

func (s *Cursor) Newline() {
	fmt.Fprintf(s.screen, NEWLINE_START)
	s.X++
}

func (s *Cursor) Render() {
	fmt.Fprintf(s.buffer, "%#v", s.screen)
	s.buffer.Flush()
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
