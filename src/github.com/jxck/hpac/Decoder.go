package hpac

import (
	"bytes"
	"encoding/binary"
	"log"
)

type Frame interface {
}

// Indexed Header Representation
type IndexedHeader struct {
	Flag1 uint8
	Index uint8
}

// TODO: RENAME

// Literal Header without Indexing - New Name
type NewNameWithoutIndexing struct {
	Flag1       uint8
	Flag2       uint8
	Flag3       uint8
	Index       uint32
	NameLength  uint32
	NameString  string
	ValueLength uint32
	ValueString string
}

// Literal Header without Indexing - Indexed Name
type IndexedNameWithoutIndexing struct {
}

// Literal Header with Incremental Indexing - Indexed Name
type IndexedNameWithIncrementalIndexing struct {
	Flag1       uint8
	Flag2       uint8
	Flag3       uint8
	Index       uint32
	ValueLength uint32
	ValueString string
}

// Literal Header with Incremental Indexing - New Name
type NewNameWithIncrementalIndexing struct {
	Flag1       uint8
	Flag2       uint8
	Flag3       uint8
	Index       uint8
	NameLength  uint32
	NameString  string
	ValueLength uint32
	ValueString string
}

// Literal Header with Substitution Indexing - Indexed Name
type IndexedNameWithSubstitutionIndexing struct {
	Flag1            uint8
	Flag2            uint8
	Index            uint32
	SubstitutedIndex uint32
	ValueLength      uint32
	ValueString      string
}

// Literal Header with Substitution Indexing - New Name
type NewNameWithSubstitutionIndexing struct {
	Flag1            uint8
	Flag2            uint8
	Index            uint8
	NameLength       uint32
	NameString       string
	SubstitutedIndex uint32
	ValueLength      uint32
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

		frame := &NewNameWithSubstitutionIndexing{}
		frame.Flag1 = 0
		frame.Flag2 = 0
		frame.Index = 0
		frame.NameLength = DecodePrefixedInteger(buf, 8)
		frame.NameString = DecodeString(buf, frame.NameLength)
		frame.SubstitutedIndex = DecodePrefixedInteger(buf, 8)
		frame.ValueLength = DecodePrefixedInteger(buf, 8)
		frame.ValueString = DecodeString(buf, frame.ValueLength)

		log.Println("Literal Header with Substitution Indexing - New Name")
		return frame

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

		frame := &NewNameWithIncrementalIndexing{}
		frame.Flag1 = 0
		frame.Flag2 = 1
		frame.Flag3 = 0
		frame.Index = 0
		frame.NameLength = DecodePrefixedInteger(buf, 8)
		frame.NameString = DecodeString(buf, frame.NameLength)
		frame.ValueLength = DecodePrefixedInteger(buf, 8)
		frame.ValueString = DecodeString(buf, frame.ValueLength)

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

		var frame = &NewNameWithoutIndexing{}
		frame.Flag1 = 0
		frame.Flag2 = 1
		frame.Flag3 = 1
		frame.Index = 0
		frame.NameLength = DecodePrefixedInteger(buf, 8)
		frame.NameString = DecodeString(buf, frame.NameLength)
		frame.ValueLength = DecodePrefixedInteger(buf, 8)
		frame.ValueString = DecodeString(buf, frame.ValueLength)

		log.Println("Literal Header without Indexing - New Name")
		return frame

	} else if types>>5 == 0x2 {

		// 0   1   2   3   4   5   6   7
		// +---+---+---+---+---+---+---+---+
		// | 0 | 1 | 0 |    Index (5+)     |
		// +---+---+---+-------------------+
		// |       Value Length (8+)       |
		// +-------------------------------+
		// | Value String (Length octets)  |
		// +-------------------------------+

		// unread first byte for parse frame
		buf.UnreadByte()

		var frame = &IndexedNameWithIncrementalIndexing{}
		frame.Flag1 = 0
		frame.Flag2 = 1
		frame.Flag3 = 0
		frame.Index = DecodePrefixedInteger(buf, 5) - 1
		frame.ValueLength = DecodePrefixedInteger(buf, 8)
		frame.ValueString = DecodeString(buf, frame.ValueLength)

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

		// unread first byte for parse frame
		buf.UnreadByte()

		var frame = &IndexedNameWithSubstitutionIndexing{}
		frame.Flag1 = 0
		frame.Flag2 = 0
		frame.Index = DecodePrefixedInteger(buf, 6) - 1
		frame.SubstitutedIndex = DecodePrefixedInteger(buf, 8)
		frame.ValueLength = DecodePrefixedInteger(buf, 8)
		frame.ValueString = DecodeString(buf, frame.ValueLength)

		log.Println("Literal Header with Substitution Indexing - Indexed Name")
		return frame

	}
	return nil
}

func DecodeString(buf *bytes.Buffer, n uint32) string {
	valueBytes := make([]byte, n)
	binary.Read(buf, binary.BigEndian, &valueBytes) // err
	return string(valueBytes)
}
