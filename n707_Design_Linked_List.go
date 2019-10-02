package main

type node struct {
	val int
	next *node
}

type MyLinkedList struct {
	head, tail *node
	size int
}


/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	t := &node{}
	h := &node{next: t}
	return MyLinkedList {
		head: h,
		tail: t,
		size: 0,
	}
}


/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || this.size <= index {
		return -1
	}

	i, cur := 0, this.head.next
	for i < index {
		i++
		cur = cur.next
	}
	return cur.val
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
	newNode := &node{val: val}
	newNode.next = this.head.next
	this.head.next = newNode
	this.size++
}


/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
	this.tail.val = val
	newNode := &node{}
	this.tail.next = newNode
	this.tail = newNode
	this.size++
}


/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	switch {
	case index < 0 || this.size < index:
		return
	case index == 0:
		this.AddAtHead(val)
		return
	case index == this.size:
		this.AddAtTail(val)
		return
	}

	i, cur := -1, this.head
	for i + 1 < index {
		i++
		cur = cur.next
	}

	newNode := &node{val: val}
	newNode.next = cur.next
	cur.next = newNode
	this.size++
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
	if index < 0 ||this.size <= index {
		return
	}

	i, cur := -1, this.head
	for i + 1 < index {
		i++
		cur = cur.next
	}

	cur.next = cur.next.next
	this.size--
}

