package parsersimple

import (
	l "arit/lexer_simple"
)

type Exp interface { ToString() string }
type NilExp struct {}

func (e *NilExp) ToString() string { return "nilExp" }  

type ValueExp interface {}
type NumberExp interface {}

type IntExp struct { pos l.Position; s string }
type FloatExp struct { pos l.Position; s string }

func (e *IntExp) ToString() string { return e.s  }  
func (e *FloatExp) ToString() string { return e.s  }  

type StringExp struct { pos l.Position; s string }
func (e *StringExp) ToString() string { return e.s  }  

type BinExp struct {
  pos l.Position
  e1,e2 Exp
  op Op 
}

func (e *BinExp) ToString() string { return e.e1.ToString() +  e.op.ToOpString() + e.e2.ToString() } 
