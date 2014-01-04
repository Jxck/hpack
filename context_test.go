package hpack

import (
	. "github.com/jxck/color"
	. "github.com/jxck/logger"
	"reflect"
	"sort"
	"testing"
)

func TestRequestWithoutHuffman(t *testing.T) {
	var (
		context    Context
		buf        []byte
		expectedES *EmittedSet
		expectedHT *HeaderTable
		expectedRS ReferenceSet
	)

	context = NewContext(REQUEST, DEFAULT_HEADER_TABLE_SIZE)

	/**
	 * First Request
	 */
	Debug(Pink("\n========== First Request ==============="))

	buf = []byte{
		0x82, 0x87, 0x86, 0x04,
		0x0f, 0x77, 0x77, 0x77,
		0x2e, 0x65, 0x78, 0x61,
		0x6d, 0x70, 0x6c, 0x65,
		0x2e, 0x63, 0x6f, 0x6d,
	}

	expectedES = &EmittedSet{
		HeaderField{":method", "GET"},
		HeaderField{":scheme", "http"},
		HeaderField{":path", "/"},
		HeaderField{":authority", "www.example.com"},
	}

	expectedHT = NewHeaderTable(DEFAULT_HEADER_TABLE_SIZE)
	expectedHT.HeaderFields = []*HeaderField{
		NewHeaderField(":authority", "www.example.com"),
		NewHeaderField(":path", "/"),
		NewHeaderField(":scheme", "http"),
		NewHeaderField(":method", "GET"),
	}

	expectedRS = ReferenceSet{
		{NewHeaderField(":authority", "www.example.com"), false},
		{NewHeaderField(":path", "/"), false},
		{NewHeaderField(":scheme", "http"), false},
		{NewHeaderField(":method", "GET"), false},
	}

	context.Decode(buf)

	// test Header Table
	if context.HT.Size() != expectedHT.Size() {
		t.Errorf("\n got %v\nwant %v", context.HT.Size(), expectedHT.Size())
	}

	// test Header Table
	for i, hf := range expectedHT.HeaderFields {
		if !(*context.HT.HeaderFields[i] == *hf) {
			t.Errorf("\n got %v\nwant %v", *context.HT.HeaderFields[i], *hf)
		}
	}

	// test Emitted Set
	sort.Sort(context.ES)
	sort.Sort(expectedES)
	if !reflect.DeepEqual(context.ES, expectedES) {
		t.Errorf("\n got %v\nwant %v", context.ES, expectedES)
	}

	// test Reference Set
	if expectedRS.Len() != context.RS.Len() {
		t.Errorf("\n got %v\nwant %v", context.RS.Dump(), expectedRS.Dump())
	}

	for i, hf := range *context.RS {
		i := len(expectedRS) - i - 1
		if *expectedRS[i].HeaderField != *hf.HeaderField {
			t.Errorf("\n got %v\nwant %v", context.RS.Dump(), expectedRS.Dump())
		}
	}

	/**
	 * Second Request
	 */
	Debug(Pink("\n========== Second Request ==============="))

	buf = []byte{
		0x1b, 0x08, 0x6e, 0x6f,
		0x2d, 0x63, 0x61, 0x63,
		0x68, 0x65,
	}

	expectedES = &EmittedSet{
		HeaderField{":method", "GET"},
		HeaderField{":scheme", "http"},
		HeaderField{":path", "/"},
		HeaderField{":authority", "www.example.com"},
		HeaderField{"cache-control", "no-cache"},
	}

	expectedHT = NewHeaderTable(DEFAULT_HEADER_TABLE_SIZE)
	expectedHT.HeaderFields = []*HeaderField{
		NewHeaderField("cache-control", "no-cache"),
		NewHeaderField(":authority", "www.example.com"),
		NewHeaderField(":path", "/"),
		NewHeaderField(":scheme", "http"),
		NewHeaderField(":method", "GET"),
	}

	expectedRS = ReferenceSet{
		{NewHeaderField("cache-control", "no-cache"), false},
		{NewHeaderField(":authority", "www.example.com"), false},
		{NewHeaderField(":path", "/"), false},
		{NewHeaderField(":scheme", "http"), false},
		{NewHeaderField(":method", "GET"), false},
	}

	context.Decode(buf)

	// test Header Table
	if context.HT.Size() != expectedHT.Size() {
		t.Errorf("\n got %v\nwant %v", context.HT.Size(), expectedHT.Size())
	}

	// test Header Table
	for i, hf := range expectedHT.HeaderFields {
		if !(*context.HT.HeaderFields[i] == *hf) {
			t.Errorf("\n got %v\nwant %v", *context.HT.HeaderFields[i], *hf)
		}
	}

	// test Emitted Set
	sort.Sort(context.ES)
	sort.Sort(expectedES)
	if !reflect.DeepEqual(context.ES, expectedES) {
		t.Errorf("\n got %v\nwant %v", context.ES, expectedES)
	}

	// test Reference Set
	if expectedRS.Len() != context.RS.Len() {
		t.Errorf("\n got %v\nwant %v", context.RS.Dump(), expectedRS.Dump())
	}

	for i, hf := range *context.RS {
		i := len(expectedRS) - i - 1
		if *expectedRS[i].HeaderField != *hf.HeaderField {
			t.Errorf("\n got %v\nwant %v", context.RS.Dump(), expectedRS.Dump())
		}
	}

	/**
	 * Third Request
	 */
	Debug(Pink("\n========== Third Request ==============="))

	buf = []byte{
		0x80, 0x85, 0x8c, 0x8b,
		0x84, 0x00, 0x0a, 0x63,
		0x75, 0x73, 0x74, 0x6f,
		0x6d, 0x2d, 0x6b, 0x65,
		0x79, 0x0c, 0x63, 0x75,
		0x73, 0x74, 0x6f, 0x6d,
		0x2d, 0x76, 0x61, 0x6c,
		0x75, 0x65,
	}

	expectedES = &EmittedSet{
		HeaderField{":method", "GET"},
		HeaderField{":scheme", "https"},
		HeaderField{":path", "/index.html"},
		HeaderField{":authority", "www.example.com"},
		HeaderField{"custom-key", "custom-value"},
	}

	expectedHT = NewHeaderTable(DEFAULT_HEADER_TABLE_SIZE)
	expectedHT.HeaderFields = []*HeaderField{
		NewHeaderField("custom-key", "custom-value"),
		NewHeaderField(":path", "/index.html"),
		NewHeaderField(":scheme", "https"),
		NewHeaderField("cache-control", "no-cache"),
		NewHeaderField(":authority", "www.example.com"),
		NewHeaderField(":path", "/"),
		NewHeaderField(":scheme", "http"),
		NewHeaderField(":method", "GET"),
	}

	expectedRS = ReferenceSet{
		{NewHeaderField("custom-key", "custom-value"), true},
		{NewHeaderField(":authority", "www.example.com"), true},
		{NewHeaderField(":path", "/index.html"), true},
		{NewHeaderField(":scheme", "https"), true},
		{NewHeaderField(":method", "GET"), true},
	}

	context.Decode(buf)

	// test Header Table
	if context.HT.Size() != expectedHT.Size() {
		t.Errorf("\n got %v\nwant %v", context.HT.Size(), expectedHT.Size())
	}

	// test Header Table
	for i, hf := range expectedHT.HeaderFields {
		if !(*context.HT.HeaderFields[i] == *hf) {
			t.Errorf("\n got %v\nwant %v", *context.HT.HeaderFields[i], *hf)
		}
	}

	// test Emitted Set
	sort.Sort(context.ES)
	sort.Sort(expectedES)
	if !reflect.DeepEqual(context.ES, expectedES) {
		t.Errorf("\n got %v\nwant %v", context.ES, expectedES)
	}

	// test Reference Set
	if expectedRS.Len() != context.RS.Len() {
		t.Errorf("\n got %v\nwant %v", context.RS.Dump(), expectedRS.Dump())
	}

	for i, hf := range *context.RS {
		i := len(expectedRS) - i - 1
		if *expectedRS[i].HeaderField != *hf.HeaderField {
			t.Errorf("\n got %v\nwant %v", context.RS.Dump(), expectedRS.Dump())
		}
	}
}

