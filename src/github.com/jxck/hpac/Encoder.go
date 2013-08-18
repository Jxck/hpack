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

	// If I < 2^N - 1, encode I on N bits
	if I < int(math.Pow(2, float64(N)))-1 {
		err := binary.Write(buf, binary.BigEndian, uint8(I))
		if err != nil {
			log.Println("binary.Write failed:", err)
		}
	}



	return buf
}
