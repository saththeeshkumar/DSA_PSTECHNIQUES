package main

import "fmt"

type Queue struct {
	Front *Node
	Rear  *Node
}

type Node struct {
	Data int
	Next *Node
	Prev *Node
}

func main() {
	q := &Queue{}
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	q.PrintQueue() // 10 20 30

	fmt.Println(q.Dequeue()) // Output: 10, true
	q.PrintQueue()

	fmt.Println(q.Dequeue()) // Output: 20, true
	q.PrintQueue()

	fmt.Println(q.IsEmpty()) // False
}

// Enqueue adds an element to the rear of the queue
func (q *Queue) Enqueue(value int) {
	newNode := &Node{Data: value}
	if q.Front == nil {
		q.Front, q.Rear = newNode, newNode
	}

	q.Rear.Next = newNode
	newNode.Prev = q.Rear
	q.Rear = newNode
}

func (q *Queue) Dequeue() (int, bool) {

	if q.Front == nil {
		return 0, false
	}

	value := q.Front.Data // to return the value. if we change next element to Front. the current value will be removed. so that i take backup to return the value

	/*
		if q.Front.Next != nil {
			q.Front = q.Front.Next
			return value, true
		}
		q = nil    The line q = nil does not correctly update the queue's reference. In Go, q is a pointer receiver, but assigning q = nil only changes the local copy of the pointer inside the function, not the actual queue instance.
	*/

	q.Front = q.Front.Next // move the Front to the Next Node
	if q.Front == nil {
		q.Rear = nil // If queue is empty, set Rear to nil
	} else {
		q.Front.Prev = nil // Ensure the new Front has no previous reference
	}
	return value, true

}

// Peek returns the front element without removing it
func (q *Queue) Peek() (int, bool) {
	if q.Front == nil {
		return 0, false
	}
	return q.Front.Data, true
}

// IsEmpty checks if the queue is empty
func (q *Queue) IsEmpty() bool {
	return q.Front == nil
}

// PrintQueue prints all elements in the queue from front to rear
func (q *Queue) PrintQueue() {
	if q.Front == nil {
		fmt.Println("Queue is empty")
		return
	}
	current := q.Front
	fmt.Print("Queue: ")
	for current != nil {
		fmt.Print(current.Data, " ")
		current = current.Next
	}
	fmt.Println()
}
