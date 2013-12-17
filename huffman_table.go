package hpack

import (
	"fmt"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

// ===== Encode =====

// 結果 1 Byte を生成するための struct
type Byte struct {
	value  uint32
	remain uint8
}

func NewByte() Byte {
	return Byte{0, 8}
}

type HuffmanCode struct {
	code   uint32
	length uint8
}

func HuffmanEncodeRequest(raw []byte) (encoded []byte) {
	return HuffmanEncode(&RequestHuffmanTable, raw)
}
func HuffmanEncodeResponse(raw []byte) (encoded []byte) {
	return HuffmanEncode(&ResponseHuffmanTable, raw)
}

func HuffmanEncode(table *[257]HuffmanCode, raw []byte) (encoded []byte) {
	// 1 byte の入れ物
	byt := NewByte()

	for _, v := range raw {

		// huffman table で変換
		huff := table[v]

		for huff.length > 0 { // huff.code を使いきるまで

			if byt.remain > huff.length {
				// huff.code の全てを入れる

				// 左シフトして桁を合わせる
				shift := byt.remain - huff.length
				tmp := huff.code << shift

				// byt に追加、入れ多分だけ長さを減らす
				byt.value += tmp
				byt.remain = shift

				// huff は空に
				huff.length = 0

			} else {
				// huff.code の一部を入れる

				// 右シフトして入れる分だけ切り出す
				shift := huff.length - byt.remain
				tmp := huff.code >> shift

				// byt に追加、もう入らない
				byt.value += tmp
				byt.remain = 0

				// huff から使った分だけ減らす
				huff.code -= (tmp << shift)
				huff.length = shift
			}

			if byt.remain == 0 {
				// byt が埋まったら配列に移して初期化
				encoded = append(encoded, byte(byt.value))
				byt = NewByte()
			}
		}
	}

	if byt.remain > 0 { // EOS でパディング
		// パディング分切り出す
		eos := RequestHuffmanTable[256]
		shift := eos.length - byt.remain
		padding := eos.code >> shift

		// byt に加える
		byt.value += padding
		byt.remain = 0

		// 配列に移す。最後なので初期化はしない。
		encoded = append(encoded, byte(byt.value))
	}

	return encoded
}

// ===== Decode =====

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

func BuildRequestTree() (root *node) {
	return BuildTree(&RequestHuffmanTable)
}

func BuildResponseTree() (root *node) {
	return BuildTree(&ResponseHuffmanTable)
}

// ツリーを構築する
func BuildTree(table *[257]HuffmanCode) (root *node) {
	root = NewNode()

	for i, e := range table { // 全てのコードについて実施
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
func HuffmanDecode(root *node, codes []byte) []int {
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

/*
  RequestHuffmanTable  [257]HuffmanCode
  ResponseHuffmanTable [257]HuffmanCode
*/
var RequestHuffmanTable = [257]HuffmanCode{
	{0x7ffffba, 27},
	{0x7ffffbb, 27},
	{0x7ffffbc, 27},
	{0x7ffffbd, 27},
	{0x7ffffbe, 27},
	{0x7ffffbf, 27},
	{0x7ffffc0, 27},
	{0x7ffffc1, 27},
	{0x7ffffc2, 27},
	{0x7ffffc3, 27},
	{0x7ffffc4, 27},
	{0x7ffffc5, 27},
	{0x7ffffc6, 27},
	{0x7ffffc7, 27},
	{0x7ffffc8, 27},
	{0x7ffffc9, 27},
	{0x7ffffca, 27},
	{0x7ffffcb, 27},
	{0x7ffffcc, 27},
	{0x7ffffcd, 27},
	{0x7ffffce, 27},
	{0x7ffffcf, 27},
	{0x7ffffd0, 27},
	{0x7ffffd1, 27},
	{0x7ffffd2, 27},
	{0x7ffffd3, 27},
	{0x7ffffd4, 27},
	{0x7ffffd5, 27},
	{0x7ffffd6, 27},
	{0x7ffffd7, 27},
	{0x7ffffd8, 27},
	{0x7ffffd9, 27},
	{0xe8, 8},
	{0xffc, 12},
	{0x3ffa, 14},
	{0x7ffc, 15},
	{0x7ffd, 15},
	{0x24, 6},
	{0x6e, 7},
	{0x7ffe, 15},
	{0x7fa, 11},
	{0x7fb, 11},
	{0x3fa, 10},
	{0x7fc, 11},
	{0xe9, 8},
	{0x25, 6},
	{0x4, 5},
	{0x0, 4},
	{0x5, 5},
	{0x6, 5},
	{0x7, 5},
	{0x26, 6},
	{0x27, 6},
	{0x28, 6},
	{0x29, 6},
	{0x2a, 6},
	{0x2b, 6},
	{0x2c, 6},
	{0x1ec, 9},
	{0xea, 8},
	{0x3fffe, 18},
	{0x2d, 6},
	{0x1fffc, 17},
	{0x1ed, 9},
	{0x3ffb, 14},
	{0x6f, 7},
	{0xeb, 8},
	{0xec, 8},
	{0xed, 8},
	{0xee, 8},
	{0x70, 7},
	{0x1ee, 9},
	{0x1ef, 9},
	{0x1f0, 9},
	{0x1f1, 9},
	{0x3fb, 10},
	{0x1f2, 9},
	{0xef, 8},
	{0x1f3, 9},
	{0x1f4, 9},
	{0x1f5, 9},
	{0x1f6, 9},
	{0x1f7, 9},
	{0xf0, 8},
	{0xf1, 8},
	{0x1f8, 9},
	{0x1f9, 9},
	{0x1fa, 9},
	{0x1fb, 9},
	{0x1fc, 9},
	{0x3fc, 10},
	{0x3ffc, 14},
	{0x7ffffda, 27},
	{0x1ffc, 13},
	{0x3ffd, 14},
	{0x2e, 6},
	{0x7fffe, 19},
	{0x8, 5},
	{0x2f, 6},
	{0x9, 5},
	{0x30, 6},
	{0x1, 4},
	{0x31, 6},
	{0x32, 6},
	{0x33, 6},
	{0xa, 5},
	{0x71, 7},
	{0x72, 7},
	{0xb, 5},
	{0x34, 6},
	{0xc, 5},
	{0xd, 5},
	{0xe, 5},
	{0xf2, 8},
	{0xf, 5},
	{0x10, 5},
	{0x11, 5},
	{0x35, 6},
	{0x73, 7},
	{0x36, 6},
	{0xf3, 8},
	{0xf4, 8},
	{0xf5, 8},
	{0x1fffd, 17},
	{0x7fd, 11},
	{0x1fffe, 17},
	{0xffd, 12},
	{0x7ffffdb, 27},
	{0x7ffffdc, 27},
	{0x7ffffdd, 27},
	{0x7ffffde, 27},
	{0x7ffffdf, 27},
	{0x7ffffe0, 27},
	{0x7ffffe1, 27},
	{0x7ffffe2, 27},
	{0x7ffffe3, 27},
	{0x7ffffe4, 27},
	{0x7ffffe5, 27},
	{0x7ffffe6, 27},
	{0x7ffffe7, 27},
	{0x7ffffe8, 27},
	{0x7ffffe9, 27},
	{0x7ffffea, 27},
	{0x7ffffeb, 27},
	{0x7ffffec, 27},
	{0x7ffffed, 27},
	{0x7ffffee, 27},
	{0x7ffffef, 27},
	{0x7fffff0, 27},
	{0x7fffff1, 27},
	{0x7fffff2, 27},
	{0x7fffff3, 27},
	{0x7fffff4, 27},
	{0x7fffff5, 27},
	{0x7fffff6, 27},
	{0x7fffff7, 27},
	{0x7fffff8, 27},
	{0x7fffff9, 27},
	{0x7fffffa, 27},
	{0x7fffffb, 27},
	{0x7fffffc, 27},
	{0x7fffffd, 27},
	{0x7fffffe, 27},
	{0x7ffffff, 27},
	{0x3ffff80, 26},
	{0x3ffff81, 26},
	{0x3ffff82, 26},
	{0x3ffff83, 26},
	{0x3ffff84, 26},
	{0x3ffff85, 26},
	{0x3ffff86, 26},
	{0x3ffff87, 26},
	{0x3ffff88, 26},
	{0x3ffff89, 26},
	{0x3ffff8a, 26},
	{0x3ffff8b, 26},
	{0x3ffff8c, 26},
	{0x3ffff8d, 26},
	{0x3ffff8e, 26},
	{0x3ffff8f, 26},
	{0x3ffff90, 26},
	{0x3ffff91, 26},
	{0x3ffff92, 26},
	{0x3ffff93, 26},
	{0x3ffff94, 26},
	{0x3ffff95, 26},
	{0x3ffff96, 26},
	{0x3ffff97, 26},
	{0x3ffff98, 26},
	{0x3ffff99, 26},
	{0x3ffff9a, 26},
	{0x3ffff9b, 26},
	{0x3ffff9c, 26},
	{0x3ffff9d, 26},
	{0x3ffff9e, 26},
	{0x3ffff9f, 26},
	{0x3ffffa0, 26},
	{0x3ffffa1, 26},
	{0x3ffffa2, 26},
	{0x3ffffa3, 26},
	{0x3ffffa4, 26},
	{0x3ffffa5, 26},
	{0x3ffffa6, 26},
	{0x3ffffa7, 26},
	{0x3ffffa8, 26},
	{0x3ffffa9, 26},
	{0x3ffffaa, 26},
	{0x3ffffab, 26},
	{0x3ffffac, 26},
	{0x3ffffad, 26},
	{0x3ffffae, 26},
	{0x3ffffaf, 26},
	{0x3ffffb0, 26},
	{0x3ffffb1, 26},
	{0x3ffffb2, 26},
	{0x3ffffb3, 26},
	{0x3ffffb4, 26},
	{0x3ffffb5, 26},
	{0x3ffffb6, 26},
	{0x3ffffb7, 26},
	{0x3ffffb8, 26},
	{0x3ffffb9, 26},
	{0x3ffffba, 26},
	{0x3ffffbb, 26},
	{0x3ffffbc, 26},
	{0x3ffffbd, 26},
	{0x3ffffbe, 26},
	{0x3ffffbf, 26},
	{0x3ffffc0, 26},
	{0x3ffffc1, 26},
	{0x3ffffc2, 26},
	{0x3ffffc3, 26},
	{0x3ffffc4, 26},
	{0x3ffffc5, 26},
	{0x3ffffc6, 26},
	{0x3ffffc7, 26},
	{0x3ffffc8, 26},
	{0x3ffffc9, 26},
	{0x3ffffca, 26},
	{0x3ffffcb, 26},
	{0x3ffffcc, 26},
	{0x3ffffcd, 26},
	{0x3ffffce, 26},
	{0x3ffffcf, 26},
	{0x3ffffd0, 26},
	{0x3ffffd1, 26},
	{0x3ffffd2, 26},
	{0x3ffffd3, 26},
	{0x3ffffd4, 26},
	{0x3ffffd5, 26},
	{0x3ffffd6, 26},
	{0x3ffffd7, 26},
	{0x3ffffd8, 26},
	{0x3ffffd9, 26},
	{0x3ffffda, 26},
	{0x3ffffdb, 26},
	{0x3ffffdc, 26},
}

var ResponseHuffmanTable = [257]HuffmanCode{
	{0x1ffffbc, 25},
	{0x1ffffbd, 25},
	{0x1ffffbe, 25},
	{0x1ffffbf, 25},
	{0x1ffffc0, 25},
	{0x1ffffc1, 25},
	{0x1ffffc2, 25},
	{0x1ffffc3, 25},
	{0x1ffffc4, 25},
	{0x1ffffc5, 25},
	{0x1ffffc6, 25},
	{0x1ffffc7, 25},
	{0x1ffffc8, 25},
	{0x1ffffc9, 25},
	{0x1ffffca, 25},
	{0x1ffffcb, 25},
	{0x1ffffcc, 25},
	{0x1ffffcd, 25},
	{0x1ffffce, 25},
	{0x1ffffcf, 25},
	{0x1ffffd0, 25},
	{0x1ffffd1, 25},
	{0x1ffffd2, 25},
	{0x1ffffd3, 25},
	{0x1ffffd4, 25},
	{0x1ffffd5, 25},
	{0x1ffffd6, 25},
	{0x1ffffd7, 25},
	{0x1ffffd8, 25},
	{0x1ffffd9, 25},
	{0x1ffffda, 25},
	{0x1ffffdb, 25},
	{0x0, 4},
	{0xffa, 12},
	{0x6a, 7},
	{0x1ffa, 13},
	{0x3ffc, 14},
	{0x1ec, 9},
	{0x3f8, 10},
	{0x1ffb, 13},
	{0x1ed, 9},
	{0x1ee, 9},
	{0xffb, 12},
	{0x7fa, 11},
	{0x22, 6},
	{0x23, 6},
	{0x24, 6},
	{0x6b, 7},
	{0x1, 4},
	{0x2, 4},
	{0x3, 4},
	{0x8, 5},
	{0x9, 5},
	{0xa, 5},
	{0x25, 6},
	{0x26, 6},
	{0xb, 5},
	{0xc, 5},
	{0xd, 5},
	{0x1ef, 9},
	{0xfffa, 16},
	{0x6c, 7},
	{0x1ffc, 13},
	{0xffc, 12},
	{0xfffb, 16},
	{0x6d, 7},
	{0xea, 8},
	{0xeb, 8},
	{0xec, 8},
	{0xed, 8},
	{0xee, 8},
	{0x27, 6},
	{0x1f0, 9},
	{0xef, 8},
	{0xf0, 8},
	{0x3f9, 10},
	{0x1f1, 9},
	{0x28, 6},
	{0xf1, 8},
	{0xf2, 8},
	{0x1f2, 9},
	{0x3fa, 10},
	{0x1f3, 9},
	{0x29, 6},
	{0xe, 5},
	{0x1f4, 9},
	{0x1f5, 9},
	{0xf3, 8},
	{0x3fb, 10},
	{0x1f6, 9},
	{0x3fc, 10},
	{0x7fb, 11},
	{0x1ffd, 13},
	{0x7fc, 11},
	{0x7ffc, 15},
	{0x1f7, 9},
	{0x1fffe, 17},
	{0xf, 5},
	{0x6e, 7},
	{0x2a, 6},
	{0x2b, 6},
	{0x10, 5},
	{0x6f, 7},
	{0x70, 7},
	{0x71, 7},
	{0x2c, 6},
	{0x1f8, 9},
	{0x1f9, 9},
	{0x72, 7},
	{0x2d, 6},
	{0x2e, 6},
	{0x2f, 6},
	{0x30, 6},
	{0x1fa, 9},
	{0x31, 6},
	{0x32, 6},
	{0x33, 6},
	{0x34, 6},
	{0x73, 7},
	{0xf4, 8},
	{0x74, 7},
	{0xf5, 8},
	{0x1fb, 9},
	{0xfffc, 16},
	{0x3ffd, 14},
	{0xfffd, 16},
	{0xfffe, 16},
	{0x1ffffdc, 25},
	{0x1ffffdd, 25},
	{0x1ffffde, 25},
	{0x1ffffdf, 25},
	{0x1ffffe0, 25},
	{0x1ffffe1, 25},
	{0x1ffffe2, 25},
	{0x1ffffe3, 25},
	{0x1ffffe4, 25},
	{0x1ffffe5, 25},
	{0x1ffffe6, 25},
	{0x1ffffe7, 25},
	{0x1ffffe8, 25},
	{0x1ffffe9, 25},
	{0x1ffffea, 25},
	{0x1ffffeb, 25},
	{0x1ffffec, 25},
	{0x1ffffed, 25},
	{0x1ffffee, 25},
	{0x1ffffef, 25},
	{0x1fffff0, 25},
	{0x1fffff1, 25},
	{0x1fffff2, 25},
	{0x1fffff3, 25},
	{0x1fffff4, 25},
	{0x1fffff5, 25},
	{0x1fffff6, 25},
	{0x1fffff7, 25},
	{0x1fffff8, 25},
	{0x1fffff9, 25},
	{0x1fffffa, 25},
	{0x1fffffb, 25},
	{0x1fffffc, 25},
	{0x1fffffd, 25},
	{0x1fffffe, 25},
	{0x1ffffff, 25},
	{0xffff80, 24},
	{0xffff81, 24},
	{0xffff82, 24},
	{0xffff83, 24},
	{0xffff84, 24},
	{0xffff85, 24},
	{0xffff86, 24},
	{0xffff87, 24},
	{0xffff88, 24},
	{0xffff89, 24},
	{0xffff8a, 24},
	{0xffff8b, 24},
	{0xffff8c, 24},
	{0xffff8d, 24},
	{0xffff8e, 24},
	{0xffff8f, 24},
	{0xffff90, 24},
	{0xffff91, 24},
	{0xffff92, 24},
	{0xffff93, 24},
	{0xffff94, 24},
	{0xffff95, 24},
	{0xffff96, 24},
	{0xffff97, 24},
	{0xffff98, 24},
	{0xffff99, 24},
	{0xffff9a, 24},
	{0xffff9b, 24},
	{0xffff9c, 24},
	{0xffff9d, 24},
	{0xffff9e, 24},
	{0xffff9f, 24},
	{0xffffa0, 24},
	{0xffffa1, 24},
	{0xffffa2, 24},
	{0xffffa3, 24},
	{0xffffa4, 24},
	{0xffffa5, 24},
	{0xffffa6, 24},
	{0xffffa7, 24},
	{0xffffa8, 24},
	{0xffffa9, 24},
	{0xffffaa, 24},
	{0xffffab, 24},
	{0xffffac, 24},
	{0xffffad, 24},
	{0xffffae, 24},
	{0xffffaf, 24},
	{0xffffb0, 24},
	{0xffffb1, 24},
	{0xffffb2, 24},
	{0xffffb3, 24},
	{0xffffb4, 24},
	{0xffffb5, 24},
	{0xffffb6, 24},
	{0xffffb7, 24},
	{0xffffb8, 24},
	{0xffffb9, 24},
	{0xffffba, 24},
	{0xffffbb, 24},
	{0xffffbc, 24},
	{0xffffbd, 24},
	{0xffffbe, 24},
	{0xffffbf, 24},
	{0xffffc0, 24},
	{0xffffc1, 24},
	{0xffffc2, 24},
	{0xffffc3, 24},
	{0xffffc4, 24},
	{0xffffc5, 24},
	{0xffffc6, 24},
	{0xffffc7, 24},
	{0xffffc8, 24},
	{0xffffc9, 24},
	{0xffffca, 24},
	{0xffffcb, 24},
	{0xffffcc, 24},
	{0xffffcd, 24},
	{0xffffce, 24},
	{0xffffcf, 24},
	{0xffffd0, 24},
	{0xffffd1, 24},
	{0xffffd2, 24},
	{0xffffd3, 24},
	{0xffffd4, 24},
	{0xffffd5, 24},
	{0xffffd6, 24},
	{0xffffd7, 24},
	{0xffffd8, 24},
	{0xffffd9, 24},
	{0xffffda, 24},
	{0xffffdb, 24},
	{0xffffdc, 24},
	{0xffffdd, 24},
}
