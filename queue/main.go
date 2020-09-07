package main

import "dataStructure/queue/linkedQueue"

func main() {
	q := linkedQueue.NewLinkedQueue()
	q.DumpQueue()
	q.EnQueue(4)
	q.DumpQueue()
	q.EnQueue(2)
	q.EnQueue(3)
	q.DumpQueue()
	q.DeQueue()
	q.DumpQueue()
	q.DeQueue()
	q.DeQueue()
	q.DumpQueue()
	q.DeQueue()
	q.DumpQueue()
}