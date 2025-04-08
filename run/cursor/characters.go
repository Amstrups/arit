package cursor

type CHAR string

/*
Constants named 'F_*'

*/

// Ansi, often used
const (
	NEWLINE_START          = "\n\033[1G" // consider changing to carriage return
	NEWLINE_START_CONCAT   = "%s" + NEWLINE_START
	NEWLINE_START_CONCAT_V = "%v" + NEWLINE_START

	MOVE_TOP_LEFT = "\033[1;1H"

	SAVE    = "\033[s"
	RESTORE = "\033[u"

	F_RGB_BG = "\033[48;2;%d;%d;%dm"
	F_RGB_FG = "\033[38;2;%d;%d;%dm"

	F_POSITION = "\033[%d;%dH" // vertical, then horizontal axis
)

// Ascii ints
const (
	// <C-c>
	CtrlC = 3

	// <C-d>
	CtrlD = 4

	// <C-l>
	CtrlL = 12

	// Enter/carriage return
	CR = 13

	// <C-t>
	CtrlT = 20

	// <C-u>
	CtrlU = 21

	// <C-y>
	CtrlY = 25

	// <ESC>
	ESC = 27

	// BACKSPACE
	BACKSPACE = 127
)

const (
	PASTE_WRAPPER_HEAD = "\033[200~"
	PASTE_WRAPPER_TAIl = "\033[201~"
)

const (
	SPACE     = ' ' // it's just space
	STR_SPACE = string(SPACE)
)

const (
	BORDER_VERT = '\u2503'
	BORDER_HORI = '\u2501'
	BORDER_TL   = '\u250F'
	BORDER_TR   = '\u2513'
	BORDER_BL   = '\u2517'
	BORDER_BR   = '\u251B'
)

const (
	CORNER_BR = '\u25E2'
	CORNER_BL = '\u25E3'
	CORNER_TL = '\u25E4'
	CORNER_TR = '\u25E5'
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

// Preset formats and consts
const (
	F_MODULE_HELP        = "%-20s %-10s %s"
	F_MODULE_HELP_HEADER = "\033[1m" + F_MODULE_HELP + "\033[0m"

	F_SUBMODULE_HELP = "\033[1G%-s %-s %s\n"

	FUNCTION_HEADER = "Function"
	COMMAND_HEADER  = "Command"
)

var (
	FUNCTION_BYTES     = []byte(FUNCTION_HEADER)
	FUNCTION_BYTES_LEN = len(FUNCTION_HEADER)

	COMMAND_BYTES     = []byte(COMMAND_HEADER)
	COMMAND_BYTES_LEN = len(COMMAND_HEADER)
)
