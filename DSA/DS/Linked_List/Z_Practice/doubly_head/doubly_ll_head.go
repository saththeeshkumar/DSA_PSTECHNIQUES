package main

import (
	"fmt"
	"log"
	"os"
)

type doublyLinkedList struct {
	head *node
}

type node struct {
	data string
	next *node
	prev *node
}

func main() {

	args := make([]string, 0, len(os.Args)-1)
	args = os.Args[1:]
	log.Println("given args ", args)

	dll := &doublyLinkedList{}
	for _, v := range args {
		dll.append(v)
	}

	dll.insertAtBeginning("10")
	dll.insertAtBeginning("9")
	dll.insertAtBeginning("8")
	dll.insertAtBeginning("7")

	Current := dll.head

	for Current != nil {
		log.Printf("Nodes %+v \n ", Current)
		Current = Current.next
	}

	dll.delete("6")

	Current = dll.head
	fmt.Println("After deletion")
	for Current != nil {
		log.Printf("Nodes %+v \n ", Current)
		Current = Current.next
	}

}

func (dll *doublyLinkedList) append(value string) {
	newnode := &node{data: value}
	if dll.head == nil {
		dll.head = newnode
		return
	}

	current := dll.head

	// put if instead of for
	for current.next != nil {
		current = current.next
	}

	current.next = newnode
	newnode.prev = current

}

func (dll *doublyLinkedList) insertAtBeginning(value string) {
	newNode := &node{data: value}

	// this is what i wrote. we can concise the code.
	// if dll.head == nil {
	// 	dll.head = newNode
	// 	return
	// }

	// newNode.next = dll.head
	// dll.head.next = newNode
	// dll.head = newNode

	if dll.head != nil {
		newNode.next = dll.head
		dll.head.prev = newNode
	}
	dll.head = newNode //  If list is empty, make newNode the head

}

func (dll *doublyLinkedList) delete(value string) (string, bool) {
	if dll.head == nil {
		return "0", false
	}

	current := dll.head

	// Note : need deep understanding. forgot about this implementation
	for current != nil {
		if current.data == value {
			if current.prev != nil {
				current.prev.next = current.next
			} else {
				dll.head = current.next // forgot about this case
				if dll.head != nil {    // Ensure the new head has no previous node and forgot about this case
					dll.head.prev = nil
				}
			}
			if current.next != nil {
				current.next.prev = current.prev
			}
			return value, true
		}
		current = current.next
	}
	return "0", false

}
