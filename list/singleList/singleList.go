package singleList

import "fmt"

// -------- singleList -------- //
type SingleList struct {
	head *singleListNode
	tail *singleListNode
}

type singleListNode struct {
	val int
	next *singleListNode
}

func NewSingleList() *SingleList {
	return &SingleList{}
}

func (singleList *SingleList) DumpList() {
	fmt.Print("list: ")
	for cur := singleList.head; cur != nil; cur = cur.next {
		fmt.Print(cur.val)
		if cur.next != nil {
			fmt.Print("->")
		}
	}
	fmt.Println()
}

func (singleList *SingleList) AddHead(val int) {
	singleList.head = &singleListNode{val: val, next: singleList.head}
	if singleList.tail == nil {
		singleList.tail = singleList.head
	}
}

func (singleList *SingleList) AddTail(val int) {
	newTail := &singleListNode{val: val}
	if singleList.tail == nil {
		singleList.head = newTail
	} else {
		singleList.tail.next = newTail
	}
	singleList.tail = newTail
}

func (singleList *SingleList) DelHead() (bool, int) {
	if singleList.head == nil {
		return false, -1
	}
	oldHead := singleList.head
	singleList.head = singleList.head.next
	if singleList.head == nil {
		singleList.tail = nil
	}
	return true, oldHead.val
}

// 删除尾结点O(n)，不应使用单链表实现
/**
	list := NewSingleList()
	list.DumpList()
	list.AddHead(5)
	list.DumpList()
	list.AddTail(4)
	list.DumpList()
	list.DelHead()
	list.DumpList()
	list.DelHead()
	list.DumpList()
*/

