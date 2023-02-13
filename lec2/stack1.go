package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

type Stack struct {
	head *Node
}

func (s *Stack) pop() *Node {
	temp := s.head
	if s.head != nil {
		if s.head.next == nil {
			s.head = nil
		} else {
			s.head = s.head.next
		}
	}
	return temp
}

func (s *Stack) peek() int {
	return s.head.val
}

func (s *Stack) push(val int) {
	node := &Node{val: val, next: nil}
	if s.head == nil {
		s.head = node
	} else {
		node.next = s.head
		s.head = node
	}
}

func (s *Stack) clear() {
	s.head = nil
}

func (s *Stack) contains(val int) {
	temp := s.head
	for temp != nil {
		if temp.val == val {
			fmt.Println("Yes")
			return
		}
		temp = temp.next
	}
	fmt.Println("No")
}

func (s *Stack) increment(val int) {
	cur := s.head
	for cur != nil {
		cur.val += val
		cur = cur.next
	}
}

func (s *Stack) print() {
	cur := s.head
	for cur != nil {
		fmt.Println(cur.val)
		cur = cur.next
	}
}

func (s *Stack) printReverse() {
	ar := make([]int, 0)
	cur := s.head
	for cur != nil {
		ar = append(ar, cur.val)
		cur = cur.next
	}
	for i := len(ar) - 1; i >= 0; i-- {
		fmt.Println(ar[i])
	}
}

func main() {
	s := Stack{head: nil}
	s.push(1)
	s.push(2)
	s.push(3)
	// s.pop()
	// // fmt.Println(s.peek())
	// s.print()
	// s.contains(3)
	s.printReverse()
	// s.clear()
	// s.print()
}
