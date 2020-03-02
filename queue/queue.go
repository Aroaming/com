package queue

const QueueMinSize = 32

type Queue struct {
	item              []interface{}
	front, rear, size int
}

func New() *Queue {
	return &Queue{
		item: make([]interface{}, QueueMinSize),
	}
}

func (q *Queue) Put(item interface{}) {
	if q.size == len(q.item) {
		q.resize()
	}
	q.item[q.rear+1] = item //入队
	q.size++
	q.rear = q.rear + 1
}
func (q *Queue) Cut() interface{} {
	if q.size <= 0 {
		panic("queue:cut called on empty queue")
	}
	ret := q.item[q.front]
	q.item[q.front] = nil
	q.front = q.front + 1
	q.size--

	return ret //from front
}
func (q *Queue) Size() int {
	return q.size
}
func (q *Queue) Get(id int) interface{} {
	if id < 0 {
		id = id + q.size
	}
	if id < 0 || id > q.size {
		panic("queue: get id range out of queue size")
	}
	return q.item[q.front+id]
}

func (q *Queue) Clear() {
	q.size, q.front, q.rear = 0, 0, 0
	q.item = make([]interface{}, QueueMinSize)
}

func (q *Queue) resize() {
	newBuf := make([]interface{}, q.size<<1)
	if q.front < q.rear {
		copy(newBuf, q.item[q.front:q.rear])
	} else {
		n := copy(newBuf, q.item[q.front:])
		copy(newBuf[n:], q.item[:q.rear])
	}
	q.front = 0
	q.rear = q.size
	q.item = newBuf
}
