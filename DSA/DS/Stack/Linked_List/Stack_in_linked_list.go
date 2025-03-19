package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type Stack struct {
	Top *Node
}

func main() {

	stack := Stack{}
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	fmt.Println(stack.Pop())     // 30
	fmt.Println(stack.Peek())    // 20
	fmt.Println(stack.IsEmpty()) // false

}

// Push adds an item to the stack
func (st *Stack) Push(value int) {
	newNode := &Node{Data: value, Next: st.Top}
	st.Top = newNode

}

// Pop removes and returns the top item of the stack
func (st *Stack) Pop() (int, bool) {

	if st.Top == nil {
		return 0, false
	}

	current := st.Top.Data

	st.Top = st.Top.Next
	return current, true

}

// Peek returns the top item without removing it
func (st *Stack) Peek() (int, bool) {
	if st.Top == nil {
		return 0, false
	}

	return st.Top.Data, true
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return s.Top == nil
}
