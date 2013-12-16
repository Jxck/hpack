package main

import (
	"fmt"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

type huff struct {
	data string
	bit  byte
}

var Table = []huff{
	{"a", 0},
	{"b", 1},
	{"c", 2},
	{"d", 3},
	{"e", 4},
	{"f", 5},
	{"g", 6},
	{"h", 7},
}

type node struct {
	left, right *node
	data        string
}

func (n *node) String() string {
	return fmt.Sprintf("%p %p %p (%s)\n", n, n.left, n.right, n.data)
}

func main() {
	root := BuildTree()
	Show(root)
	var bit byte = 3
	log.Println(Decode(root, bit))
}

func BuildTree() (root *node) {
	root = &node{}

	for _, e := range Table {
		current := root
		for {
			if e.bit%2 == 0 {
				next := current.left
				if next == nil {
					next = &node{}
					current.left = next
				}
				current = next
			} else {
				next := current.right
				if next == nil {
					next = &node{}
					current.right = next
				}
				current = next
			}
			e.bit = e.bit / 2
			if e.bit == 0 {
				current.data = e.data
				break
			}
		}
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

func Decode(root *node, bit byte) (data string) {
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
