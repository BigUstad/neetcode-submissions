type MinStack struct {
	curstack *list.List
	minstack *list.List
}

func Constructor() MinStack {
	c := list.New()
	m := list.New()
	return MinStack{
		curstack: c,
		minstack: m,
	}
}

func (this *MinStack) Push(val int) {
	if this.curstack.Len() == 0 {
		this.curstack.PushFront(val)
		this.minstack.PushFront(val)
		return
	}
	minTop := this.minstack.Front().Value.(int)
	min := min(minTop, val)
	this.curstack.PushFront(val)
	this.minstack.PushFront(min)
}

func (this *MinStack) Pop() {
	curFront := this.curstack.Front()
	minFront := this.minstack.Front()
	this.curstack.Remove(curFront)
	this.minstack.Remove(minFront)
}

func (this *MinStack) Top() int {
	return this.curstack.Front().Value.(int)
}

func (this *MinStack) GetMin() int {
	return this.minstack.Front().Value.(int)
}
