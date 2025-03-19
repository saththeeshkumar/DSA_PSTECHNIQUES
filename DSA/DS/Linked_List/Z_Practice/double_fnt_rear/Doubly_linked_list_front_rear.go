package main

import (
	"fmt"
	"log"
)

type doublyLinkedList struct {
	front *node
	rear  *node
}

type node struct {
	data string
	next *node
	prev *node
}

func main() {

	dll := doublyLinkedList{}

	// Insert values
	dll.append("A")
	dll.append("B")
	dll.append("C")
	dll.append("D")

	fmt.Println("Original List:")
	dll.DisplayForward()

	// Delete middle node
	dll.Delete("B")
	fmt.Println("After Deleting 'B':")
	dll.DisplayForward()

	// Delete first node
	dll.Delete("A")
	fmt.Println("After Deleting 'A':")
	dll.DisplayForward()

	// Delete last node
	dll.Delete("D")
	fmt.Println("After Deleting 'D':")
	dll.DisplayForward()

	// Try deleting a non-existent node
	dll.Delete("Z")

}

func (dll *doublyLinkedList) append(value string) {
	newNode := &node{data: value}
	if dll.front == nil {
		dll.front, dll.rear = newNode, newNode
		return
	}

	newNode.prev = dll.rear
	dll.rear.next = newNode
	dll.rear = newNode

}

func (dll *doublyLinkedList) preppend(value string) {
	newNode := &node{data: value}
	if dll.front == nil {
		dll.front, dll.rear = newNode, newNode
		return
	}

	newNode.next = dll.front
	dll.front.prev = newNode
	dll.front = newNode

}

// Note : couldn't construct this func. keep depth understanding here.
func (dll *doublyLinkedList) Delete(value string) (string, bool) {
	if dll.front == nil {
		log.Println("front is nil in DeleteatFront")
		return "0", false
	}

	// this how i tried to solve this problem
	// if dll.front.data == value {
	// 	dll.front = dll.front.next
	// 	dll.front.prev = nil
	// 	return value, true
	// }

	// current := dll.front

	// for current.next != nil {
	// 	if current.next.data == value {
	// 		current.next = current.next.next
	// 	}
	// 	current = current.next
	// }

	// Actual solutions
	current := dll.front

	// Note : i missed to thik '&& current.data != value'
	if current != nil && current.data != value {
		current = current.next
	}

	// this too
	if current == nil {
		log.Println("given value", value, " not found in delete method")
		return "0", false
	}

	// Delete the first node
	if dll.front == current {
		dll.front = current.next
		if dll.front != nil {
			dll.front.prev = nil
		} else {
			dll.rear = nil
		}
		return value, true
	}

	// Delete the last node
	if dll.rear == current {
		dll.rear = current.prev
		if dll.rear != nil {
			dll.rear.next = nil
		} else {
			dll.front = nil
		}
		return value, true
	}

	// delete in the middle
	current.prev.next = current.next
	current.next.prev = current.prev

	return value, true
}

func (dll *doublyLinkedList) Search(value string) (string, bool) {
	if dll.front == nil {
		log.Println("front is nil in Search")
		return "0", false
	}

	current := dll.front

	for current != nil && current.data == value {
		return value, true
	}
	return "0", false
}

func (dll *doublyLinkedList) DisplayForward() {
	if dll.front == nil {
		log.Println("front is nil in DisplayForward")
	}

	current := dll.front

	if current != nil {
		fmt.Printf("%s <-> ", current.data)
		current = current.next
	}
}

func (dll *doublyLinkedList) DisplayBackward() {
	if dll.front == nil {
		log.Println("front is nil in DisplayForward")
	}

	current := dll.rear

	if current != nil {
		fmt.Printf("%s <-> ", current.data)
		current = current.prev
	}
}
