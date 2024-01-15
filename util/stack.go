package util

import "errors"

type IntStack []int
var ErrStackEmpty = errors.New("stack is empty")

func (s IntStack) Push(i int) IntStack { 
  return append(s, i)
}
func (s IntStack) Pop() (IntStack,int,error) { 
  l := len(s)
  if (l == 0) { 
    return s,-1,ErrStackEmpty 
  }
  return s[:l-1], s[l-1], nil
}

func (s IntStack) IsEmpty() bool { 
  return len(s) == 0
}
