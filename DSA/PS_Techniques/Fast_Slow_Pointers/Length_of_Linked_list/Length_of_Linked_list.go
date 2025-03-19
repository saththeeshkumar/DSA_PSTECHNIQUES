package main

import (
	"fmt"
	"log"
)

// Node represents a node in the linked list
type Node struct {
	data int
	next *Node
}

func main() {
	// Creating nodes
	node1 := &Node{data: 1}
	node2 := &Node{data: 2}
	node3 := &Node{data: 3}
	node4 := &Node{data: 4}
	node5 := &Node{data: 5}

	// Linking nodes to form a linked list with a cycle
	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5
	node5.next = node2 // Creates a cycle starting at node2
	log.Printf("node2 %p \n", node2)
	log.Printf("node5 %p \n", node5)
	log.Println("node5 ", node5)
	log.Printf("node5.next %p \n", node5.next)

	// Finding the total length of the linked list
	totalLength := findTotalLength(node1)
	fmt.Printf("The total length of the linked list is: %d\n", totalLength)
}

func findTotalLength(head *Node) int {
	if head == nil {
		return 0
	}
	// Find the given list cyclic or not.
	meetingPoint := hasCyclic(head)
	log.Println("meetingPoint ", meetingPoint)
	log.Printf("meetingPoint %p \n", meetingPoint)
	// In this case, the given linkedlist is not cyclic.
	if meetingPoint == nil {
		length := 0
		for current := head; current != nil; current = current.next {
			length++
		}
		return length
	}

	// find the cycleLength
	cyclicStart := findCycleStart(head, meetingPoint)
	log.Println("cyclicStart ", cyclicStart)
	log.Printf("cyclicStart %p \n", cyclicStart)
	cycleLength := findCycleLength(cyclicStart)
	nonCycleLength := findNonCycleLength(head, cyclicStart)
	fmt.Printf("cycleLength %d \n nonCycleLength %d \n totalLength %d ", cycleLength, nonCycleLength, cycleLength+nonCycleLength)
	return cycleLength + nonCycleLength
}

func hasCyclic(head *Node) *Node {
	if head == nil {
		return nil
	}

	slow, fast := head, head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			return slow
		}
	}
	return nil
}

func findCycleStart(head, meetingPoint *Node) *Node {
	ptr1 := head
	ptr2 := meetingPoint
	for ptr1 != ptr2 {
		ptr1 = ptr1.next
		ptr2 = ptr2.next
	}
	return ptr1
}

func findCycleLength(meetingPoint *Node) int {
	log.Println("meetingPoint in findCycleLength ", meetingPoint)
	log.Printf("meetingPoint in findCycleLength %p \n", meetingPoint)
	current := meetingPoint
	log.Println("current in findCycleLength ", current)
	log.Printf("current in findCycleLength %p \n", current)
	length := 0

	for {
		current = current.next
		if current.data == 2 {
			log.Printf("\n\n\n")
			log.Println("current in findCycleLength inside 2", current)
			log.Printf("current in findCycleLength inside 2 %p \n", current)
		}
		length++
		if current == meetingPoint {
			log.Printf("\n\n\n")
			log.Println("current in findCycleLength inside break", current)
			log.Printf("current in findCycleLength inside break %p \n", current)
			log.Println("meetingPoint in findCycleLength inside break ", meetingPoint)
			log.Printf("meetingPoint in findCycleLength inside break %p \n", meetingPoint)
			break
		}
	}
	return length

}

func findNonCycleLength(head, cycleStart *Node) int {

	nonCycleLength := 0

	for current := head; current != cycleStart; current = current.next {
		nonCycleLength++
	}

	return nonCycleLength
}
