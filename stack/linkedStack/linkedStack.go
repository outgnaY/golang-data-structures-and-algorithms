package linkedStack

import "fmt"

// -------- linkedStack -------- //
type LinkedStack struct {
	top *linkedStackNode
	size int
}

type linkedStackNode struct {
	val int
	next *linkedStackNode
}

func NewLinkedStack() *LinkedStack {
	return &LinkedStack{}
}

func (linkedStack *LinkedStack) DumpStack() {
	fmt.Print("stack: ")
	for cur := linkedStack.top; cur != nil; cur = cur.next {
		fmt.Print(cur.val)
		if cur.next != nil {
			fmt.Print("->")
		}
	}
	fmt.Println()
}

func (linkedStack *LinkedStack) Push(val int) {
	node := &linkedStackNode{val: val, next: linkedStack.top}
	linkedStack.top = node
	linkedStack.size += 1
}

func (linkedStack *LinkedStack) Pop() (bool, int) {
	if linkedStack.size == 0 {
		return false, -1
	}
	oldTop := linkedStack.top
	linkedStack.top = linkedStack.top.next
	linkedStack.size -= 1
	return true, oldTop.val
}

func (linkedStack *LinkedStack) Empty() bool {
	return linkedStack.size == 0
}

func (linkedStack *LinkedStack) Peek() (bool, int) {
	if linkedStack.size == 0 {
		return false, -1
	}
	return true, linkedStack.top.val
}
