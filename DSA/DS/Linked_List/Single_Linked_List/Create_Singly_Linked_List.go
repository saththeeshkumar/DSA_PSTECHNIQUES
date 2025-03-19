package main

import (
	"log"
	"os"
	"strconv"
)

type LinkedList struct {
	Head *Node
}

type Node struct {
	Data int
	Next *Node
}

func main() {

	args := make([]string, 0, len(os.Args)-1)
	args = os.Args[1:]
	log.Println("given args ", args)

	ll := &LinkedList{}

	for _, v := range args {
		log.Println("value ", v)
		vint, _ := strconv.Atoi(v)
		ll.AppendLinkedList(vint)
		// ll.InsertAtBeginning(vint)
	}

	temp := ll.Head
	for temp != nil {
		log.Printf("Nodes %+v \n ", temp)
		temp = temp.Next
	}

	ll.DeletetheNode(len(args)) // Delete the last input
	log.Println("After deleted")

	temp = ll.Head
	for temp != nil {
		log.Printf("Nodes %+v \n ", temp)
		temp = temp.Next
	}

	found := ll.SearchtheNode(len(args) - 1)
	log.Printf("the value %v  %v", len(args)-1, found)

}

func (ll *LinkedList) AppendLinkedList(value int) {
	log.Println("ll.Head == nil ", ll.Head == nil, value)
	newNode := &Node{Data: value, Next: nil} // Create a new node

	if ll.Head == nil { // If the list is empty
		ll.Head = newNode // Set the new node as the head
		return
	}

	temp := ll.Head // Start from the head

	for temp.Next != nil { // Traverse until the last node
		temp = temp.Next // Move to the next node
	}
	temp.Next = newNode // Link the last node to the new node
}

func (ll *LinkedList) InsertAtBeginning(value int) {
	newNode := &Node{Data: value, Next: ll.Head} // Create a new node with the given value and Make the new node point to the current head
	ll.Head = newNode                            // Update the head to the new node
}

func (ll *LinkedList) DeletetheNode(value int) {
	log.Println("Delete the node which containd value of ", value)
	if ll.Head == nil { // If the list is empty, there's nothing to delete
		return
	}

	// Case 1: If the node to be deleted is the head node
	if ll.Head.Data == value {
		ll.Head = ll.Head.Next // Update the head to point to the next node
		return
	}

	// Case 2: Traverse the list to find the node and delete it
	temp := ll.Head
	for temp.Next != nil {

		if temp.Next.Data == value { // If the next node contains the value
			temp.Next = temp.Next.Next // Bypass the node to delete it
			return
		}

		temp = temp.Next // Move to the next node
	}

	// If we reached here, it means the value was not found
	log.Println("Node with value", value, "not found.")
}

func (ll *LinkedList) SearchtheNode(value int) bool {
	log.Println("TO be found ", value)
	temp := ll.Head

	for temp.Next != nil {
		if temp.Next.Data == value {
			return true
		}
		temp = temp.Next
	}
	return false
}
