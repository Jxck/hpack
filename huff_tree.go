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

// huffman tree の node
type node struct {
	left, right *node // 左右の子ノードへのポインタ
	data        int   // Table へのインデクスを持つ
}

// 空のノードを生成
// data int を持たないことを -1 で表わす
func NewNode() *node {
	return &node{nil, nil, -1}
}

// デバッグ用
func (n *node) String() string {
	return fmt.Sprintf(
		"%p %p %p (%d)\n",
		n, n.left, n.right, n.data)
}

// デバッグ用
func Show(current *node) {
	if current != nil {
		fmt.Println(current)
		Show(current.left)
		Show(current.right)
	}
}

// ツリーを構築する
func BuildTree() (root *node) {
	root = NewNode()

	for i, e := range RequestHuffmanTable { // 全てのコードについて実施
		current := root // たどる時は根から

		for e.length > 0 { // 1 コードの長さを消化するまで繰り返す
			e.length -= 1
			mask := uint32(1 << e.length) // length=5 なら 10000 でマスクする
			if e.code&mask == mask {      // 11011 & 10000 = 10000 マスク結果とマスクの比較でわかる
				next := current.right // 1 だったら右
				if next == nil {      // 無かったらノードを足す
					next = NewNode()
					current.right = next
				}
				current = next
			} else {
				next := current.left // 0 だったら左
				if next == nil {     // 無かったらノードを足す
					next = NewNode()
					current.left = next
				}
				current = next
			}
		} // 木をコード長まで降りたところ
		// ここにテーブルのインデックスを入れる
		current.data = i
	}
	return root
}

/* テストデータ
[141,       127,       91       ]
[1000 1101, 0111 1111, 0101 1011]
   e   d    c  h   h    f   d  d

0 a 000,   4 e 100
1 b 001,   5 f 101
2 c 010,   6 g 110
3 d 011,   7 h 111
*/

// バイト配列を渡すと木を辿り、その葉にあるテーブルへのインデックスの配列を返す。
func Decode(root *node, codes []byte) []int {
	// 初期化
	var result []int
	var mask byte

	// スタート地点
	current := root
	for _, code := range codes { // バイト配列の頭から順に
		mask = 128 // 常に 1000 0000 をマスクに使う。 1<<7 から定数に変更した
		log.Printf("%b", code)
		for mask > 0 { // 8bit 読み終わるまで
			log.Printf("%b %b %b", code, mask, code&mask)
			if code&mask == mask { // 1100 & 1000 = 1000 マスクすると先頭ビットが分かる
				current = current.right // 1 なら右
			} else {
				current = current.left // 0 なら左
			}
			if current.data != -1 { // 移動先で値を見つけたら
				log.Println("==", current.data)
				result = append(result, current.data) // それを結果に追加
				current = root                        // 残った bit を使ってまた根からたどる
			}
			mask = mask >> 1 // 1bit 処理したら減らす
		}
	}
	return result
}

func main() {
	root := BuildTree()
	Show(root)
	var code = []byte{141, 127, 91}
	log.Println(Decode(root, code))
}
