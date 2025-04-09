package run

import (
	"arit/run/cursor"
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
)

const linecount = 2603

type logger struct {
	*bufio.Writer
}

func fffy() string {
	file, err := os.Open("bin/ffs")
	if err != nil {
		panic(err)
	}

	i := rand.Intn(2603)
	c := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if c >= i {
			return line

		}
		c++

	}

	return "I ran out of fun facts."
}

func TTS() {
	fmt.Print("Starting logger...")

	msg := fmt.Sprintf(
		"\"Hey, did you know? %s\"", fffy(),
	)

	cmd := exec.Command("say", "-v", "Samantha", msg)

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func NewLogger(file string) *logger {
	log_file, err := os.Create(file)
	writer := bufio.NewWriter(log_file)
	if err != nil {
		panic(err)
	}

	return &logger{
		Writer: writer,
	}
}

func (logger *logger) Log(x string) {
	fmt.Fprintf(logger, "%s\n", x)
	logger.Flush()
}

func (logger *logger) LogInput(n int, b []byte, buffered int) {
	str := fmt.Sprintf("%q", b[:n])

	if n == 6 {
		str = "PASTE"
	}

	if n == 1 {
		switch b[0] {
		case cursor.CtrlL:
			str = "FORM FEED"
		case cursor.CtrlC:
			str = "INTERRUPT"
		case cursor.CR:
			str = "NEWLINE"
		case cursor.ESC:
			str = "ESCAPE"
		case cursor.BACKSPACE:
			str = "BACKSPACE"
		case cursor.SPACE:
			str = "SPACE"
		}
	}

	logger.Log(fmt.Sprintf("%3v= %-10s - read: %d, rem:%d", b, str, n, buffered))
}
