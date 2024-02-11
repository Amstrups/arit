package parsersimple

import (
	l "arit/lexer_simple"
)

type Op interface { ToOpString() string }

type EqOp struct { pos l.Position; }
type NeqOp struct { pos l.Position; }
type LtOp struct { pos l.Position; }
type LeOp struct { pos l.Position; }
type GtOp struct { pos l.Position; }
type GeOp struct { pos l.Position; }
type PlusOp struct { pos l.Position; }
type MinusOp struct { pos l.Position; }
type TimesOp struct { pos l.Position; }
type DivideOp struct { pos l.Position; }
type ExponentOp struct { pos l.Position; }

func (o *EqOp) ToOpString() string { return "==" }
func (o *NeqOp) ToOpString() string { return "!=" }
func (o *LtOp) ToOpString() string { return "<" }
func (o *LeOp) ToOpString() string { return "<=" }
func (o *GtOp) ToOpString() string { return ">" }
func (o *GeOp) ToOpString() string { return ">=" }
func (o *PlusOp) ToOpString() string { return "+" }
func (o *MinusOp) ToOpString() string { return "-" }
func (o *TimesOp) ToOpString() string { return "*" }
func (o *DivideOp) ToOpString() string { return "/" }
func (o *ExponentOp) ToOpString() string { return "^" }

