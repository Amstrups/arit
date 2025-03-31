package run

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
)

const linecount = 2603

type logger struct {
	io.WriteCloser
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

func NewLogger() *logger {
	fmt.Print("Starting logger...")

	msg := fmt.Sprintf(
		"\"Hey, did you know? %s\"", fffy(),
	)

	cmd := exec.Command("say", "-v", "Samantha", msg)
	out, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}

	go func() {
		if err := cmd.Run(); err != nil {
			panic(err)
		}
	}()

	fmt.Println(" Logger attached!")

	return &logger{out}
}
