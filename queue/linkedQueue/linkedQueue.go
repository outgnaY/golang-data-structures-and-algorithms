package linkedQueue

import "fmt"

// -------- linkedQueue -------- //
type LinkedQueue struct {
	head *linkedQueueNode
	tail *linkedQueueNode
	size int
}

type linkedQueueNode struct {
	val int
	next *linkedQueueNode
}

func NewLinkedQueue() *LinkedQueue {
	return &LinkedQueue{}
}

func (linkedQueue *LinkedQueue) DumpQueue() {
	fmt.Print("queue: ")
	for cur := linkedQueue.head; cur != nil; cur = cur.next {
		fmt.Print(cur.val)
		if cur.next != nil {
			fmt.Print("->")
		}
	}
	fmt.Println()
}

func (linkedQueue *LinkedQueue) EnQueue(val int) {
	node := &linkedQueueNode{val: val}
	if linkedQueue.size == 0 {
		linkedQueue.head = node
	} else {
		linkedQueue.tail.next = node
	}
	linkedQueue.tail = node
	linkedQueue.size += 1
}

func (linkedQueue *LinkedQueue) DeQueue() (bool, int) {
	if linkedQueue.size == 0 {
		return false, -1
	}
	linkedQueue.size -= 1
	oldHead := linkedQueue.head
	linkedQueue.head = linkedQueue.head.next
	if linkedQueue.size == 0 {
		linkedQueue.tail = nil
	}
	return true, oldHead.val
}

func (linkedQueue *LinkedQueue) Empty() bool {
	return linkedQueue.size == 0
}
