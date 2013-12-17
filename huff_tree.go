package main

import (
	"fmt"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

type HuffmanCode struct {
	code   uint32
	length uint8
}

var RequestHuffmanTable = [257]HuffmanCode{
	{0, 3},
	{1, 3},
	{2, 3},
	{3, 3},
	{4, 3},
	{5, 3},
	{6, 3},
	{7, 3},
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

	for i, e := range RequestHuffmanTable {
		current := root
		for e.length > 0 {
			e.length -= 1
			mask := uint32(1 << e.length)
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

func Show(current *node) {
	if current != nil {
		fmt.Println(current)
		Show(current.left)
		Show(current.right)
	}
}

// byte の配列をデコードする際に、
// 値をうまく引き継ぐようにする。

/* テストデータ
[141,       127,       91       ]
[1000 1101, 0111 1111, 0101 1011]
   e   d    c  h   h    f   d  d

0 a 000,   4 e 100
1 b 001,   5 f 101
2 c 010,   6 g 110
3 d 011,   7 h 111
*/
func Decode(root *node, codes []byte) []int {
	var result []int
	current := root
	var mask byte
	for _, code := range codes {
		mask = 1 << 7
		log.Printf("%b", code)
		for mask > 0 {
			log.Printf("%b %b %b", code, mask, code&mask)
			if code&mask == mask {
				current = current.right
			} else {
				current = current.left
			}
			if current.data != -1 {
				log.Println("==", current.data)
				result = append(result, current.data)
				current = root
			}
			mask = mask >> 1
		}
	}
	return result
}
