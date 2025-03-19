package main

import (
	"log"
	"os"
	"strconv"
)

type SinglyLinkedList struct {
	Head *Node
}

type Node struct {
	Data int
	Next *Node
}

func main() {
	inputs := os.Args[1:]
	ll := &SinglyLinkedList{}
	for _, v := range inputs {
		integer, _ := strconv.Atoi(v)
		ll.createLinkedList(integer)
	}

	temp := ll.Head
	for temp != nil {
		log.Println(" Node ", temp)
		temp = temp.Next
	}

	log.Println("Middle of the Node ", ll.findMiddleOftheNode())
}

func (ll *SinglyLinkedList) createLinkedList(value int) {
	newNode := &Node{Data: value}

	if ll.Head == nil {
		ll.Head = newNode
		return
	}
	temp := ll.Head

	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = newNode
}

func (ll *SinglyLinkedList) findMiddleOftheNode() *Node {
	slow, fast := ll.Head, ll.Head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
