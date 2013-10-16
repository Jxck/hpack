package hpack

import (
	"bytes"
	"encoding/binary"
	"log"
	"math"
)

// Encode Integer to N bit prefix
// Integer Representation
//
// [Logic]
// If I < 2^N - 1, encode I on N bits
// Else
//     encode 2^N - 1 on N bits
//     I = I - (2^N - 1)
//     While I >= 128
//          Encode (I % 128 + 128) on 8 bits
//          I = I / 128
//     encode (I) on 8 bits
func EncodeInteger(I uint64, N uint8) *bytes.Buffer {
	buf := new(bytes.Buffer)

	// 2^N -1
	boundary := uint64(math.Pow(2, float64(N))) - 1

	if I < boundary {
		// If I < 2^N - 1, encode I on N bits
		err := binary.Write(buf, binary.BigEndian, uint8(I))
		if err != nil {
			log.Fatal("binary.Write failed:", err)
		}
	} else {
		// encode 2^N - 1 on N bits
		err := binary.Write(buf, binary.BigEndian, uint8(boundary))
		if err != nil {
			log.Fatal("binary.Write failed:", err)
		}

		// I = I - (2^N - 1)
		I = I - boundary

		// While I >= 128
		for I >= 128 {
			// Encode (I % 128 + 128) on 8 bits
			err := binary.Write(buf, binary.BigEndian, uint8(I%128+128))
			if err != nil {
				log.Fatal("binary.Write failed:", err)
			}
			// I = I / 128
			I = I / 128
		}

		// encode (I) on 8 bits
		err = binary.Write(buf, binary.BigEndian, uint8(I))
		if err != nil {
			log.Fatal("binary.Write failed:", err)
		}
	}
	return buf
}

// Decode N bit prefixed Representation
// to Integer
//
// Read N bit from first 1 byte as I
// If I < 2^N-1
//     decode I
// Else
//     i = 0
//     read next 1 byte as b
//     While b > 128
//         I += (b - 128) * 128^(i-1)
//         i++
func DecodeInteger(buf []byte, N uint8) uint64 {
	boundary := uint64(math.Pow(2, float64(N)) - 1)
	I := uint64(buf[0])
	if I < boundary {
		return I
	}
	for i, b := range buf[1:] {
		i++
		if b > 128 {
			shift := uint64(math.Pow(128, float64(i-1)))
			I += ((uint64(b) - 128) * shift)
		} else {
			shift := uint64(math.Pow(128, float64(i-1)))
			I += (uint64(b) * shift)
			break
		}
	}
	return I
}

// read prefixed bytes from buffer
func ReadPrefixedInteger(buf *bytes.Buffer, N uint8) *bytes.Buffer {
	var tmp uint8
	boundary := byte(math.Pow(2, float64(N)) - 1)
	binary.Read(buf, binary.BigEndian, &tmp) // err

	tmp = tmp & boundary
	prefix := bytes.NewBuffer([]byte{tmp})

	// if first byte is smaller than boundary
	// it's end of the prefixed bytes
	if tmp < boundary {
		return prefix
	}

	// read bytes while bytes smaller than 128
	for {
		binary.Read(buf, binary.BigEndian, &tmp) // err
		prefix.WriteByte(tmp)
		if tmp < 128 {
			break
		}
	}

	return prefix
}
