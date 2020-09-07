package doubleList

import "fmt"

// -------- doubleList -------- //
type DoubleList struct {
	head *doubleListNode
	tail *doubleListNode
}

type doubleListNode struct {
	val int
	prev *doubleListNode
	next *doubleListNode
}

func NewDoubleList() *DoubleList {
	return &DoubleList{}
}

func (doubleList *DoubleList) DumpList() {
	fmt.Print("list: ")
	for cur := doubleList.head; cur != nil; cur = cur.next {
		fmt.Print(cur.val)
		if cur.next != nil {
			fmt.Print("->")
		}
	}
	fmt.Println()
}

func (doubleList *DoubleList) DumpListReverse() {
	fmt.Print("list reverse : ")
	for cur := doubleList.tail; cur != nil; cur = cur.prev {
		fmt.Print(cur.val)
		if cur.prev != nil {
			fmt.Print("<-")
		}
	}
	fmt.Println()
}

func (doubleList *DoubleList) AddHead(val int) {
	newHead := &doubleListNode{val: val, next: doubleList.head}
	doubleList.head = newHead
	if doubleList.tail == nil {
		doubleList.tail = newHead
	} else {
		newHead.next.prev = newHead
	}
}

func (doubleList *DoubleList) AddTail(val int) {
	newTail := &doubleListNode{val: val, prev: doubleList.tail}
	if doubleList.tail == nil {
		doubleList.tail = newTail
		doubleList.head = newTail
	} else {
		doubleList.tail.next = newTail
	}
	doubleList.tail = newTail
}

func (doubleList *DoubleList) DelHead() (bool, int) {
	if doubleList.head == nil {
		return false, -1
	}
	oldHead := doubleList.head
	doubleList.head = doubleList.head.next
	if doubleList.head == nil {
		doubleList.tail = nil
	} else {
		doubleList.head.prev = nil
	}
	return true, oldHead.val
}

func (doubleList *DoubleList) DelTail() (bool, int) {
	if doubleList.tail == nil {
		return false, -1
	}
	oldTail := doubleList.tail
	doubleList.tail = doubleList.tail.prev
	if doubleList.tail == nil {
		doubleList.head = nil
	} else {
		doubleList.tail.next = nil
	}
	return true, oldTail.val
}






