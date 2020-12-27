package queue

// Queue 利用数组实现基本的FIFO队列（环形队列）
type Queue struct {
	maxSize int   //队列数字长度
	Q       []int //
	head    int
	tail    int
}

// NewQueue 返回一个新的队列
func NewQueue(cap int) *Queue {
	q := &Queue{maxSize: cap + 1}
	q.Q = make([]int, q.maxSize)
	q.head = 0
	q.tail = 0
	return q
}

func (q *Queue) isFull() bool {
	return q.head == ((q.tail + 1) % q.maxSize)
}

func (q *Queue) isEmpty() bool {
	return q.head == q.tail
}

func (q *Queue) enqueue(v int) {
	if q.isFull() {
		panic("queue is full")
	}
	q.Q[q.tail] = v
	if q.tail+1 == q.maxSize {
		q.tail = 0
	} else {
		q.tail++
	}
}

func (q *Queue) dequeue() int {
	if q.isEmpty() {
		panic("queue is empty")
	}
	v := q.Q[q.head]
	if q.head+1 == q.maxSize {
		q.head = 0
	} else {
		q.head++
	}
	return v
}