func TestRequestWithHuffman(t *testing.T) {
	var (
		context    Context
		buf        []byte
		expectedES *EmittedSet
		expectedHT *HeaderTable
		expectedRS ReferenceSet
	)

	context = NewContext(REQUEST, DEFAULT_HEADER_TABLE_SIZE)

	/**
	 * First Request
	 */
	Debug(Pink("\n========== First Request ==============="))

	buf = []byte{
		0x82, 0x87, 0x86, 0x04,
		0x8b, 0xdb, 0x6d, 0x88,
		0x3e, 0x68, 0xd1, 0xcb,
		0x12, 0x25, 0xba, 0x7f,
	}

	expectedES = &EmittedSet{
		HeaderField{":method", "GET"},
		HeaderField{":scheme", "http"},
		HeaderField{":path", "/"},
		HeaderField{":authority", "www.example.com"},
	}

	expectedHT = NewHeaderTable(DEFAULT_HEADER_TABLE_SIZE)
	expectedHT.HeaderFields = []*HeaderField{
		NewHeaderField(":authority", "www.example.com"),
		NewHeaderField(":path", "/"),
		NewHeaderField(":scheme", "http"),
		NewHeaderField(":method", "GET"),
	}

	expectedRS = ReferenceSet{
		{NewHeaderField(":authority", "www.example.com"), false},
		{NewHeaderField(":path", "/"), false},
		{NewHeaderField(":scheme", "http"), false},
		{NewHeaderField(":method", "GET"), false},
	}

	context.Decode(buf)

	// test Header Table
	if context.HT.Size() != expectedHT.Size() {
		t.Errorf("\n got %v\nwant %v", context.HT.Size(), expectedHT.Size())
	}

	// test Header Table
	for i, hf := range expectedHT.HeaderFields {
		if !(*context.HT.HeaderFields[i] == *hf) {
			t.Errorf("\n got %v\nwant %v", *context.HT.HeaderFields[i], *hf)
		}
	}

	// test Emitted Set
	sort.Sort(context.ES)
	sort.Sort(expectedES)
	if !reflect.DeepEqual(context.ES, expectedES) {
		t.Errorf("\n got %v\nwant %v", context.ES, expectedES)
	}

	// test Reference Set
	if expectedRS.Len() != context.RS.Len() {
		t.Errorf("\n got %v\nwant %v", context.RS.Dump(), expectedRS.Dump())
	}

	for i, hf := range *context.RS {
		i := len(expectedRS) - i - 1
		if *expectedRS[i].HeaderField != *hf.HeaderField {
			t.Errorf("\n got %v\nwant %v", context.RS.Dump(), expectedRS.Dump())
		}
	}

	/**
	 * Second Request
	 */
	Debug(Pink("\n========== Second Request ==============="))

	buf = []byte{
		0x1b, 0x86, 0x63, 0x65,
		0x4a, 0x13, 0x98, 0xff,
	}

	expectedES = &EmittedSet{
		HeaderField{":method", "GET"},
		HeaderField{":scheme", "http"},
		HeaderField{":path", "/"},
		HeaderField{":authority", "www.example.com"},
		HeaderField{"cache-control", "no-cache"},
	}

	expectedHT = NewHeaderTable(DEFAULT_HEADER_TABLE_SIZE)
	expectedHT.HeaderFields = []*HeaderField{
		NewHeaderField("cache-control", "no-cache"),
		NewHeaderField(":authority", "www.example.com"),
		NewHeaderField(":path", "/"),
		NewHeaderField(":scheme", "http"),
		NewHeaderField(":method", "GET"),
	}

	expectedRS = ReferenceSet{
		{NewHeaderField("cache-control", "no-cache"), false},
		{NewHeaderField(":authority", "www.example.com"), false},
		{NewHeaderField(":path", "/"), false},
		{NewHeaderField(":scheme", "http"), false},
		{NewHeaderField(":method", "GET"), false},
	}

	context.Decode(buf)

	// test Header Table
	if context.HT.Size() != expectedHT.Size() {
		t.Errorf("\n got %v\nwant %v", context.HT.Size(), expectedHT.Size())
	}

	// test Header Table
	for i, hf := range expectedHT.HeaderFields {
		if !(*context.HT.HeaderFields[i] == *hf) {
			t.Errorf("\n got %v\nwant %v", *context.HT.HeaderFields[i], *hf)
		}
	}

	// test Emitted Set
	sort.Sort(context.ES)
	sort.Sort(expectedES)
	if !reflect.DeepEqual(context.ES, expectedES) {
		t.Errorf("\n got %v\nwant %v", context.ES, expectedES)
	}

	// test Reference Set
	if expectedRS.Len() != context.RS.Len() {
		t.Errorf("\n got %v\nwant %v", context.RS.Dump(), expectedRS.Dump())
	}

	for i, hf := range *context.RS {
		i := len(expectedRS) - i - 1
		if *expectedRS[i].HeaderField != *hf.HeaderField {
			t.Errorf("\n got %v\nwant %v", context.RS.Dump(), expectedRS.Dump())
		}
	}

	/**
	 * Third Request
	 */
	Debug(Pink("\n========== Third Request ==============="))

	buf = []byte{
		0x80, 0x85, 0x8c, 0x8b,
		0x84, 0x00, 0x88, 0x4e,
		0xb0, 0x8b, 0x74, 0x97,
		0x90, 0xfa, 0x7f, 0x89,
		0x4e, 0xb0, 0x8b, 0x74,
		0x97, 0x9a, 0x17, 0xa8,
		0xff,
	}

	expectedES = &EmittedSet{
		HeaderField{":method", "GET"},
		HeaderField{":scheme", "https"},
		HeaderField{":path", "/index.html"},
		HeaderField{":authority", "www.example.com"},
		HeaderField{"custom-key", "custom-value"},
	}

	expectedHT = NewHeaderTable(DEFAULT_HEADER_TABLE_SIZE)
	expectedHT.HeaderFields = []*HeaderField{
		NewHeaderField("custom-key", "custom-value"),
		NewHeaderField(":path", "/index.html"),
		NewHeaderField(":scheme", "https"),
		NewHeaderField("cache-control", "no-cache"),
		NewHeaderField(":authority", "www.example.com"),
		NewHeaderField(":path", "/"),
		NewHeaderField(":scheme", "http"),
		NewHeaderField(":method", "GET"),
	}

	expectedRS = ReferenceSet{
		{NewHeaderField("custom-key", "custom-value"), true},
		{NewHeaderField(":authority", "www.example.com"), true},
		{NewHeaderField(":path", "/index.html"), true},
		{NewHeaderField(":scheme", "https"), true},
		{NewHeaderField(":method", "GET"), true},
	}

	context.Decode(buf)

	// test Header Table
	if context.HT.Size() != expectedHT.Size() {
		t.Errorf("\n got %v\nwant %v", context.HT.Size(), expectedHT.Size())
	}

	// test Header Table
	for i, hf := range expectedHT.HeaderFields {
		if !(*context.HT.HeaderFields[i] == *hf) {
			t.Errorf("\n got %v\nwant %v", *context.HT.HeaderFields[i], *hf)
		}
	}

	// test Emitted Set
	sort.Sort(context.ES)
	sort.Sort(expectedES)
	if !reflect.DeepEqual(context.ES, expectedES) {
		t.Errorf("\n got %v\nwant %v", context.ES, expectedES)
	}

	// test Reference Set
	if expectedRS.Len() != context.RS.Len() {
		t.Errorf("\n got %v\nwant %v", context.RS.Dump(), expectedRS.Dump())
	}

	for i, hf := range *context.RS {
		i := len(expectedRS) - i - 1
		if *expectedRS[i].HeaderField != *hf.HeaderField {
			t.Errorf("\n got %v\nwant %v", context.RS.Dump(), expectedRS.Dump())
		}
	}
}

