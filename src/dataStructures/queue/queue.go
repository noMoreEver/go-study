package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	data int
	next *Node
}
type Queue struct {
	start  *Node
	end    *Node
	length int
}

func main() {
	q := new(Queue)
	q.push(3)
	q.push(4)
	q.push(5)
	fmt.Println("peek: ", q.peek())
	fmt.Println("queue state: ", q)
	q.pop()
	fmt.Println("queue state: ", q)
	fmt.Println("peek: ", q.peek())
	q.pop()
	fmt.Println("queue state: ", q)
	q.pop()
	fmt.Println("queue state: ", q)
	q.pop()
	fmt.Println("queue state: ", q)
}

func (q *Queue) push(v int) {
	if q.length == 0 {
		q.start = &Node{data: v}
		q.end = q.start
	} else {
		n := &Node{data: v}
		q.end.next = n
		q.end = n
	}
	q.length += 1
}

func (q *Queue) pop() int {
	if q.length == 0 {
		return 0
	} else {
		v := q.start.data
		q.start = q.start.next
		q.length -= 1
		return v
	}
}

func (q *Queue) Length() int {
	return q.length
}

func (q *Queue) peek() int {
	if q.length == 0 {
		return 0
	} else {
		return q.start.data
	}
}

func (q *Queue) String() string {
	if q.length == 0 {
		return "[]"
	}
	res := "["
	n := q.start
	for {
		if n == nil {
			break
		}
		res += strconv.Itoa(n.data) + "->"
		n = n.next
	}
	return res + "]"
}
