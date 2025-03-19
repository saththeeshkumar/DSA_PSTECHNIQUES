package main

import "log"

type linkedlist struct {
	head *node
}

type node struct {
	data string
	next *node
}

func main() {

	ll := &linkedlist{}
	ll.append("1")
}

func (ll *linkedlist) append(value string) {
	newNode := &node{data: value}
	if ll.head == nil {
		ll.head = newNode
		return
	}

	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode

}

func (ll *linkedlist) preppend(value string) {
	newNode := &node{data: value}
	if ll.head == nil {
		ll.head = newNode
		return
	}

	newNode.next = ll.head
	ll.head = newNode

	// You can simply do this by below two lines.
	// newNode := &Node{Data: value, Next: ll.head} // Create a new node with the given value and Make the new node point to the current head
	// ll.head = newNode                            // Update the head to the new node
}

func (ll *linkedlist) delete(value string) (string, bool) {
	if ll.head == nil {
		log.Println("head is empty in delete method")
		return "0", false
	}

	if ll.head.data == value {
		ll.head = ll.head.next
		return value, true
	}

	current := ll.head

	for current.next != nil {
		if current.next.data == value {
			current.next = current.next.next
			return value, true
		}
		current = current.next
	}
	log.Println("Node with value", value, "not found.")
	return "0", false
}

func (ll *linkedlist) search(value string) bool {
	// The below commented line for nil check of head is not required to do search operation
	/*
		if ll.head == nil {
			log.Println("head is empty in search method")
			return false
		} */
	current := ll.head
	for current != nil {
		if current.data == value {
			log.Println("Search value ", value, " found")
			return true
		}
		current = current.next
	}
	log.Println("Search value ", value, " not found")
	return false
}
