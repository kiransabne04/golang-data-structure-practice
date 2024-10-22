package singlelylinkedlist

type Node[T any] struct {
	data T
	next *Node[T]
}

type Slist[T any] struct {
	head *Node[T]
	size int
}
