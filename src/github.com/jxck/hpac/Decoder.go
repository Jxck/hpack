package main

import (
	"bytes"
	"encoding/binary"
	"log"
)

func DecodeHeader(buf *bytes.Buffer) {
	var types uint8
	if err := binary.Read(buf, binary.BigEndian, &types); err != nil {
		log.Println("binary.Read failed:", err)
	}
	log.Printf("%b\n", types)
	log.Printf("%b\n", types&0x80)
	if types&0x80 != 1 {
		// 	0   1   2   3   4   5   6   7
		// +---+---+---+---+---+---+---+---+
		// | 1 |        Index (7+)         |
		// +---+---------------------------+
		log.Println("Indexed Header Representation")
	} else {

		// 0   1   2   3   4   5   6   7
		// +---+---+---+---+---+---+---+---+
		// | 0 | 0 |           0           |
		// +---+---+-----------------------+
		// |       Name Length (8+)        |
		// +-------------------------------+
		// |  Name String (Length octets)  |
		// +-------------------------------+
		// |    Substituted Index (8+)     |
		// +-------------------------------+
		// |       Value Length (8+)       |
		// +-------------------------------+
		// | Value String (Length octets)  |
		// +-------------------------------+
		log.Println("Literal Header with Substitution Indexing - New Name")

		// 0   1   2   3   4   5   6   7
		// +---+---+---+---+---+---+---+---+
		// | 0 | 0 |      Index (6+)       |
		// +---+---+-----------------------+
		// |    Substituted Index (8+)     |
		// +-------------------------------+
		// |       Value Length (8+)       |
		// +-------------------------------+
		// | Value String (Length octets)  |
		// +-------------------------------+
		log.Println("Literal Header with Substitution Indexing - Indexed Name")

		// 0   1   2   3   4   5   6   7
		// +---+---+---+---+---+---+---+---+
		// | 0 | 1 | 0 |         0         |
		// +---+---+---+-------------------+
		// |       Name Length (8+)        |
		// +-------------------------------+
		// |  Name String (Length octets)  |
		// +-------------------------------+
		// |       Value Length (8+)       |
		// +-------------------------------+
		// | Value String (Length octets)  |
		// +-------------------------------+
		log.Println("Literal Header with Incremental Indexing - New Name")

		// 0   1   2   3   4   5   6   7
		// +---+---+---+---+---+---+---+---+
		// | 0 | 1 | 0 |    Index (5+)     |
		// +---+---+---+-------------------+
		// |       Value Length (8+)       |
		// +-------------------------------+
		// | Value String (Length octets)  |
		// +-------------------------------+
		log.Println("Literal Header with Incremental Indexing - Indexed Name")

		// 0   1   2   3   4   5   6   7
		// +---+---+---+---+---+---+---+---+
		// | 0 | 1 | 1 |         0         |
		// +---+---+---+-------------------+
		// |       Name Length (8+)        |
		// +-------------------------------+
		// |  Name String (Length octets)  |
		// +-------------------------------+
		// |       Value Length (8+)       |
		// +-------------------------------+
		// | Value String (Length octets)  |
		// +-------------------------------+
		log.Println("Literal Header without Indexing - New Name")

		//   0   1   2   3   4   5   6   7
		// +---+---+---+---+---+---+---+---+
		// | 0 | 1 | 1 |    Index (5+)     |
		// +---+---+---+-------------------+
		// |       Value Length (8+)       |
		// +-------------------------------+
		// | Value String (Length octets)  |
		// +-------------------------------+
		log.Println("Literal Header without Indexing - Indexed Name")

	}
}

func main() {
	buf := new(bytes.Buffer)

	var d uint8 = 0x44
	log.Printf("%b\n", d)

	err := binary.Write(buf, binary.BigEndian, d)
	if err != nil {
		log.Println("binary.Write failed:", err)
	}

	DecodeHeader(buf)
}
