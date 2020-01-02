package main

//采用双栈结构，一个用来存储元素，另一个用来存储最小元素
type MinStack struct {
	stack []int
	min []int
	size int
}

/** initialize your data structure here. */
func ConstructorMinStack() MinStack {
	return MinStack{
		stack: []int{},
		min: []int{},
		size: 0,
	}
}

func (this *MinStack) Push(x int)  {
	this.stack = append(this.stack, x)
	n := len(this.min)
	if n == 0 || x <= this.min[n - 1]{
		//注意是小于等于，因为如果连续有两个相同的最小值入栈，只存一个的话后续出栈一个之后结果就不对了
		this.min = append(this.min, x)
	}
	this.size++
}

func (this *MinStack) Pop()  {
	n := len(this.min) - 1
	if this.Top() == this.min[n] {
		this.min = this.min[:n]
	}
	this.stack = this.stack[:this.size - 1]
	this.size--
}

func (this *MinStack) Top() int {
	if this.size == 0 {
		return 0
	}
	return this.stack[this.size - 1]
}

func (this *MinStack) GetMin() int {
	n := len(this.min)
	if n == 0 { //注意最小栈内的元素个数判断不能省略，否则对空栈调用GetMin就会越界
		return 0
	}
	return this.min[len(this.min) - 1]
}
