package structures

type Queue struct {
	Queue []interface{}
}

func (q *Queue) Enqueue(value interface{}) {
	q.Queue = append(q.Queue, value)
}
func (q *Queue) Dequeue() {
	q.Queue = q.Queue[1:]
}
func (q *Queue) Peek() interface{} {
	return q.Queue[0]
}
func (q *Queue) Size() int {
	return len(q.Queue)
}
func (q *Queue) IsEmpty() bool {
	return len(q.Queue) == 0
}
