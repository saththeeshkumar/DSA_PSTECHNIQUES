package main

import "fmt"

type priorityqueue struct {
	front *node
}

type node struct {
	data     int
	priority int
	next     *node
	prev     *node
}

func main() {
	pq := &priorityqueue{}

	pq.Enqueue(10, 2)
	pq.Enqueue(20, 1) // Higher priority
	pq.Enqueue(30, 3)
	pq.Enqueue(40, 2)

	pq.printqueue() // Output: (20, P1) (10, P2) (40, P2) (30, P3)

	fmt.Println(pq.Dequeue())
	pq.printqueue() // Output: (10, P2) (40, P2) (30, P3)
}

func (pq *priorityqueue) Enqueue(value, priority int) {
	newNode := &node{data: value, priority: priority}

	// Case 1: Queue is empty
	if pq.front == nil {
		pq.front = newNode
		return
	}

	current := pq.front

	// Case 2: Insert at the front if priority is higher
	if priority < pq.front.priority {
		newNode.next = pq.front
		pq.front.prev = newNode
		pq.front = newNode
		return
	}

	// Case 3: Traverse and insert at the correct position
	for current.next != nil && current.next.priority <= priority {
		current = current.next
	}

	// Insert the new node after `current`
	newNode.prev = current
	newNode.next = current.next
	if current.next != nil {
		current.next.prev = newNode // Need to know why?
	}
	current.next = newNode

}

func (pq *priorityqueue) Dequeue() (int, bool) {
	if pq.front == nil {
		return 0, false
	}

	value := pq.front.data
	pq.front = pq.front.next

	if pq.front != nil {
		pq.front.prev = nil
	}
	return value, true

}
func (pq *priorityqueue) printqueue() {
	if pq.front == nil {
		return
	}
	current := pq.front

	for current != nil {
		fmt.Printf("(%d, P%d) ", current.data, current.priority)
		current = current.next
	}
	fmt.Println()

}
