package main

import (
	"testing"
)

func TestBinaryTreeInsert(t *testing.T) {
	n := Node{Value: 5, Data: "fives"}
	n.Insert(6, "sixes")
	if n.Right.Value != 6 {
		t.Error("binary tree doesn't insert higher value into Right Pointer")
	}

	n.Insert(4, "fours")
	if n.Left.Value != 4 {
		t.Error("binary tree doesn't insert lower value into Left Pointer")
	}
}
