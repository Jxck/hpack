package hpack

type Header struct {
	Name  string
	Value string
}

func (h *Header) Size() int {
	return len(h.Name) + len(h.Value) + 32
}

type Headers []Header

var DEFAULT_HEADER_TABLE_SIZE int = 4096

type HeaderTable struct {
	HEADER_TABLE_SIZE int
	Headers
}

func (ht *HeaderTable) Size() int {
	var sum int
	for _, h := range ht.Headers {
		sum += h.Size()
	}
	return sum
}

func (ht *HeaderTable) Add(name, value string) {
	header := Header{name, value}
	if header.Size() > ht.HEADER_TABLE_SIZE {
		ht.DeleteAll()
	} else {
		ht.AllocSpace(header.Size())
		ht.Headers = append(ht.Headers, header)
	}
}

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

// name と value が HeaderTable にあるかを探す
// name, value とも一致 => index, *Header
// name はある          => index, nil
// ない                 => -1, nil
func (ht HeaderTable) SearchHeader(name, value string) (int, *Header) {
	// name が複数一致した時のために格納しておく
	// MEMO: スライスで持たず単一で最初だけもってもいいかもしれないが
	// もし無かった場合 0 になって、それが index=0 と紛らわしいので
	// slice でもって、長さで判断できるようにした
	var matching_name_indexes = []int{}

	// ヘッダテーブルの頭から探す
	for i, h := range ht.Headers {

		// Name がヘッダテーブルにあった場合
		if h.Name == name {

			// Value も一致したら
			if h.Value == value {
				// 一致した index とそこにある値を返す
				return i, &h // index header
			}

			// name は一致したのでそのインデックスを加えておく
			matching_name_indexes = append(matching_name_indexes, i)
		}
	}

	// Name があっても value までは一致しなかった場合
	// 一番最初のヘッダを返す
	if len(matching_name_indexes) > 0 {
		return matching_name_indexes[0], nil // literal with index
	}

	// Name も一致しなかったら -1, nil
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