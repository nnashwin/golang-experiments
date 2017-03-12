package main

import (
	"errors"
	"log"
	"reflect"
)

type Node struct {
	Value int
	Data  string
	Left  *Node
	Right *Node
}

func (n *Node) Insert(value int, data string) error {
	if n == nil {
		return errors.New("Cannot insert val into a nil tree")
	}

	switch {
	case value == n.Value:
		return nil

	case value < n.Value:
		if n.Left == nil {
			n.Left = &Node{Value: value, Data: data}
			return nil
		}
		return n.Left.Insert(value, data)

	case value > n.Value:
		if n.Right == nil {
			n.Right = &Node{Value: value, Data: data}
			return nil
		}
		return n.Right.Insert(value, data)
	}
	return nil
}

func main() {
	n := Node{Value: 5, Data: "data"}
	n.Insert(6, "sixes")
	log.Println(n.Right.Value)
	log.Println(reflect.TypeOf(n.Right.Value))

}
