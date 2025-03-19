package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type SinglyLinkedList struct {
	Head *Node
}

type Node struct {
	Data int
	Next *Node
}

func main() {
	inputs := os.Args[1:]
	ll := &SinglyLinkedList{}
	for _, v := range inputs {
		integer, _ := strconv.Atoi(v)
		ll.createCyclicLinkedList(integer)
	}

	temp := ll.Head
	for temp.Next != ll.Head {
		log.Println(" Node ", temp)
		temp = temp.Next
	}
	isCyclic, node := ll.findCyclicNode()
	log.Println("cyclic of the Node ", isCyclic, node)
}

func (ll *SinglyLinkedList) createCyclicLinkedList(value int) {
	newNode := &Node{Data: value}

	if ll.Head == nil {
		ll.Head = newNode
		ll.Head.Next = ll.Head
		return
	}
	temp := ll.Head

	for temp.Next != ll.Head {
		temp = temp.Next
	}
	newNode.Next = ll.Head
	temp.Next = newNode
}

func (ll *SinglyLinkedList) findCyclicNode() (bool, *Node) {
	slow, fast := ll.Head, ll.Head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			fmt.Printf("slow %+v and %p\n", slow, slow)
			fmt.Printf("fast %+v and %p\n", fast, fast)
			return true, slow
		}
	}
	return false, nil
}
