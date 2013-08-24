package hpac

import (
	"bytes"
	"encoding/binary"
	"log"
)

type Frame interface {
}

type IndexedHeader struct {
	Flag1 uint8
	Index uint8
}

// TODO: RENAME

type IncrementalIndexingName struct {
	Flag1       uint8
	Flag2       uint8
	Flag3       uint8
	Index       uint32
	ValueLength uint32
	ValueString string
}

type IncrementalNewName struct {
	Flag1       uint8
	Flag2       uint8
	Flag3       uint8
	Flag4       uint8
	NameLength  uint8
	NameString  string
	ValueLength uint8
	ValueString string
}

type SubstitutionIndexedName struct {
	Flag1            uint8
	Flag2            uint8
	Index            uint8
	SubstitutedIndex uint8
	ValueLength      uint8
	ValueString      string
}

func DecodeHeader(buf *bytes.Buffer) Frame {
	log.SetFlags(log.Lshortfile)
	var types uint8
	if err := binary.Read(buf, binary.BigEndian, &types); err != nil {
		log.Println("binary.Read failed:", err)
	}
	if types>>7 == 1 {

		// 	0   1   2   3   4   5   6   7
		// +---+---+---+---+---+---+---+---+
		// | 1 |        Index (7+)         |
		// +---+---------------------------+

		frame := &IndexedHeader{}
		frame.Flag1 = 1
		frame.Index = types & 0x7F

		log.Println("Indexed Header Representation")
		return frame

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

		frame := &IncrementalNewName{}
		frame.Flag1 = 0
		frame.Flag2 = 1
		frame.Flag3 = 0
		frame.Flag4 = 0

		binary.Read(buf, binary.BigEndian, &frame.NameLength) // err
		nameBytes := make([]byte, frame.NameLength)
		binary.Read(buf, binary.BigEndian, &nameBytes) // err
		frame.NameString = string(nameBytes)

		binary.Read(buf, binary.BigEndian, &frame.ValueLength) // err
		valueBytes := make([]byte, frame.ValueLength)
		binary.Read(buf, binary.BigEndian, &valueBytes) // err
		frame.ValueString = string(valueBytes)

		log.Println("Literal Header with Incremental Indexing - New Name")
		return frame

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

		var frame = &IncrementalIndexingName{}

		frame.Flag1 = 0
		frame.Flag2 = 1
		frame.Flag3 = 0

		// unread first byte for parse frame
		buf.UnreadByte()

		frame.Index = DecodePrefixedInteger(buf, 5) - 1

		frame.ValueLength = DecodePrefixedInteger(buf, 8)
		valueBytes := make([]byte, frame.ValueLength)
		binary.Read(buf, binary.BigEndian, &valueBytes) // err
		frame.ValueString = string(valueBytes)

		log.Println("Literal Header with Incremental Indexing - Indexed Name")
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
		var frame = &SubstitutionIndexedName{}

		frame.Flag1 = 0
		frame.Flag2 = 0
		frame.Index = (types & 0x1F) - 1

		binary.Read(buf, binary.BigEndian, &frame.SubstitutedIndex) // err
		binary.Read(buf, binary.BigEndian, &frame.ValueLength)      // err
		valueBytes := make([]byte, frame.ValueLength)
		binary.Read(buf, binary.BigEndian, &valueBytes) // err
		frame.ValueString = string(valueBytes)

		log.Println("Literal Header with Substitution Indexing - Indexed Name")
		return frame

	}
	return nil
}
