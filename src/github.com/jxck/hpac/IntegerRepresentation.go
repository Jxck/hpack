package hpac

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
//     While I >= 128
//          Encode (I % 128 + 128) on 8 bits
//          I = I / 128
//     encode (I) on 8 bits
func EncodeInteger(I int, N int) *bytes.Buffer {
	buf := new(bytes.Buffer)

	// 2^N -1
	boundary := int(math.Pow(2, float64(N))) - 1

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
// [sample]
// 40 [31, 9]
// b xxx1 1111    40>31 : e(31), I=40-31=9
// a xxxx 1001    9 <128: e(9)
//
// b) 9 + 31 = 40
//
//
// 1337 [31, 154, 10]
// a xxx1 1111    1337>31 : e(31), I=1337-31=1306
// b 1001 1010    1306>128: e(1306%128+128), I=1306/128=10
// c 0000 1010    10  <128: 3(10)
//
// b) (10*128) + (154-128) = 1306
// a) 1306 + 31 = 1337
//
//
// 3000000 [31 161 141 183 1]
// a xxx1 1111    3000000>31  : e(31), I=3000000-31=2999969
// b 1010 0001    2999969>128 : e(2999969%128+128), I=2999969/128=23437
// c 1000 1101      23437>128 : e(23437%128+128), I=23437/128=183
// d 1011 0111        183>128 : e(183%128+128), I=183/128=1
// e 0000 0001          1<128 : e(1)
//
// d) (1*128) + (183-128) = 183
// c) (183*128) + (141-128) = 23473
// b) (23437*128) + (161-128) = 2999969
// a) 2999969 + 31 = 3000000
func DecodeInteger(buf []byte, N uint8) uint32 {
	boundary := byte(math.Pow(2, float64(N)) - 1)
	if buf[0] == boundary {
		var I uint32
		i := len(buf) - 1
		I += uint32(buf[i])
		for i > 1 {
			I *= 128
			i--
			I += uint32(buf[i]) - 128
		}
		I += 31

		return I
	}
	return uint32(buf[0])
}
