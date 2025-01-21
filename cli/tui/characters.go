package tui

type CHAR string

const (
	SPACE = " " // it's just space
)

const (
	BORDER_VERT = "\u2503"
	BORDER_HORI = "\u2501"
	BORDER_TL   = "\u250F"
	BORDER_TR   = "\u2513"
	BORDER_BL   = "\u2517"
	BORDER_BR   = "\u251B"
)

const (
	CORNER_BR = "\u25E2"
	CORNER_BL = "\u25E3"
	CORNER_TL = "\u25E4"
	CORNER_TR = "\u25E5"
)

/*
Set

	background: "\033[48"
	foreground: "\033[38"
*/
const (
	// I wanna say "sky"
	FG1 = "\033[38;2;181;194;205m"
	// Yellow ish
	FG2 = "\033[38;2;243;196;106m"
	// Yellow ish
	BG1 = "\033[48;2;243;196;106m"
	// Blue/gray thing
	BG2 = "\033[48;2;52;71;86m"
	// Blue/gray thing
	BGDarker = "\033[48;2;24;34;41m"
	// I wanna say "sky"
	BG3 = "\033[48;2;181;194;205m"
)
