package bootcamps

import "fmt"

type Node struct {
	value int
	next  *Node
}

type Queue struct {
	head *Node
	tail *Node
}

func (q *Queue) Enqueue(item int) {
	newNode := &Node{value: item, next: nil}
	if q.tail == nil {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
}

func (q *Queue) Dequeue() int {
	if q.head == nil {
		return -1 // QUEUE KOSONG
	}
	item := q.head.value
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	return item
}

func main() {
	q := Queue{}

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
}
