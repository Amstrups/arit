package parser
type Exp interface {}

type IntExp struct {
  val int
}
type FloatExp struct {}

type BinExp struct { 
  op int
  left Exp
  right Exp
}

/*
func (bin BinExp) eval() float64 { 


  return 0.0
}
*/