func TestResponseWithoutHuffman(t *testing.T) {
	var (
		context    Context
		buf        []byte
		expectedES *EmittedSet
		expectedHT *HeaderTable
		expectedRS ReferenceSet
	)

	var HeaderTableSize int = 256
	context = NewContext(RESPONSE, HeaderTableSize)

	/**
	 * First Response
	 */
	Debug(Pink("\n========== First Response ==============="))

	buf = []byte{
		0x08, 0x82, 0x40, 0x9f,
		0x18, 0x86, 0xc3, 0x1b,
		0x39, 0xbf, 0x38, 0x7f,
		0x22, 0x92, 0xa2, 0xfb,
		0xa2, 0x03, 0x20, 0xf2,
		0xab, 0x30, 0x31, 0x24,
		0x01, 0x8b, 0x49, 0x0d,
		0x32, 0x09, 0xe8, 0x77,
		0x30, 0x93, 0xe3, 0x9e,
		0x78, 0x64, 0xdd, 0x7a,
		0xfd, 0x3d, 0x3d, 0x24,
		0x87, 0x47, 0xdb, 0x87,
		0x28, 0x49, 0x55, 0xf6,
		0xff,
	}

	expectedES = &EmittedSet{
		HeaderField{":status", "302"},
		HeaderField{"cache-control", "private"},
		HeaderField{"date", "Mon, 21 Oct 2013 20:13:21 GMT"},
		HeaderField{"location", "https://www.example.com"},
	}

	expectedHT = NewHeaderTable(DEFAULT_HEADER_TABLE_SIZE)
	expectedHT.HeaderFields = []*HeaderField{
		NewHeaderField("location", "https://www.example.com"),
		NewHeaderField("date", "Mon, 21 Oct 2013 20:13:21 GMT"),
		NewHeaderField("cache-control", "private"),
		NewHeaderField(":status", "302"),
	}

	expectedRS = ReferenceSet{
		{NewHeaderField("location", "https://www.example.com"), false},
		{NewHeaderField("date", "Mon, 21 Oct 2013 20:13:21 GMT"), false},
		{NewHeaderField("cache-control", "private"), false},
		{NewHeaderField(":status", "302"), false},
	}

	context.Decode(buf)

	// test Header Table
	if context.HT.Size() != expectedHT.Size() {
		t.Errorf("\n got %v\nwant %v", context.HT.Size(), expectedHT.Size())
	}

	// test Header Table
	for i, hf := range expectedHT.HeaderFields {
		if !(*context.HT.HeaderFields[i] == *hf) {
			t.Errorf("\n got %v\nwant %v", *context.HT.HeaderFields[i], *hf)
		}
	}

	// test Emitted Set
	sort.Sort(context.ES)
	sort.Sort(expectedES)
	if !reflect.DeepEqual(context.ES, expectedES) {
		t.Errorf("\n got %v\nwant %v", context.ES, expectedES)
	}

	// test Reference Set
	if expectedRS.Len() != context.RS.Len() {
		t.Errorf("\n got %v\nwant %v", context.RS.Dump(), expectedRS.Dump())
	}

	for i, hf := range *context.RS {
		i := len(expectedRS) - i - 1
		if *expectedRS[i].HeaderField != *hf.HeaderField {
			t.Errorf("\n got %v\nwant %v", context.RS.Dump(), expectedRS.Dump())
		}
	}

	/**
	 * Second Response
	 */
	Debug(Pink("\n========== Second Response ==============="))

	buf = []byte{
		0x84, 0x8c,
	}

	expectedES = &EmittedSet{
		HeaderField{":status", "200"},
		HeaderField{"cache-control", "private"},
		HeaderField{"date", "Mon, 21 Oct 2013 20:13:21 GMT"},
		HeaderField{"location", "https://www.example.com"},
	}

	expectedHT = NewHeaderTable(DEFAULT_HEADER_TABLE_SIZE)
	expectedHT.HeaderFields = []*HeaderField{
		NewHeaderField(":status", "200"),
		NewHeaderField("location", "https://www.example.com"),
		NewHeaderField("date", "Mon, 21 Oct 2013 20:13:21 GMT"),
		NewHeaderField("cache-control", "private"),
	}

	expectedRS = ReferenceSet{
		{NewHeaderField(":status", "200"), false},
		{NewHeaderField("location", "https://www.example.com"), false},
		{NewHeaderField("date", "Mon, 21 Oct 2013 20:13:21 GMT"), false},
		{NewHeaderField("cache-control", "private"), false},
	}

	context.Decode(buf)

	// test Header Table
	if context.HT.Size() != expectedHT.Size() {
		t.Errorf("\n got %v\nwant %v", context.HT.Size(), expectedHT.Size())
	}

	// test Header Table
	for i, hf := range expectedHT.HeaderFields {
		if !(*context.HT.HeaderFields[i] == *hf) {
			t.Errorf("\n got %v\nwant %v", *context.HT.HeaderFields[i], *hf)
		}
	}

	// test Emitted Set
	sort.Sort(context.ES)
	sort.Sort(expectedES)
	if !reflect.DeepEqual(context.ES, expectedES) {
		t.Errorf("\n got %v\nwant %v", context.ES, expectedES)
	}

	// test Reference Set
	if expectedRS.Len() != context.RS.Len() {
		t.Errorf("\n got %v\nwant %v", context.RS.Dump(), expectedRS.Dump())
	}

	for i, hf := range *context.RS {
		i := len(expectedRS) - i - 1
		if *expectedRS[i].HeaderField != *hf.HeaderField {
			t.Errorf("\n got %v\nwant %v", context.RS.Dump(), expectedRS.Dump())
		}
	}

	/**
	 * Third Response
	 */
	Debug(Pink("\n========== Third Response ==============="))

	buf = []byte{
		0x83, 0x84, 0x84, 0x03,
		0x1d, 0x4d, 0x6f, 0x6e,
		0x2c, 0x20, 0x32, 0x31,
		0x20, 0x4f, 0x63, 0x74,
		0x20, 0x32, 0x30, 0x31,
		0x33, 0x20, 0x32, 0x30,
		0x3a, 0x31, 0x33, 0x3a,
		0x32, 0x32, 0x20, 0x47,
		0x4d, 0x54, 0x1d, 0x04,
		0x67, 0x7a, 0x69, 0x70,
		0x84, 0x84, 0x83, 0x83,
		0x3a, 0x38, 0x66, 0x6f,
		0x6f, 0x3d, 0x41, 0x53,
		0x44, 0x4a, 0x4b, 0x48,
		0x51, 0x4b, 0x42, 0x5a,
		0x58, 0x4f, 0x51, 0x57,
		0x45, 0x4f, 0x50, 0x49,
		0x55, 0x41, 0x58, 0x51,
		0x57, 0x45, 0x4f, 0x49,
		0x55, 0x3b, 0x20, 0x6d,
		0x61, 0x78, 0x2d, 0x61,
		0x67, 0x65, 0x3d, 0x33,
		0x36, 0x30, 0x30, 0x3b,
		0x20, 0x76, 0x65, 0x72,
		0x73, 0x69, 0x6f, 0x6e,
		0x3d, 0x31,
	}

	expectedES = &EmittedSet{
		HeaderField{":status", "200"},
		HeaderField{"cache-control", "private"},
		HeaderField{"date", "Mon, 21 Oct 2013 20:13:22 GMT"},
		HeaderField{"location", "https://www.example.com"},
		HeaderField{"content-encoding", "gzip"},
		HeaderField{"set-cookie", "foo=ASDJKHQKBZXOQWEOPIUAXQWEOIU; max-age=3600; version=1"},
	}

	expectedHT = NewHeaderTable(DEFAULT_HEADER_TABLE_SIZE)
	expectedHT.HeaderFields = []*HeaderField{
		NewHeaderField("set-cookie", "foo=ASDJKHQKBZXOQWEOPIUAXQWEOIU; max-age=3600; version=1"),
		NewHeaderField("content-encoding", "gzip"),
		NewHeaderField("date", "Mon, 21 Oct 2013 20:13:22 GMT"),
	}

	expectedRS = ReferenceSet{
		{NewHeaderField("set-cookie", "foo=ASDJKHQKBZXOQWEOPIUAXQWEOIU; max-age=3600; version=1"), false},
		{NewHeaderField("content-encoding", "gzip"), false},
		{NewHeaderField("date", "Mon, 21 Oct 2013 20:13:22 GMT"), false},
	}

	context.Decode(buf)

	// test Header Table
	if context.HT.Size() != expectedHT.Size() {
		t.Errorf("\n got %v\nwant %v", context.HT.Size(), expectedHT.Size())
	}

	// test Header Table
	for i, hf := range expectedHT.HeaderFields {
		if !(*context.HT.HeaderFields[i] == *hf) {
			t.Errorf("\n got %v\nwant %v", *context.HT.HeaderFields[i], *hf)
		}
	}

	// test Emitted Set
	sort.Sort(context.ES)
	sort.Sort(expectedES)
	if !reflect.DeepEqual(context.ES, expectedES) {
		t.Errorf("\n got %v\nwant %v", context.ES, expectedES)
	}

	// test Reference Set
	if expectedRS.Len() != context.RS.Len() {
		t.Errorf("\n got %v\nwant %v", context.RS.Dump(), expectedRS.Dump())
	}

	for i, hf := range *context.RS {
		i := len(expectedRS) - i - 1
		if *expectedRS[i].HeaderField != *hf.HeaderField {
			t.Errorf("\n got %v\nwant %v", context.RS.Dump(), expectedRS.Dump())
		}
	}
}

