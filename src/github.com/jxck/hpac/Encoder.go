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
		// Else, encode 2^N - 1 on N bits and do the following steps:
		err := binary.Write(buf, binary.BigEndian, uint8(boundary))
		if err != nil {
			log.Println("binary.Write failed:", err)
		}

		// Set I to (I - (2^N - 1)) and Q to 1
		I = I - boundary
		Q := 1
		R := 0

		for Q > 0 {
			// Compute Q and R, quotient and remainder of I divided by 2^7
			R = I % 128
			Q = (I - R) / 128
			log.Println(Q, R)

			// If Q is strictly greater than 0, write one 1 bit; otherwise, write one 0 bit
			var b uint8 = 0
			if Q > 0 {
				b = 128
			}

			b = b | uint8(R)
			err := binary.Write(buf, binary.BigEndian, b)
			if err != nil {
				log.Println("binary.Write failed:", err)
			}
			I = Q
		}
	}

	return buf
}
