package main

import "fmt"

type Node struct {
	data string
	next *Node
}

func (n *Node) printNodes() {
	current := n
	fmt.Print("Nodes: ")
	for current != nil {
		fmt.Print(current.data, " -> ")
		current = current.next
	}
	fmt.Print("nil\n")
}

func (n *Node) findLength() int {
	current := n
	length := 0
	for current != nil {
		current = current.next
		length++
	}
	return length
}

func main() {
	myLinkedList := &Node{
		data: "Jean",
		next: &Node{
			data: "Melo",
			next: nil,
		},
	}

	myLinkedList.printNodes()
	fmt.Print("Length: ", myLinkedList.findLength())
}
