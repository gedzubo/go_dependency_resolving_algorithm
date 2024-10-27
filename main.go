package main

import "fmt"

type Node struct {
	name  string
	edges []*Node
}

func (n *Node) addEdge(e *Node) {
	n.edges = append(n.edges, e)
}

func resolveDependancies(n *Node) {
	fmt.Println(n.name)
	for _, edge := range n.edges {
		resolveDependancies(edge)
	}
}

func main() {
	a := Node{name: "A"}
	b := Node{name: "B"}
	c := Node{name: "C"}
	d := Node{name: "D"}
	e := Node{name: "E"}

	a.addEdge(&b)
	a.addEdge(&d)
	b.addEdge(&c)
	b.addEdge(&e)
	c.addEdge(&d)
	c.addEdge(&e)

	resolveDependancies(&a)
}