func TestResponseWithHuffman(t *testing.T) {
	var (
		context    Context
		buf        []byte
		expectedES *EmittedSet
		expectedHT *HeaderTable
		expectedRS ReferenceSet
	)

	var HeaderTableSize int = 256
	context = NewContext(RESPONSE, HeaderTableSize)

	/**
	 * First Response
	 */
	Debug(Pink("\n========== First Response ==============="))

	buf = []byte{
		0x08, 0x82, 0x40, 0x9f,
		0x18, 0x86, 0xc3, 0x1b,
		0x39, 0xbf, 0x38, 0x7f,
		0x22, 0x92, 0xa2, 0xfb,
		0xa2, 0x03, 0x20, 0xf2,
		0xab, 0x30, 0x31, 0x24,
		0x01, 0x8b, 0x49, 0x0d,
		0x32, 0x09, 0xe8, 0x77,
		0x30, 0x93, 0xe3, 0x9e,
		0x78, 0x64, 0xdd, 0x7a,
		0xfd, 0x3d, 0x3d, 0x24,
		0x87, 0x47, 0xdb, 0x87,
		0x28, 0x49, 0x55, 0xf6,
		0xff,
	}

	expectedES = &EmittedSet{
		HeaderField{":status", "302"},
		HeaderField{"cache-control", "private"},
		HeaderField{"date", "Mon, 21 Oct 2013 20:13:21 GMT"},
		HeaderField{"location", "https://www.example.com"},
	}

	expectedHT = NewHeaderTable(DEFAULT_HEADER_TABLE_SIZE)
	expectedHT.HeaderFields = []*HeaderField{
		NewHeaderField("location", "https://www.example.com"),
		NewHeaderField("date", "Mon, 21 Oct 2013 20:13:21 GMT"),
		NewHeaderField("cache-control", "private"),
		NewHeaderField(":status", "302"),
	}

	expectedRS = ReferenceSet{
		{NewHeaderField("location", "https://www.example.com"), false},
		{NewHeaderField("date", "Mon, 21 Oct 2013 20:13:21 GMT"), false},
		{NewHeaderField("cache-control", "private"), false},
		{NewHeaderField(":status", "302"), false},
	}

	context.Decode(buf)

	// test Header Table
	if context.HT.Size() != expectedHT.Size() {
		t.Errorf("\n got %v\nwant %v", context.HT.Size(), expectedHT.Size())
	}

	// test Header Table
	for i, hf := range expectedHT.HeaderFields {
		if !(*context.HT.HeaderFields[i] == *hf) {
			t.Errorf("\n got %v\nwant %v", *context.HT.HeaderFields[i], *hf)
		}
	}

	// test Emitted Set
	sort.Sort(context.ES)
	sort.Sort(expectedES)
	if !reflect.DeepEqual(context.ES, expectedES) {
		t.Errorf("\n got %v\nwant %v", context.ES, expectedES)
	}

	// test Reference Set
	if expectedRS.Len() != context.RS.Len() {
		t.Errorf("\n got %v\nwant %v", context.RS.Dump(), expectedRS.Dump())
	}

	for i, hf := range *context.RS {
		i := len(expectedRS) - i - 1
		if *expectedRS[i].HeaderField != *hf.HeaderField {
			t.Errorf("\n got %v\nwant %v", context.RS.Dump(), expectedRS.Dump())
		}
	}

	/**
	 * Second Response
	 */
	Debug(Pink("\n========== Second Response ==============="))

	buf = []byte{
		0x84, 0x8c,
	}

	expectedES = &EmittedSet{
		HeaderField{":status", "200"},
		HeaderField{"cache-control", "private"},
		HeaderField{"date", "Mon, 21 Oct 2013 20:13:21 GMT"},
		HeaderField{"location", "https://www.example.com"},
	}

	expectedHT = NewHeaderTable(DEFAULT_HEADER_TABLE_SIZE)
	expectedHT.HeaderFields = []*HeaderField{
		NewHeaderField(":status", "200"),
		NewHeaderField("location", "https://www.example.com"),
		NewHeaderField("date", "Mon, 21 Oct 2013 20:13:21 GMT"),
		NewHeaderField("cache-control", "private"),
	}

	expectedRS = ReferenceSet{
		{NewHeaderField(":status", "200"), false},
		{NewHeaderField("location", "https://www.example.com"), false},
		{NewHeaderField("date", "Mon, 21 Oct 2013 20:13:21 GMT"), false},
		{NewHeaderField("cache-control", "private"), false},
	}

	context.Decode(buf)

	// test Header Table
	if context.HT.Size() != expectedHT.Size() {
		t.Errorf("\n got %v\nwant %v", context.HT.Size(), expectedHT.Size())
	}

	// test Header Table
	for i, hf := range expectedHT.HeaderFields {
		if !(*context.HT.HeaderFields[i] == *hf) {
			t.Errorf("\n got %v\nwant %v", *context.HT.HeaderFields[i], *hf)
		}
	}

	// test Emitted Set
	sort.Sort(context.ES)
	sort.Sort(expectedES)
	if !reflect.DeepEqual(context.ES, expectedES) {
		t.Errorf("\n got %v\nwant %v", context.ES, expectedES)
	}

	// test Reference Set
	if expectedRS.Len() != context.RS.Len() {
		t.Errorf("\n got %v\nwant %v", context.RS.Dump(), expectedRS.Dump())
	}

	for i, hf := range *context.RS {
		i := len(expectedRS) - i - 1
		if *expectedRS[i].HeaderField != *hf.HeaderField {
			t.Errorf("\n got %v\nwant %v", context.RS.Dump(), expectedRS.Dump())
		}
	}

	/**
	 * Third Response
	 */
	Debug(Pink("\n========== Third Response ==============="))

	buf = []byte{
		0x83, 0x84, 0x84, 0x03,
		0x92, 0xa2, 0xfb, 0xa2,
		0x03, 0x20, 0xf2, 0xab,
		0x30, 0x31, 0x24, 0x01,
		0x8b, 0x49, 0x0d, 0x33,
		0x09, 0xe8, 0x77, 0x1d,
		0x84, 0xe1, 0xfb, 0xb3,
		0x0f, 0x84, 0x84, 0x83,
		0x83, 0x3a, 0xb3, 0xdf,
		0x7d, 0xfb, 0x36, 0xd3,
		0xd9, 0xe1, 0xfc, 0xfc,
		0x3f, 0xaf, 0xe7, 0xab,
		0xfc, 0xfe, 0xfc, 0xbf,
		0xaf, 0x3e, 0xdf, 0x2f,
		0x97, 0x7f, 0xd3, 0x6f,
		0xf7, 0xfd, 0x79, 0xf6,
		0xf9, 0x77, 0xfd, 0x3d,
		0xe1, 0x6b, 0xfa, 0x46,
		0xfe, 0x10, 0xd8, 0x89,
		0x44, 0x7d, 0xe1, 0xce,
		0x18, 0xe5, 0x65, 0xf7,
		0x6c, 0x2f,
	}

	expectedES = &EmittedSet{
		HeaderField{":status", "200"},
		HeaderField{"cache-control", "private"},
		HeaderField{"date", "Mon, 21 Oct 2013 20:13:22 GMT"},
		HeaderField{"location", "https://www.example.com"},
		HeaderField{"content-encoding", "gzip"},
		HeaderField{"set-cookie", "foo=ASDJKHQKBZXOQWEOPIUAXQWEOIU; max-age=3600; version=1"},
	}

	expectedHT = NewHeaderTable(DEFAULT_HEADER_TABLE_SIZE)
	expectedHT.HeaderFields = []*HeaderField{
		NewHeaderField("set-cookie", "foo=ASDJKHQKBZXOQWEOPIUAXQWEOIU; max-age=3600; version=1"),
		NewHeaderField("content-encoding", "gzip"),
		NewHeaderField("date", "Mon, 21 Oct 2013 20:13:22 GMT"),
	}

	expectedRS = ReferenceSet{
		{NewHeaderField("set-cookie", "foo=ASDJKHQKBZXOQWEOPIUAXQWEOIU; max-age=3600; version=1"), false},
		{NewHeaderField("content-encoding", "gzip"), false},
		{NewHeaderField("date", "Mon, 21 Oct 2013 20:13:22 GMT"), false},
	}

	context.Decode(buf)

	// test Header Table
	if context.HT.Size() != expectedHT.Size() {
		t.Errorf("\n got %v\nwant %v", context.HT.Size(), expectedHT.Size())
	}

	// test Header Table
	for i, hf := range expectedHT.HeaderFields {
		if !(*context.HT.HeaderFields[i] == *hf) {
			t.Errorf("\n got %v\nwant %v", *context.HT.HeaderFields[i], *hf)
		}
	}

	// test Emitted Set
	sort.Sort(context.ES)
	sort.Sort(expectedES)
	if !reflect.DeepEqual(context.ES, expectedES) {
		t.Errorf("\n got %v\nwant %v", context.ES, expectedES)
	}

	// test Reference Set
	if expectedRS.Len() != context.RS.Len() {
		t.Errorf("\n got %v\nwant %v", context.RS.Dump(), expectedRS.Dump())
	}

	for i, hf := range *context.RS {
		i := len(expectedRS) - i - 1
		if *expectedRS[i].HeaderField != *hf.HeaderField {
			t.Errorf("\n got %v\nwant %v", context.RS.Dump(), expectedRS.Dump())
		}
	}
}

func TestEncode(t *testing.T) {
	context := NewContext(REQUEST, DEFAULT_HEADER_TABLE_SIZE)

	hs := HeaderSet{
		NewHeaderField(":status", "200"),
		NewHeaderField("cache-control", "private"),
		NewHeaderField("date", "Mon, 21 Oct 2013 20:13:22 GMT"),
		NewHeaderField("location", "https://www.example.com"),
		NewHeaderField("content-encoding", "gzip"),
		NewHeaderField("set-cookie", "foo=ASDJKHQKBZXOQWEOPIUAXQWEOIU; max-age=3600; version=1"),
	}

	encoded := context.Encode(hs)

	context.Decode(encoded)
	for i, v := range *context.ES {
		if *hs[i] != v {
			t.Errorf("\n got %v\nwant %v", context.ES.Dump(), hs)
		}
	}
}
