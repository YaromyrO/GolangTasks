package queue

type Queue struct {
	rear *node
	front *node
	count int
}

type node struct {
	value string
	previous *node
}

func NewQueue() *Queue {
	queue := new(Queue)
	return queue
}

func (queue *Queue) Enqueue(value string) {

	node := new(node)
	node.value = value
	node.previous = nil
	if queue.count == 0 {
		queue.front = node
		queue.rear = node
		queue.count++
	} else {
		queue.rear.previous = node
		queue.rear = node
		queue.count++
	}
}

func (queue Queue) Size() int {

	if queue.count != 0 {
		return queue.count
	}

	return 0
}

func (queue *Queue) Dequeue() string {

	result := queue.front.value
	if queue.front.previous == nil {
		queue.rear = nil
		queue.front = nil
		queue.count--
	} else {
		queue.front = queue.front.previous
		queue.count--
	}

	return result
}

func (queue Queue) Peek() string {
	return queue.front.value
}

func (queue Queue) Contains(value string) bool {

	if queue.front.value == value {
		return true
	}
	for i := queue.count; i > 0; i-- {
		if queue.front.previous.value == value {
			return true
		}
	}

	return false
}

func (queue Queue) IsEmpty() bool {

	if queue.count == 0 {
		return true
	}

	return false
}
