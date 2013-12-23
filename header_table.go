package hpack

type Headers []Header

var DEFAULT_HEADER_TABLE_SIZE int = 4096

type HeaderTable struct {
	HEADER_TABLE_SIZE int
	Headers
}

// get total size of Header Table
func (ht *HeaderTable) Size() int {
	var sum int
	for _, h := range ht.Headers {
		sum += h.Size()
	}
	return sum
}

// add name value pair to the end of HeaderTable
// with eviction :TODO (check & test eviction more)
func (ht *HeaderTable) Add(name, value string) {
	header := Header{name, value}
	if header.Size() > ht.HEADER_TABLE_SIZE {
		ht.DeleteAll()
	} else {
		ht.AllocSpace(header.Size())
		ht.Headers = append(ht.Headers, header)
	}
}

// replace Header at index i with name, value pair
// with eviction :TODO (check & test eviction more)
func (ht *HeaderTable) Replace(name, value string, i uint64) {
	index := int(i)
	header := Header{name, value}
	if header.Size() > ht.HEADER_TABLE_SIZE {
		ht.DeleteAll()
	} else {
		existHeader := ht.Headers[index]
		needSpace := header.Size() - existHeader.Size()

		// if new replaced header is bigger than existing header
		// allocate space for replace
		if needSpace > 0 {
			removed := ht.AllocSpace(needSpace)
			// if Allocate removes entry
			// need to shift replace index
			index -= removed
			if index < 0 {
				// if replace entry removed
				// insert at first
				index = 0
			}
		}
		ht.Headers[index] = header
	}
}

// remove Header at index i
func (ht *HeaderTable) Remove(index int) {
	// https://code.google.com/p/go-wiki/wiki/SliceTricks
	copy(ht.Headers[index:], ht.Headers[index+1:])
	// avoid memory leak
	ht.Headers[len(ht.Headers)-1] = Header{}
	ht.Headers = ht.Headers[:len(ht.Headers)-1]
}

// removing entry from top
// until make space of size in Header Table
func (ht *HeaderTable) AllocSpace(size int) (removed int) {
	adjustSize := ht.HEADER_TABLE_SIZE - size
	for ht.Size() > adjustSize {
		ht.Remove(0)
		removed++
	}
	return
}

// remove all entry from HeaderTable
func (ht *HeaderTable) DeleteAll() {
	ht.Headers = Headers{}
}

// search name & value is exists in HeaderTable
// name, value   exists => index, *Header
// name          exists => index, nil
// none                 =>    -1, nil
func (ht HeaderTable) SearchHeader(name, value string) (int, *Header) {
	// name が複数一致した時のために格納しておく
	// MEMO: スライスで持たず単一で最初だけもってもいいかもしれないが
	// もし無かった場合 0 になって、それが index=0 と紛らわしいので
	// slice でもって、長さで判断できるようにした
	var matching_name_indexes = []int{}

	// search from header
	for i, h := range ht.Headers {

		// name exists
		if h.Name == name {

			// value exists
			if h.Value == value {
				return i, &h // index, *header
			}

			// only name exists
			// add the index of entry for multi hit
			matching_name_indexes = append(matching_name_indexes, i)
		}
	}

	// only name exists
	// return first muched index
	if len(matching_name_indexes) > 0 {
		return matching_name_indexes[0], nil // literal with index
	}

	// dosen't exists
	return -1, nil // literal without index
}

func NewRequestHeaderTable() HeaderTable {
	return HeaderTable{
		DEFAULT_HEADER_TABLE_SIZE,
		Headers{
			{":scheme", "http"},
			{":scheme", "https"},
			{":host", ""},
			{":path", "/"},
			{":method", "GET"},
			{"accept", ""},
			{"accept-charset", ""},
			{"accept-encoding", ""},
			{"accept-language", ""},
			{"cookie", ""},
			{"if-modified-since", ""},
			{"user-agent", ""},
			{"referer", ""},
			{"authorization", ""},
			{"allow", ""},
			{"cache-control", ""},
			{"connection", ""},
			{"content-length", ""},
			{"content-type", ""},
			{"date", ""},
			{"expect", ""},
			{"from", ""},
			{"if-match", ""},
			{"if-none-match", ""},
			{"if-range", ""},
			{"if-unmodified-since", ""},
			{"max-forwards", ""},
			{"proxy-authorization", ""},
			{"range", ""},
			{"via", ""},
		},
	}
}

func NewResponseHeaderTable() HeaderTable {
	return HeaderTable{
		DEFAULT_HEADER_TABLE_SIZE,
		Headers{
			{":status", "200"},
			{"age", ""},
			{"cache-control", ""},
			{"content-length", ""},
			{"content-type", ""},
			{"date", ""},
			{"etag", ""},
			{"expires", ""},
			{"last-modified", ""},
			{"server", ""},
			{"set-cookie", ""},
			{"vary", ""},
			{"via", ""},
			{"access-control-allow-origin", ""},
			{"accept-ranges", ""},
			{"allow", ""},
			{"connection", ""},
			{"content-disposition", ""},
			{"content-encoding", ""},
			{"content-language", ""},
			{"content-location", ""},
			{"content-range", ""},
			{"link", ""},
			{"location", ""},
			{"proxy-authenticate", ""},
			{"refresh", ""},
			{"retry-after", ""},
			{"strict-transport-security", ""},
			{"transfer-encoding", ""},
			{"www-authenticate", ""},
		},
	}
}
