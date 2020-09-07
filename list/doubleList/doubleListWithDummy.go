package doubleList

import "fmt"

// -------- doubleListWithDummy -------- //
type DoubleListWithDummy struct {
	dummyHead *doubleListWithDummyNode
	dummyTail *doubleListWithDummyNode
}

type doubleListWithDummyNode struct {
	val int
	prev *doubleListWithDummyNode
	next *doubleListWithDummyNode
}

func NewDoubleListWithDummy() *DoubleListWithDummy {
	dummyHead := &doubleListWithDummyNode{}
	dummyTail := &doubleListWithDummyNode{}
	dummyHead.next = dummyTail
	dummyTail.prev = dummyHead
	return &DoubleListWithDummy{dummyHead: dummyHead, dummyTail: dummyTail}
}

func (doubleListWithDummy *DoubleListWithDummy) DumpList() {
	fmt.Print("list: ")
	for cur := doubleListWithDummy.dummyHead.next; cur != doubleListWithDummy.dummyTail; cur = cur.next {
		fmt.Print(cur.val)
		if cur.next != doubleListWithDummy.dummyTail {
			fmt.Print("->")
		}
	}
	fmt.Println()
}

func (doubleListWithDummy *DoubleListWithDummy) DumpListReverse() {
	fmt.Print("list reverse: ")
	for cur := doubleListWithDummy.dummyTail.prev; cur != doubleListWithDummy.dummyHead; cur = cur.prev {
		fmt.Print(cur.val)
		if cur.prev != doubleListWithDummy.dummyHead {
			fmt.Print("<-")
		}
	}
	fmt.Println()
}

func (doubleListWithDummy *DoubleListWithDummy) AddHead(val int) {
	newHead := &doubleListWithDummyNode{val: val, next: doubleListWithDummy.dummyHead.next, prev: doubleListWithDummy.dummyHead}
	doubleListWithDummy.dummyHead.next.prev = newHead
	doubleListWithDummy.dummyHead.next = newHead
}

func (doubleListWithDummy *DoubleListWithDummy) AddTail(val int) {
	newTail := &doubleListWithDummyNode{val: val, next: doubleListWithDummy.dummyTail, prev: doubleListWithDummy.dummyTail.prev}
	doubleListWithDummy.dummyTail.prev.next = newTail
	doubleListWithDummy.dummyTail.prev = newTail
}

func (doubleListWithDummy *DoubleListWithDummy) DelHead() (bool, int) {
	if doubleListWithDummy.dummyHead.next == doubleListWithDummy.dummyTail {
		return false, -1
	}
	oldHead := doubleListWithDummy.dummyHead.next
	oldHead.next.prev = doubleListWithDummy.dummyHead
	doubleListWithDummy.dummyHead.next = oldHead.next
	oldHead.prev = nil
	oldHead.next = nil
	return true, oldHead.val
}

func (doubleListWithDummy *DoubleListWithDummy) DelTail() (bool, int) {
	if doubleListWithDummy.dummyHead.next == doubleListWithDummy.dummyTail {
		return false, -1
	}
	oldTail := doubleListWithDummy.dummyTail.prev
	oldTail.prev.next = doubleListWithDummy.dummyTail
	doubleListWithDummy.dummyTail.prev = oldTail.prev
	oldTail.prev = nil
	oldTail.next = nil
	return true, oldTail.val
}