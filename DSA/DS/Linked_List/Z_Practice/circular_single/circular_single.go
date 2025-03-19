package main

type CircularLinkedList struct {
	Head *Node
}

type Node struct {
	Data string
	Next *Node
}

func main() {

	cll := &CircularLinkedList{}

}

func (cll *CircularLinkedList) Append(value string) {
	freshNode := &Node{Data: value}
	if cll.Head == nil {
		cll.Head = freshNode
		cll.Head.Next = cll.Head
		return
	}

	current := cll.Head
	for current.Next != cll.Head {
		current = current.Next
	}
	current.Next = freshNode
	freshNode.Next = cll.Head
}

func (cll *CircularLinkedList) Prepend(value string) {
	freshNode := &Node{Data: value}
	if cll.Head == nil {
		cll.Head = freshNode
		cll.Head.Next = cll.Head
		return
	}

	last := cll.Head
	for last.Next != cll.Head {
		last = last.Next
	}

	freshNode.Next = cll.Head
	last.Next = freshNode
	cll.Head = freshNode
}

func (cll *CircularLinkedList) Delete(value string) bool {

	if cll.Head == nil { // Case 1: Empty list
		return false
	}

	current := cll.Head
	var prev *Node

	// Case 2: Deleting the head node
	if current.Data == value {

		//  if only one is there
		if current.Next == cll.Head {
			cll.Head = nil
			return true
		}

		// NOTE : can make mistake here.understand why we create last var here.
		// Find the last node to update its Next pointer
		last := cll.Head
		for last.Next != cll.Head {
			last = last.Next
		}
		cll.Head = cll.Head.Next // Move head
		last.Next = cll.Head     // Maintain circular connection
		return true
	}

	// Case 3: Deleting a non-head node
	for current.Next != cll.Head {
		if current.Data == value {
			prev.Next = current.Next
			return true
		}
		prev = current
		current = current.Next

	}

	// NOTE: we are retriving above. so anyway, at in current the last node will be present.
	// Case 4: If the last node is the target
	if current.Data == value {
		prev.Next = cll.Head
		return true
	}
	return false // Value not found
}
