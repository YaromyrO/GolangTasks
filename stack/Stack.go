package main

type Stack struct {
	top *node
	count int
}
type node struct {
	value float64
	next *node
}

func NewStack() *Stack {
	stack := new(Stack)
	return stack
}

func (stack *Stack) Push(value float64) {

	node := new(node)
	node.value = value
	node.next = nil
	if stack.count == 0 {
		stack.top = node
		stack.count++
	} else {
		node.next = stack.top
		stack.top = node
		stack.count++
	}
}

func (stack Stack) Size() int {

	if stack.top != nil {
		return stack.count
	}

	return 0
}

func (stack *Stack) Pop() float64 {

	result := stack.top.value
	if stack.top.next == nil {
		stack.top = nil
		stack.count--
	} else {
		stack.top = stack.top.next
		stack.count--
	}

	return result
}

func (stack Stack) Peek() float64 {

	if stack.top != nil {
		return stack.top.value
	}

	return 0.0
}

func (stack Stack) Contains(value float64) bool {

	if stack.top.value == value {
		return true
	}
	for i := 0; i < stack.count; i++ {
		if stack.top.next.value == value {
			return true
		}
	}

	return false
}

func (stack Stack) IsEmpty() bool {

	if stack.count == 0 {
		return true
	}

	return false
}

