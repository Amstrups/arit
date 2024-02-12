package parsersimple

type expr interface { String() string }

type term struct {
  name string
  value string
}

func (t *term) String() string { 
  return "[" + t.name + " : " + t.value + "]" 
}

type nonterm struct {
  name string
  subExps []expr
}

func (n *nonterm) String() string {
  out := "[" + n.name + " : \n" 
  for _,e := range n.subExps {
    out += e.String()
  }
  return out
}

