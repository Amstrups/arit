package parser

import (
	"fmt"
	"io"
	"strconv"
)

type CmdToken int

const (
	UNKNOWN = iota
	EOF
	IDENT
	NUMBER
	STRING
	FLAG
	FLAG_ARG_IDENT
	FLAG_ARG_STRING
	FLAG_ARG_NUMBER
	PIPE
)

type (
	Argument struct {
		T CmdToken
		V any
	}

	Flag interface {
		flag()
	}

	StringFlag struct {
		Flag, Value string
	}

	IdentFlag struct {
		Flag, Value string
	}

	NumberFlag struct {
		Flag  string
		Value int64
	}

	BoolFlag struct {
		Flag  string
		Value bool
	}
)

func (StringFlag) flag() {}
func (IdentFlag) flag()  {}
func (NumberFlag) flag() {}
func (BoolFlag) flag()   {}

type Command struct {
	In  *io.Reader
	Out *io.Writer

	Module string
	Func   string
	Args   []Argument
	Flags  []Flag
}

func (c *Command) AddStringFlag(vi string, vj any) error {
	vtj, ok := vj.(string)
	if !ok {
		return fmt.Errorf("givem flag value %v could not be parsed as string", vj)
	}

	c.Flags = append(c.Flags,
		StringFlag{
			Flag:  vi[2:],
			Value: vtj,
		})

	return nil
}

func (c *Command) AddIdentFlag(vi string, vj any) error {
	vtj, ok := vj.(string)
	if !ok {
		return fmt.Errorf("givem flag value %v could not be parsed as string", vj)
	}

	if vtj == "true" || vtj == "false" {
		c.Flags = append(c.Flags,
			BoolFlag{
				Flag:  vi[2:],
				Value: vtj == "true",
			})
		return nil
	}

	c.Flags = append(c.Flags,
		IdentFlag{
			Flag:  vi[2:],
			Value: vtj,
		})

	return nil

}

func (c *Command) AddNumberFlag(vi string, vj string) error {
	vtj, err := strconv.ParseInt(vj, 0, 64)
	if err != nil {
		return fmt.Errorf("givem flag value %v could not be parsed as int. %e", vj, err)
	}

	c.Flags = append(c.Flags,
		NumberFlag{
			Flag:  vi[2:],
			Value: vtj,
		})

	return nil
}
