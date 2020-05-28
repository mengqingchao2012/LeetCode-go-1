package main

type MyStack struct {
	enqueue []int
	dequeue []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{[]int{}, []int{}}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.enqueue = append(this.enqueue, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	size := len(this.enqueue)
	this.dequeue = append(this.dequeue, this.enqueue[:size-1]...)
	this.enqueue = this.enqueue[size-1:]
	item := this.enqueue[0]
	this.enqueue = this.dequeue
	this.dequeue = []int{}
	return item
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.enqueue[len(this.enqueue)-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.enqueue) == 0
}
