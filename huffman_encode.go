package hpack

import (
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

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
