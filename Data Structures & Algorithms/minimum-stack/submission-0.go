type MinStack struct {
	stack []int
	min int
}

func Constructor() MinStack {
	return MinStack {
		stack: make([]int, 0),
		min: math.MaxInt32,
	}
}

func (this *MinStack) Push(val int) {
	if len(this.stack) == 0 {
		this.stack = append(this.stack, 0)
		this.min = val
	} else {
		this.stack = append(this.stack, val - this.min)
		if val < this.min {
			this.min = val
		}
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	pop := this.stack[len(this.stack) - 1]
	this.stack = this.stack[:len(this.stack) - 1]
	if pop < 0 {
		this.min = this.min - pop
	}
}

func (this *MinStack) Top() int {
	top := this.stack[len(this.stack) - 1]
	if top > 0 {
		return this.min + top
	}
	return this.min
}

func (this *MinStack) GetMin() int {
	return this.min
}
