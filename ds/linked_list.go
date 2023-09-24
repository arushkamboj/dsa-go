package ds

import (
	"fmt"
	"math/rand"
)

type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	head *Node
}

// Insert adds the provided element at the end of the LinkedList
func (list *LinkedList) Insert(val int) {
	newNode := &Node{val: val, next: nil}

	// check if the head is nil, this would mean the list is empty, therefore, make the new node the head
	if list.head == nil {
		list.head = newNode
		return
	}

	// if head was not nil, it means there were elements already present in the list
	// go to the end of the list, add the node at the end
	temp := list.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode
}

// Iterate traverses the list and prints each element
func (list *LinkedList) Iterate() {
	temp := list.head
	for temp != nil {
		fmt.Printf("-> %v ", temp.val)
		temp = temp.next
	}
}

func ImplLinkedList() {
	list := LinkedList{}

	// Create a list of 7 random integers
	for i := 0; i < 7; i++ {
		list.Insert(rand.Intn(100))
	}

	// Iterate over the list and print the value
	list.Iterate()
}
