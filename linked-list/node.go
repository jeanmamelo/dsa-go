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
	fmt.Print(nil)
}

func (n *Node) printLength() {
	current := n
	length := 0
	for current != nil {
		current = current.next
		length++
	}
	fmt.Print("\nLength: ", length)
}

func main() {
	myLinkedList := &Node{
		data: "Jean",
		next: &Node{
			data: "Marcos",
			next: &Node{
				data: "Albino",
				next: &Node{
					data: "de",
					next: &Node{
						data: "Melo",
						next: nil,
					},
				},
			},
		},
	}

	myLinkedList.printNodes()
	myLinkedList.printLength()
}
