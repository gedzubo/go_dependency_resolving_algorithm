package main

import (
	"reflect"
	"testing"
)

func TestResolveDependancies(t *testing.T) {
	t.Run("should return a list of resolved dependacies", func(t *testing.T) {
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

		var resolved, unresolved []*Node
		expected := []*Node{&d, &e, &c, &b, &a}
		got, _, _ := ResolveDependancies(&a, resolved, unresolved)

		if !reflect.DeepEqual(expected, got) {
			t.Errorf("got: %v ,expected: %v", got, expected)
		}
	})

	t.Run("should return an error if circular dependancy detected", func(t *testing.T) {
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
		d.addEdge(&b)

		var resolved, unresolved []*Node
		expected := "circular dependancy detected: A -> D"
		_, _, got := ResolveDependancies(&a, resolved, unresolved)

		if expected != got.Error() {
			t.Errorf("got: %v ,expected: %v", got, expected)
		}
	})
}
