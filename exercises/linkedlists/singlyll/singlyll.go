package main

import (
	"fmt"
)

// define a head pointer
type Node struct {
	data int
	next *Node
}

// define the linked list
type LinkedList struct {
	head *Node
	//length int // define the length of the linked list
}

func (l *LinkedList) Insert(data int) {
	newNode := &Node{data: data}

	if l.head == nil {
		l.head = newNode
	} else {

		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}

	// for func (list *LinkedList) prepend(next *Node){}
	// second := l.head
	// l.head = n
	// l.head.next = second
	// l.length++

}

func (l *LinkedList) Display() {

	current := l.head

	if current == nil {
		fmt.Println("List is empty")
	}
	for current != nil {
		fmt.Printf("%d", current.data)
		current = current.next
	}
}

func main() {

	l := LinkedList{}

	l.Insert(10)
	// node1 := &Node{data: 48}
	// node2 := &Node{data: 45}
	// node3 := &Node{data: 43}
	// mylist.prepend(node1)
	// mylist.prepend(node2)
	// mylist.prepend(node3)
	// fmt.Print(mylist)
	l.Display()
	fmt.Println()
}
