package main

import "fmt"

type Node struct {
	data string
	next *Node
}

func printNodes(nodes *Node) {
	current := nodes
	fmt.Print("Nodes: ")
	for current != nil {
		fmt.Print(current.data, " -> ")
		current = current.next
	}
	fmt.Print(nil)
}

func printLength(nodes *Node) {
	current := nodes
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

	printNodes(myLinkedList)
	printLength(myLinkedList)
}
