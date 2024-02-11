package parsersimple


type (
	stack struct {
		top *state
		length int
	}

	state struct {
		val int
		prev *state
	}	
)

func NewStack(first int) *stack {
  s := &stack{nil, 0}
  s.push(0)
  return s
}

func (s *stack) push(v int) {
  t := s.top
  s.top = &state{v, t}
}

func (s *stack) pop() state {
  s.length--
  t := *s.top
  s.top = t.prev
  return t
}

func (s *stack) peek() int {
  return s.top.val
}
