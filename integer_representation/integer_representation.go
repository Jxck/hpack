package integer_representation

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
func Encode(I uint64, N uint8) (buf []byte) {
	boundary := uint64(1<<N - 1) // 2^N-1

	if I < boundary {
		// If I < 2^N - 1, encode I on N bits
		buf = append(buf, byte(I))
	} else {
		// encode 2^N - 1 on N bits
		buf = append(buf, byte(boundary))

		// I = I - (2^N - 1)
		I = I - boundary

		// While I >= 128
		for I >= 128 {
			// Encode (I % 128 + 128) on 8 bits
			buf = append(buf, byte(I%128+128))

			// I = I / 128
			I = I / 128
		}

		// encode (I) on 8 bits
		buf = append(buf, byte(I))
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
func Decode(buf []byte, N uint8) uint64 {
	boundary := uint64(1<<N - 1) // 2^N-1
	I := uint64(buf[0])          // Read N bit from first 1 byte as I
	if I < boundary {            // less than 2^N-1
		return I // as is
	}
	for i, b := range buf[1:] { // continue while follow bites are bigger than 128
		shift := uint8(7 * i)
		if b >= 128 { // if first bit is 1
			// to 0 at first bit (- 128) and shift 7*i bit
			// and add
			I += uint64(b-128) << shift
		} else { // if first bit is 0
			// shit 7*i shift
			// and add
			I += uint64(b) << shift
			break
		}
	}
	return I
}

// read prefixed N bytes from buffer
// if N bit of first byte is 2^N-1 (ex 1111 in N=4)
// read follow byte until it's smaller than 128
func ReadPrefixedInteger(buf []byte, N uint8) []byte {
	boundary := byte(1<<N - 1) // 2^N-1
	first := buf[0]

	first = first & boundary // mask N bit
	prefix := []byte{first}

	// if first byte is smaller than boundary
	// it's end of the prefixed bytes
	if first < boundary {
		return prefix
	}

	// read bytes while bytes smaller than 128
	for i := 1; ; i++ {
		tmp := buf[i]
		prefix = append(prefix, tmp)
		if tmp < 128 {
			break
		}
	}

	return prefix
}
