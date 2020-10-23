package structures

import "fmt"

// Dynamic Queue

type Queue struct {
	first *Node
	last  *Node
}

func (q *Queue) Insert(value interface{}) {
	if q.first == nil {
		q.first = &Node{value, nil}
		q.last = q.first
	} else {
		q.last.next = &Node{value, nil}
		q.last = q.last.next
	}
}

func (q *Queue) Remove() interface{} {
	if q.first == nil {
		return nil
	} else {
		value := q.first.value
		q.first = q.first.next
		return value
	}
}

func (q *Queue) Head() interface{} {
	return q.first.value
}

func (q *Queue) Count() int {
	curr := q.first
	count := 0
	for curr != nil {
		count++
		curr = curr.next
	}
	return count
}

func (q *Queue) String() string {
	curr := q.first
	output := ""
	for curr != nil {
		output += fmt.Sprintf("%v <- ", curr.value)
		curr = curr.next
	}
	return output
}
