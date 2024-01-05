package main

import (
	"fmt"
)

const MAXSIZE = 10

type stackI interface {
	stackcount() int
	push(i int)
	pop() int
	isEmpty() bool
	show()
}

type queueI interface {
	enqueue(i int)
	dequeue() int
	queuecount() int
	isEmpty() bool
	show()
}

type stack struct {
	sd []int
}

type queue struct {
	qd []int
}

func main() {
	var sIf stackI
	si := stack{}
	sIf = &si
	sIf.push(10)
	sIf.push(20)
	sIf.show()
	sIf.pop()
	sIf.show()

	var qIf queueI = &queue{}

	qIf.enqueue(100)
	qIf.enqueue(200)
	qIf.show()
	qIf.dequeue()
	qIf.show()
}

func (s *stack) show() {
	for i := range s.sd {
		fmt.Println(s.sd[i])
	}
}

func (s *stack) isEmpty() bool {
	return len(s.sd) == 0
}

func (s *stack) stackcount() int {
	return len(s.sd)
}

func (s *stack) push(i int) {
	if s.isEmpty() {
		s.sd = append(s.sd, i)
	} else if s.stackcount() < MAXSIZE {
		s.sd = append(s.sd, i)
	} else {
		fmt.Println("Stack is full")
	}
}

func (s *stack) pop() int {
	if !s.isEmpty() {
		d := s.sd[len(s.sd)-1]
		s.sd = s.sd[:len(s.sd)-1] //s.sd = slices.Delete(s.sd, len(s.sd)-1, len(s.sd)) alternate but uses slices std library
		return d
	}
	fmt.Println("Stack is empty")
	return -1
}

func (q *queue) enqueue(i int) {
	q.qd = append(q.qd, i)
}

func (q *queue) dequeue() int {
	if len(q.qd) == 0 {
		fmt.Println("Queue is empty")
		return -1
	}
	d := q.qd[0]
	q.qd = q.qd[1:]
	return d
}

func (q *queue) queuecount() int {
	return len(q.qd)
}

func (q *queue) isEmpty() bool {
	return len(q.qd) == 0
}

func (q *queue) show() {
	for i := range q.qd {
		fmt.Println(q.qd[i])
	}
}
