package main

import "dataStructure/stack/linkedStack"

func main() {
	s := linkedStack.NewLinkedStack()
	s.DumpStack()
	s.Push(1)
	s.DumpStack()
	s.Pop()
	s.Pop()
	s.DumpStack()
	s.Push(4)
	s.Push(5)
	s.Push(6)
	s.DumpStack()
	s.Pop()
	s.DumpStack()
	s.Pop()
	s.Pop()
	s.DumpStack()
	s.Pop()
	s.DumpStack()
}
