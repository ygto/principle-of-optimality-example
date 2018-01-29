package main

import (
	"math"
	"fmt"
)

type Node struct {
	neighbours []*Node
	value      float64
	name       string
}

func NewNode(name string) *Node {
	n := Node{}
	n.name = name
	n.neighbours = make([]*Node, 0, 0)
	return &n
}
func (node *Node) addNeighbours(neighbours ... *Node) *Node {

	//check old neighbours to protect multiplexing
	for _, neighbour := range neighbours {
		for _, n := range node.neighbours {
			if n.name == neighbour.name {
				return node
			}
		}
		node.neighbours = append(node.neighbours, neighbour)

		// two way neighbourhood
		neighbour.addNeighbours(node)
	}

	return node
}
func (node *Node) Calculate() float64 {
	if node.Calculated() {
		return node.value
	}
	var max float64
	for _, n := range node.neighbours {
		max = math.Max(max, node.value+(n.value*lambda))
	}
	node.value = max

	return node.value
}
func (node *Node) Calculated() bool {
	return node.value != 0
}

func (node *Node) SetAsFinish() {
	node.value = 1
}
func (node *Node) Travel() {

	if node.value == 1 {
		fmt.Printf("%s\n\r", node.name)
		return
	} else {
		fmt.Printf("%s->", node.name)
	}
	var max float64
	var bestNeighbour *Node

	//choose best neighbour
	for _, n := range node.neighbours {
		if n.value > max {
			max = n.value
			bestNeighbour = n
		}
	}
	//continue to travel with best neighbour
	bestNeighbour.Travel()

}

var lambda = 0.9

func main() {

	a := NewNode("a")
	a1 := NewNode("a1")
	b := NewNode("b")
	c := NewNode("c")
	d := NewNode("d")
	e := NewNode("e")
	f := NewNode("f")

	a.addNeighbours(b,a1)
	a1.addNeighbours(f)
	b.addNeighbours(c, d)
	c.addNeighbours(d)
	d.addNeighbours(e)
	e.addNeighbours(f)
	e.addNeighbours(f)
	e.addNeighbours(f)
	e.addNeighbours(f)

	f.SetAsFinish()
	queue := make([]*Node, 0, 0)

	queue = append(queue, f)
	for {
		//if queue is empty calculation completed
		if len(queue) == 0 {
			break
		}
		
		var tempNode *Node
		tempNode, queue = queue[0], queue[1:]
		tempNode.Calculate()
		for _, n := range tempNode.neighbours {
			if !n.Calculated() {
				queue = append(queue, n)
			}
		}
	}

	fmt.Println("a", a.value)
	fmt.Println("a1", a1.value)
	fmt.Println("b", b.value)
	fmt.Println("c", c.value)
	fmt.Println("d", d.value)
	fmt.Println("e", e.value)
	fmt.Println("f", f.value)

	a.Travel()
	c.Travel()
}
