package hpac

type Header struct {
	Name  string
	Value string
}

type HeaderTable []Header

func (ht *HeaderTable) Add(header Header) {
	*ht = append(*ht, header)
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
	for i, h := range ht {

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

// func Search(headers http.Header, headerTable HeaderTable) {
// 	for name, values := range headers {
// 		for _, value := range values {
// 			index, h := headerTable.SearchHeader(name, value)
// 			if h != nil {
// 				frame := hpac.NewIndexedHeader()
// 				frame.Index = uint64(index)
// 				f := hpac.EncodeHeader(frame)
// 				log.Printf("indexed header [%v:%v] is in HT[%v]=%v  %v", name, value, index, h, f.Bytes())
// 			} else if index != -1 {
// 				frame := hpac.NewIndexedNameWithIncrementalIndexing()
// 				frame.Index = uint64(index)
// 				frame.ValueLength = uint64(len(value))
// 				frame.ValueString = value
// 				f := hpac.EncodeHeader(frame)
// 				log.Printf("literal with index [%v:%v] is in HT[%v] %v", name, value, index, f.Bytes())
// 			} else {
// 				frame := hpac.NewNewNameWithoutIndexing()
// 				frame.NameLength = uint64(len(name))
// 				frame.NameString = name
// 				frame.ValueLength = uint64(len(value))
// 				frame.ValueString = value
// 				f := hpac.EncodeHeader(frame)
// 				log.Printf("literal without index [%v:%v] is not in HT %v", name, value, f.Bytes())
// 			}
// 		}
// 	}
// }

func NewRequestHeaderTable() HeaderTable {
	return HeaderTable{
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
	}
}

func NewResponseHeaderTable() HeaderTable {
	return HeaderTable{
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
	}
}
