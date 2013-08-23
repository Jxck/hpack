package hpac

import (
	"bytes"
	"encoding/binary"
	"log"
)

type Frame struct {
	Flag1       uint8
	Flag2       uint8
	Flag3       uint8
	Index       uint8
	ValueLength uint8
	ValueString string
}

func DecodeHeader(buf *bytes.Buffer) *Frame {
	var types uint8
	if err := binary.Read(buf, binary.BigEndian, &types); err != nil {
		log.Println("binary.Read failed:", err)
	}
	log.Printf("%b\n", types)
	if types>>7 == 1 {

		// 	0   1   2   3   4   5   6   7
		// +---+---+---+---+---+---+---+---+
		// | 1 |        Index (7+)         |
		// +---+---------------------------+
		log.Println("Indexed Header Representation")

	} else if types == 0 {

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

	} else if types>>5 == 0x2 {

		// 0   1   2   3   4   5   6   7
		// +---+---+---+---+---+---+---+---+
		// | 0 | 1 | 0 |    Index (5+)     |
		// +---+---+---+-------------------+
		// |       Value Length (8+)       |
		// +-------------------------------+
		// | Value String (Length octets)  |
		// +-------------------------------+

		// 0x44      (literal header with incremental indexing, name index = 3)
		// 0x16      (header value string length = 22)
		// /my-example/index.html

		var frame = &Frame{}

		frame.Flag1 = types >> 7
		frame.Flag2 = (types & 0x40) >> 6
		frame.Flag3 = (types & 0x20) >> 5
		frame.Index = (types & 0x1F) - 1

		binary.Read(buf, binary.BigEndian, &frame.ValueLength) // err

		valueBytes := make([]byte, frame.ValueLength)

		binary.Read(buf, binary.BigEndian, &valueBytes) // err

		frame.ValueString = string(valueBytes)

		log.Println("Literal Header with Incremental Indexing - Indexed Name")
		log.Println("&v", frame)

		return frame

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
	return nil
}
