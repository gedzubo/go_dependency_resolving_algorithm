package main

import (
	"fmt"
	"slices"
)

type Node struct {
	name  string
	edges []*Node
}

func (n *Node) addEdge(e *Node) {
	n.edges = append(n.edges, e)
}

func removeNodeFromList(list []*Node, element *Node) []*Node {
	var newList []*Node

	for _, e := range list {
		if e.name != element.name {
			newList = append(newList, e)
		}
	}

	return newList
}

func resolveDependancies(node *Node, resolved []*Node, unresolved []*Node) ([]*Node, []*Node, error) {
	unresolved = append(unresolved, node)
	for _, edge := range node.edges {
		if !slices.Contains(resolved, edge) {
			if slices.Contains(unresolved, edge) {
				return resolved, unresolved, fmt.Errorf("circular dependancy detected: %s -> %s ", node.name, edge.name)
			} else {
				resolved, unresolved, _ = resolveDependancies(edge, resolved, unresolved)
			}
		}
	}
	resolved = append(resolved, node)
	unresolved = removeNodeFromList(unresolved, node)

	return resolved, unresolved, nil
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

	// Circular dependancy
	// d.addEdge(&b)

	var resolved, unresolved []*Node
	result, _, err := resolveDependancies(&a, resolved, unresolved)

	if err != nil {
		fmt.Print(err)
	} else {
		for _, node := range result {
			fmt.Println(node.name)
		}
	}
}
