package hpack

import ()

// TODO: what is default ?
const DEFAULT_HEADER_TABLE_SIZE int = 4096

// The header table is a component used to associate stored header
// fields to index values.
// The data stored in this table is in first-in, first-out order.
type HeaderTable struct {
	HEADER_TABLE_SIZE int
	HeaderFields      []*HeaderField
}

func NewHeaderTable(SETTINGS_HEADER_TABLE_SIZE int) *HeaderTable {
	return &HeaderTable{
		SETTINGS_HEADER_TABLE_SIZE,
		[]*HeaderField{},
	}
}

// get total size of Header Table
func (ht *HeaderTable) Size() int {
	var size int
	for _, h := range ht.HeaderFields {
		size += h.Size()
	}
	return size
}

// get length of Header Table
func (ht *HeaderTable) Len() int {
	return len(ht.HeaderFields)
}

// push new Header Field to the of HeaderTable
// with eviction
// :TODO (check & test eviction more)
func (ht *HeaderTable) Push(hf *HeaderField) {
	tmp := []*HeaderField{hf}
	ht.HeaderFields = append(tmp, ht.HeaderFields...)
}

/*
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
*/
