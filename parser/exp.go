package parser

import (
  "errors"
  "fmt"
  "log"
  "strconv"
)

var ErrDivideByZero = errors.New("divide by zero")
var ErrUnknownOperation = errors.New("unknown binop")

func getNonNilErr(l_err error, r_err error) error {
  if (l_err != nil) { 
    return l_err
  }
  return r_err
}

type Exp interface {
  Eval() (float32, error)
  String() string
}

type NegExp struct {
  val Exp
  pos, line int
}

type IntExp struct {
  val int 
  pos, line int 
  
}

func (in IntExp) Eval() (float32,error) { 
  return float32(in.val), nil
}

func (in IntExp) String() string {
  return fmt.Sprintf("IntExp(%d,%d:%d)", in.val, in.pos, in.line)
}

type FloatExp struct {
  val float32
  pos, line int 
}

func (fl FloatExp) Eval() (float32,error) { 
  return fl.val,nil 
}
func (fl FloatExp) String() string {
  return fmt.Sprintf("FloatExp(%.4f,%d:%d)", fl.val, fl.pos, fl.line)
}

type BinExp struct { 
  op int
  left Exp
  right Exp
  pos, line int
}

func (bin BinExp) extract() (float32, float32, error) {
  l, lerr := bin.left.Eval()
  r, rerr := bin.right.Eval()
  err := getNonNilErr(lerr, rerr)
  return l,r,err
}

func (bin BinExp) Eval() (float32,error) { 
  switch {
  case bin.op == ADD:
    l,r,err := bin.extract()
    if (err != nil) { 
      return -1, err
    }
    return l + r,nil 

  case bin.op == SUB:
    l,r,err := bin.extract()
    if (err != nil) { 
      return -1, err
    }
    return l - r,nil 

  case bin.op == DIV:
    l,r,err := bin.extract()
    if (err != nil) { 
      return -1, err
    }
    if (r == 0) { 
      return -1,ErrDivideByZero 
    } 
    return l / r, err
  case bin.op == MUL:
    l,r,err := bin.extract()
    if (err != nil) { 
      return -1, err
    }
    return l * r, err
  }

  return -1,ErrUnknownOperation
}

func (bin BinExp) String() string {
  return fmt.Sprintf("BinExp(%d,l:%s,r:%s,%d:%d)", bin.op, bin.left, bin.right, bin.pos, bin.line)
}


func NewIntExp(t Token) IntExp { 
  intVal, err := strconv.Atoi(t.val)
  if err != nil { 
    log.Fatal(err)
  }
  return IntExp{intVal, t.pos, t.line}
}


func NewFloatExp(t Token) FloatExp { 
  floatVal, err := strconv.ParseFloat(t.val, 64)
  if err != nil { 
    log.Fatal(err)
  }
  return FloatExp{float32(floatVal), t.pos, t.line}
}

func dynamicTermExp(t Token) Exp { 
  if (t.tokent == 1) { 
    return NewIntExp(t)
  } else {
    return NewFloatExp(t)
  }
}


func NewOpExp(tokens []Token) BinExp {
  t1,op, t2 := tokens[0], tokens[1], tokens[2]
  e1 := dynamicTermExp(t1)
  e2 := dynamicTermExp(t2)
  return BinExp{op.tokent, e1, e2, t1.pos, t1.line}
}
