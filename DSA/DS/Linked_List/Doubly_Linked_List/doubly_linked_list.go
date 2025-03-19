package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type DoublyLinkedList struct {
	Head *Node
}

type Node struct {
	Data int
	Prev *Node
	Next *Node
}

func main() {
	args := make([]string, 0, len(os.Args)-1)
	args = os.Args[1:]
	log.Println("given args ", args)

	dll := &DoublyLinkedList{}
	for _, v := range args {
		vint, _ := strconv.Atoi(v)
		dll.Append(vint)
	}

	dll.InsertAtBeginning(10)
	dll.InsertAtBeginning(9)
	dll.InsertAtBeginning(8)
	dll.InsertAtBeginning(7)

	Current := dll.Head

	for Current != nil {
		log.Printf("Nodes %+v \n ", Current)
		Current = Current.Next
	}

	dll.Delete(6)

	Current = dll.Head
	fmt.Println("After deletion")
	for Current != nil {
		log.Printf("Nodes %+v \n ", Current)
		Current = Current.Next
	}
}

func (dll *DoublyLinkedList) Append(value int) {
	newNode := &Node{Data: value}
	if dll.Head == nil {
		dll.Head = newNode
		return
	}
	current := dll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
	newNode.Prev = current
}

func (dll *DoublyLinkedList) InsertAtBeginning(value int) {
	newNode := &Node{Data: value}
	if dll.Head != nil {
		newNode.Next = dll.Head
		dll.Head.Prev = newNode
	}
	dll.Head = newNode //  If list is empty, make newNode the head
}

func (dll *DoublyLinkedList) Delete(value int) {
	log.Println("TO be deleted ", value)
	if dll.Head == nil {
		log.Println("Head nil")
		return
	}
	current := dll.Head

	for current != nil {
		if current.Data == value {
			if current.Prev != nil {
				current.Prev.Next = current.Next
			} else {
				dll.Head = current.Next
			}

			if current.Next != nil {
				current.Next.Prev = current.Prev
			}
			return
		}
		current = current.Next
	}
}
