package main

import "dataStructure/list/doubleList"

func main() {
	// -------- singleList -------- //
	/*
	list := singleList.NewSingleList()
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
	// -------- doubleList -------- //
	/*
	list := doubleList.NewDoubleList()
	list.DumpList()
	list.DumpListReverse()
	list.AddHead(1)
	list.DumpList()
	list.DumpListReverse()
	list.DelTail()
	list.DumpList()
	list.DumpListReverse()
	list.AddTail(2)
	list.AddTail(3)
	list.DumpList()
	list.DumpListReverse()
	list.DelHead()
	list.DumpList()
	list.DumpListReverse()
	list.DelTail()
	list.DumpList()
	list.DumpListReverse()
	*/
	// -------- doubleListWithDummy -------- //
	list := doubleList.NewDoubleListWithDummy()
	list.DumpList()
	list.DumpListReverse()
	list.AddHead(1)
	list.DumpList()
	list.DumpListReverse()
	list.DelTail()
	list.DumpList()
	list.DumpListReverse()
	list.AddTail(2)
	list.AddTail(3)
	list.DumpList()
	list.DumpListReverse()
	list.DelHead()
	list.DumpList()
	list.DumpListReverse()
	list.DelTail()
	list.DumpList()
	list.DumpListReverse()
}
