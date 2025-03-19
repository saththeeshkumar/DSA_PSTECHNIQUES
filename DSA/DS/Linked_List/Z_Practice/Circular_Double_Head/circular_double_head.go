package main

type CircularLinkedList struct {
	Head *Node
}

type Node struct {
	Data string
	Next *Node
	Prev *Node
}

func main() {

	cll := &CircularLinkedList{}

}

func (cll *CircularLinkedList) Append(value string) {
	freshNode := &Node{Data: value}
	if cll.Head == nil {
		cll.Head = freshNode
		cll.Head.Next = cll.Head
		cll.Head.Prev = cll.Head
		return
	}

	last := cll.Head.Prev

	freshNode.Next = cll.Head
	freshNode.Prev = last
	last.Next = freshNode
	cll.Head.Prev = freshNode
}

func (cll *CircularLinkedList) Prepend(value string) {
	freshNode := &Node{Data: value}
	if cll.Head == nil {
		cll.Head = freshNode
		cll.Head.Next = cll.Head
		cll.Head.Prev = cll.Head
		return
	}

	last := cll.Head.Prev

	freshNode.Next = cll.Head
	freshNode.Prev = last
	last.Next = freshNode
	cll.Head.Prev = freshNode
	cll.Head = freshNode
}

func (cll *CircularLinkedList) Delete(value string) bool {
	if cll.Head == nil {
		return false
	}

	current := cll.Head

	// Case 1: Single node in the list
	if current.Next == cll.Head && current.Data == value {
		cll.Head = nil
		current.Next, current.Prev = nil, nil
		return true
	}

	// Case 2: Head node needs to be deleted
	if current.Data == value {
		last := cll.Head.Prev // Get the last node

		cll.Head = current.Next
		cll.Head.Prev = last
		last.Next = cll.Head

		current.Next, current.Prev = nil, nil // Remove references
		return true
	}

	// Case 3: Deleting a middle or last node
	var prev *Node
	for current.Next != cll.Head {
		if current.Data == value {
			prev.Next = current.Next
			current.Next.Prev = prev

			current.Next, current.Prev = nil, nil // Remove references
			return true
		}
		prev = current
		current = current.Next
	}

	// Case 4: If the last node is the target
	if current.Data == value {
		prev.Next = cll.Head
		cll.Head.Prev = prev

		current.Next, current.Prev = nil, nil // Remove references
		return true
	}

	return false
}
