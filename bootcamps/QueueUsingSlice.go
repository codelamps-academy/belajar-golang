package bootcamps

import "fmt"

type Queue2 struct {
	items []int
}

func (q *Queue2) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *Queue2) Dequeu() int {
	if len(q.items) == 0 {
		return -1 // QUEUE KOSONG
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func main() {
	q := Queue2{}

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	fmt.Println(q.Dequeu())
	fmt.Println(q.Dequeu())
	fmt.Println(q.Dequeu())
	fmt.Println(q.Dequeu())
}
