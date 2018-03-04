package adt

type Stack struct {
	top *element
	Size int
}

type element struct {
	value interface{}
	prev *element
}

func (s *Stack) Push(value interface{}) {

	s.top = &element{value, s.top}
	s.Size++
}

func (s *Stack) Pop() (value interface{}) {

	if s.Size == 0 {
		panic("Stack underflowed.")
	}

	value = s.top.value
	s.top = s.top.prev
	s.Size--

	return value
}

func (s *Stack) Top() interface{} {

	if s.Size == 0 {
		return nil
	}

	return s.top.value
}
