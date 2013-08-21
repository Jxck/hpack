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
	log.Printf("0%b\n", types)
	if types&0x80 == 1 {
		// 	0   1   2   3   4   5   6   7
		// +---+---+---+---+---+---+---+---+
		// | 1 |        Index (7+)         |
		// +---+---------------------------+
		log.Println("Indexed Header Representation")
	} else {
		if types == 0 {

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

		} else if types == 0x40 {

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

		} else if types == 0x60 {

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

		} else if types&0x60 == 0x40 {
			// 0   1   2   3   4   5   6   7
			// +---+---+---+---+---+---+---+---+
			// | 0 | 1 | 0 |    Index (5+)     |
			// +---+---+---+-------------------+
			// |       Value Length (8+)       |
			// +-------------------------------+
			// | Value String (Length octets)  |
			// +-------------------------------+
			log.Println("Literal Header with Incremental Indexing - Indexed Name")

		} else if types&0x60 == 0x60 {

			//   0   1   2   3   4   5   6   7
			// +---+---+---+---+---+---+---+---+
			// | 0 | 1 | 1 |    Index (5+)     |
			// +---+---+---+-------------------+
			// |       Value Length (8+)       |
			// +-------------------------------+
			// | Value String (Length octets)  |
			// +-------------------------------+
			log.Println("Literal Header without Indexing - Indexed Name")

		} else {

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

		}
	}
}

func main() {
	log.SetFlags(log.Lshortfile)
	var buf *bytes.Buffer
	var d uint8

	d = 0x44 //     (literal header with incremental indexing, name index = 3)
	buf = new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, d)
	DecodeHeader(buf)

	//d = 0x4D //     (literal header with incremental indexing, name index = 12)
	//buf = new(bytes.Buffer)
	//binary.Write(buf, binary.BigEndian, d)
	//DecodeHeader(buf)

	//d = 0x40 //     (literal header with incremental indexing, new name)
	//buf = new(bytes.Buffer)
	//binary.Write(buf, binary.BigEndian, d)
	//DecodeHeader(buf)

	//d = 0xa6 //      (indexed header, index = 38: removal from reference set)
	//buf = new(bytes.Buffer)
	//binary.Write(buf, binary.BigEndian, d)
	//DecodeHeader(buf)

	//d = 0xa8 //      (indexed header, index = 40: removal from reference set)
	//buf = new(bytes.Buffer)
	//binary.Write(buf, binary.BigEndian, d)
	//DecodeHeader(buf)

	//d = 0x04 //      (literal header, substitution indexing, name index = 3)
	//buf = new(bytes.Buffer)
	//binary.Write(buf, binary.BigEndian, d)
	//DecodeHeader(buf)

	//d = 0x26 //      (replaced entry index = 38)
	//buf = new(bytes.Buffer)
	//binary.Write(buf, binary.BigEndian, d)
	//DecodeHeader(buf)

	//0x5f
	//0x0a  (literal header, incremental indexing, name index = 40)
	//0x06       (header value string length = 6)
}
