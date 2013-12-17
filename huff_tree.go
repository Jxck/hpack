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
	var code = []byte{141, 127, 91}
	log.Println(Decode(root, code))
}

func BuildTree() (root *node) {
	root = &node{nil, nil, -1}

	for i, e := range Table {
		current := root
		for e.bit > 0 {
			e.bit -= 1
			mask := pow(2, e.bit)
			if e.code&mask == mask {
				next := current.right
				if next == nil {
					next = &node{nil, nil, -1}
					current.right = next
				}
				current = next
			} else {
				next := current.left
				if next == nil {
					next = &node{nil, nil, -1}
					current.left = next
				}
				current = next
			}
		}
		current.data = i
	}
	return root
}

func pow(x, y uint8) (ans uint32) {
	x32 := uint32(x)
	if y == 0 {
		return 1
	}

	ans = x32
	for y > 1 {
		ans *= x32
		y--
	}
	return ans
}

func Show(current *node) {
	if current != nil {
		fmt.Println(current)
		Show(current.left)
		Show(current.right)
	}
}

func Decode(root *node, codes []byte) (data int) {
	current := root
	var mask byte = 1 << 7
	for _, code := range codes {
		log.Printf("%b", code)
		for mask > 0 {
			log.Printf("%b", mask)
			if current.data != -1 {
				break
			}
			if code&mask == mask {
				current = current.right
			} else {
				current = current.left
			}

			mask = mask >> 1
		}
		log.Println(current.data)
	}
	return current.data
}
