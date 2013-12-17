package main

import (
	"fmt"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

type huff struct {
	code uint32
	bit  uint8
}

var Table = []huff{
	{0, 3},
	{1, 3},
	{2, 3},
	{3, 3},
	{4, 3},
	//{5, 3},
	//{6, 3},
	//{7, 3},
}

type node struct {
	left, right *node
	data        int
}

func (n *node) String() string {
	return fmt.Sprintf(
		"%p %p %p (%d)\n",
		n, n.left, n.right, n.data)
}

func main() {
	root := BuildTree()
	Show(root)
	var bit byte = 4
	log.Println(Decode(root, bit))
}

func BuildTree() (root *node) {
	root = &node{nil, nil, -1}

	for i, e := range Table {
		current := root
		for e.bit > 0 {
			if e.code%2 == 0 {
				next := current.left
				if next == nil {
					next = &node{nil, nil, -1}
					current.left = next
				}
				current = next
			} else {
				next := current.right
				if next == nil {
					next = &node{nil, nil, -1}
					current.right = next
				}
				current = next
			}
			e.code = e.code / 2
			e.bit -= 1
		}
		current.data = i
	}
	return root
}

func Show(current *node) {
	if current != nil {
		fmt.Println(current)
		Show(current.left)
		Show(current.right)
	}
}

func Decode(root *node, bit byte) (data int) {
	current := root
	for bit > 0 {
		if bit%2 == 0 {
			current = current.left
		} else {
			current = current.right
		}
		bit = bit / 2
	}
	return current.data
}
