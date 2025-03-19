package main

import (
	"log"
	"os"
	"strconv"
)

type CircularLinkedList struct {
	Head *Node
}

type Node struct {
	Data int
	Next *Node
}

func main() {
	input := os.Args[1:]
	log.Println("User Input ", input)

	cll := &CircularLinkedList{}

	for _, v := range input {
		i, _ := strconv.Atoi(v)
		cll.Append(i)
	}

	// for _, v := range input {
	// 	i, _ := strconv.Atoi(v)
	// 	cll.Prepend(i)
	// }
	cll.Prepend(4)
	cll.Prepend(5)
	cll.delete(2)
	log.Printf("head %p", cll.Head)

	cll.Print()
	log.Println("End")
}

func (cll *CircularLinkedList) Append(value int) {
	newNode := &Node{Data: value}
	// log.Printf("newNode: %+v, Address: %p\n", newNode, newNode)
	if cll.Head == nil {
		cll.Head = newNode
		cll.Head.Next = cll.Head // Point to itself to form a circular list
		// log.Printf("cll.Head Set: %+v, Address: %p\n", cll.Head, cll.Head)
		return
	}

	current := cll.Head
	for current.Next != cll.Head {
		current = current.Next
	}
	newNode.Next = cll.Head
	current.Next = newNode
	// log.Printf("After Insert, cll.Head: %+v, Address: %p\n", cll.Head, cll.Head)

}

func (cll *CircularLinkedList) Print() {
	if cll.Head == nil {
		log.Println("No values to be printed")
		return
	}

	current := cll.Head

	for {
		log.Println("node ", current)
		if current.Next == cll.Head {
			break
		}
		current = current.Next
	}

}

func (cll *CircularLinkedList) Prepend(value int) {
	newNode := &Node{Data: value}
	if cll.Head == nil {
		newNode.Next = newNode
		cll.Head = newNode
		return
	}

	current := cll.Head
	for current.Next != cll.Head {
		current = current.Next
	}

	newNode.Next = cll.Head
	current.Next = newNode
	cll.Head = newNode

}

func (cll *CircularLinkedList) delete(value int) {

	if cll.Head == nil {
		return
	}

	// Case 1: Only one node in the list
	if cll.Head.Next == cll.Head && cll.Head.Data == value {
		cll.Head = nil
		return
	}

	// Case 2: Delete the Head while it's having multiple Nodes
	if cll.Head.Data == value {
		temp := cll.Head
		for temp.Next != cll.Head {
			temp = temp.Next
		}
		cll.Head = cll.Head.Next
		temp.Next = cll.Head
		return
	}

	current := cll.Head

	// Case 3: Delete the middle or last Node
	for current.Next != cll.Head {
		if current.Next.Data == value {
			current.Next = current.Next.Next // Remove the node
			return
		}
		current = current.Next
	}
}
