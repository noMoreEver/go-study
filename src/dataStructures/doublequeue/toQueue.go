package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	next *Node
	prev *Node
	data int
}
type toQueue struct {
	length int
	start  *Node
	end    *Node
}

func main() {
	q := new(toQueue)
	q.push_first(3)
	q.push_first(4)
	q.push_first(5)
	q.push_last(6)
	q.push_last(7)
	q.push_last(8)

	fmt.Println("queue state: ", q)
	q.pop_first()
	fmt.Println("queue state: ", q)
	q.pop_first()
	fmt.Println("queue state: ", q)
	q.pop_last()
	fmt.Println("queue state: ", q)
	q.pop_last()
	fmt.Println("queue state: ", q)
	q.pop_last()
	fmt.Println("queue state: ", q)
	q.pop_last()
	fmt.Println("queue state: ", q)
	q.pop_last()
	fmt.Println("queue state: ", q)
	q.pop_last()
	fmt.Println("queue state: ", q)
}

func (q *toQueue) push_first(v int) {
	fmt.Println(q.start, q.end, q.length)

	n := &Node{data: v}

	if q.length == 0 {
		q.start = n
		q.end = n
	} else {
		n.next = q.start
		q.start.prev = n
		q.start = n
	}

	q.length += 1
}

func (q *toQueue) push_last(v int) {
	fmt.Println(q.start, q.end, q.length)

	n := &Node{data: v}

	if q.length == 0 {
		q.start = n
		q.end = n
	} else {
		q.end.next = n
		n.prev = q.end
		q.end = n
	}
	q.length += 1
}

func (q *toQueue) pop_first() int {
	if q.length == 0 {
		return 0
	}
	fmt.Println(q.start, q.end, q.length)

	n := q.start
	q.start = q.start.next
	if q.start != nil {
		q.start.prev = nil
	}
	n.next = nil

	q.length -= 1
	return n.data
}

func (q *toQueue) pop_last() int {
	if q.length == 0 {
		return 0
	}
	fmt.Println(q.start, q.end, q.length)
	n := q.end
	q.end = q.end.prev
	if q.end != nil {
		q.end.next = nil
	}
	n.prev = nil

	q.length -= 1
	return n.data
}

func (q *toQueue) Length() int {
	return q.length
}

func (q *toQueue) String() string {
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
