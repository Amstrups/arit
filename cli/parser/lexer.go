package parser

import (
	"fmt"
	"regexp"
	"strings"
)

type CmdReader interface {
	Read() (string, error)
}

type CmdLexer struct {
	input string
	head  int
}

func (l *CmdLexer) Next() (x string, t CmdToken) {
	defer func() {
		x = strings.TrimSpace(x)
	}()

	if l.head >= len(l.input) {
		return "", EOF
	}

	pipeReg, _ := regexp.Compile("^\\s*(|>)")

	identReg, _ := regexp.Compile("^\\s*([A-z]+)(\\s*|$)")
	numberReg, _ := regexp.Compile("^\\s*([0-9]+(\\.[0-9]*){0,1})(\\s*|$)")
	flagReg, _ := regexp.Compile("^\\s*--([a-z]+)=")

	flagIdentValueReg, _ := regexp.Compile("^=([A-z]+)(\\s*|$)")
	flagNumberValueReg, _ := regexp.Compile("^=([0-9]+(\\.[0-9]*){0,1})(\\s*|$)")
	flagStringValueReg, _ := regexp.Compile("^=\"([A-z]+)\"(\\s*|$)")

	fmt.Printf("State: %v\n", l.input[l.head:])

	if x = pipeReg.FindString(l.input[l.head:]); x != "" {
		l.head += len(x)
		return x, IDENT
	}

	if x = identReg.FindString(l.input[l.head:]); x != "" {
		l.head += len(x)
		return x, IDENT
	}

	if x = numberReg.FindString(l.input[l.head:]); x != "" {
		l.head += len(x)
		return x, NUMBER
	}

	if x = flagReg.FindString(l.input[l.head:]); x != "" {
		l.head += len(x) - 1
		return x, FLAG
	}

	if x = flagIdentValueReg.FindString(l.input[l.head:]); x != "" {
		l.head += len(x)
		return x, FLAG_ARG_IDENT
	}

	if x = flagNumberValueReg.FindString(l.input[l.head:]); x != "" {
		l.head += len(x)
		return x, FLAG_ARG_NUMBER
	}

	if x = flagStringValueReg.FindString(l.input[l.head:]); x != "" {
		l.head += len(x)
		return x, FLAG_ARG_STRING
	}

	return "", UNKNOWN
}

func (l *CmdLexer) Unread(v string) {
	l.head -= len(v)
}

func NewFromArgs(args []string) *CmdLexer {
	con := strings.Join(args, " ")

	cl := &CmdLexer{
		input: con,
	}
	return cl
}

func NewFromText(args string) *CmdLexer {
	cl := &CmdLexer{
		input: args,
	}
	return cl
}

type CmdParser struct {
	Commands []Command
	*CmdLexer
}

func (p *CmdParser) ParseArgs(args []string) []Command {
	p.CmdLexer = NewFromArgs(args)

	for {
		v, t := p.Next()
		if t == EOF {
			return p.Commands
		}
		p.Commands = append(p.Commands, p.parseCommand(v, t))
	}
}

func (p *CmdParser) parseCommand(v0 string, t0 CmdToken) Command {
	cmd := &Command{}

	if t0 != IDENT {
		panic("expected ident")
	}

	cmd.Module = v0

	v1, t1 := p.Next()

	if t1 == EOF || t1 == PIPE {
		return *cmd
	}

	if t1 == IDENT {
		cmd.Func = v1
	} else {
		cmd.Args = append(cmd.Args, Argument{t1, v1})
	}

	for {
		vi, ti := p.Next()
		if ti == EOF || ti == PIPE {
			return *cmd
		}

		if ti == FLAG {
			vj, tj := p.Next()

			vi = vi[:len(vi)-1]
			vj = vj[1:]

			switch tj {
			case FLAG_ARG_STRING:
				cmd.AddStringFlag(vi, vj)
				continue
			case FLAG_ARG_NUMBER:
				cmd.AddNumberFlag(vi, vj)
				continue
			case FLAG_ARG_IDENT:
				cmd.AddIdentFlag(vi, vj)
				continue
			default:
				panic("its fucked")
			}
		}

		cmd.Args = append(cmd.Args, Argument{
			V: vi,
			T: ti,
		})

	}
}
