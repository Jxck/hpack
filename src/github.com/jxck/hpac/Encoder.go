package hpac

import (
	"bytes"
	"encoding/binary"
	"log"
	"math"
)

type Encoder struct {
	requestHeaderTable  HeaderTable
	responseHeaderTable HeaderTable
	referenceSet        ReferenceSet
}

func NewEncoder() Encoder {
	var encoder = Encoder{
		requestHeaderTable:  RequestHeaderTable,
		responseHeaderTable: ResponseHeaderTable,
		referenceSet:        ReferenceSet{},
	}
	return encoder
}

func EncodeInteger(I int, N int) *bytes.Buffer {
	log.Printf("%v %v = ", I, N)
	buf := new(bytes.Buffer)

	// 2^N -1
	boundary := int(math.Pow(2, float64(N))) - 1

	if I < boundary {
		// If I < 2^N - 1, encode I on N bits
		err := binary.Write(buf, binary.BigEndian, uint8(I))
		if err != nil {
			log.Println("binary.Write failed:", err)
		}
	} else {
		// encode 2^N - 1 on N bits
		err := binary.Write(buf, binary.BigEndian, uint8(boundary))
		if err != nil {
			log.Println("binary.Write failed:", err)
		}

		// I = I - (2^N - 1)
		I = I - boundary

		// While I >= 128
		for I >= 128 {
			// Encode (I % 128 + 128) on 8 bits
			err := binary.Write(buf, binary.BigEndian, uint8(I%128+128))
			if err != nil {
				log.Println("binary.Write failed:", err)
			}
			// I = I / 128
			I = I / 128
		}

		// encode (I) on 8 bits
		err = binary.Write(buf, binary.BigEndian, uint8(I))
		if err != nil {
			log.Println("binary.Write failed:", err)
		}
	}
	log.Println(buf.Bytes())
	return buf
}

func DecodeInteger(buf []byte, N uint8) uint32 {
	boundary := byte(math.Pow(2, float64(N)) - 1)
	if buf[0] == boundary {
		l := len(buf) - 1
		I := uint32(buf[l]) * 128
		for l--; l > 1; l-- {
			I = (I + uint32(buf[l]) - 128) * 128
		}
		I = I + uint32(buf[1]) - 128 + uint32(buf[0])
		return uint32(I)
	}
	return uint32(buf[0])
}
