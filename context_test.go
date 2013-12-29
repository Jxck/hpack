package hpack

import (
	"log"
	"net/http"
	"reflect"
	"testing"
)

func TestRequestWithoutHuffman(t *testing.T) {

	client := NewContext()

	/**
	 * First Request
	 */
	log.Println("========== First Request ===============")

	buf := []byte{
		0x82, 0x87, 0x86, 0x04,
		0x0f, 0x77, 0x77, 0x77,
		0x2e, 0x65, 0x78, 0x61,
		0x6d, 0x70, 0x6c, 0x65,
		0x2e, 0x63, 0x6f, 0x6d,
	}

	expectedHeader := http.Header{
		"Method":    []string{"GET"},
		"Scheme":    []string{"http"},
		"Path":      []string{"/"},
		"Authority": []string{"www.example.com"},
	}

	expectedHeaderFields := []HeaderField{
		HeaderField{":authority", "www.example.com"},
		HeaderField{":path", "/"},
		HeaderField{":scheme", "http"},
		HeaderField{":method", "GET"},
	}

	client.Decode(buf)

	// test Header Table
	if client.HT.Size() != 180 {
		t.Errorf("\n got %v\nwant %v", client.HT.Size(), 180)
	}

	// test Header Table
	for i, hf := range expectedHeaderFields {
		if !(*client.HT.HeaderFields[i] == hf) {
			t.Errorf("\n got %v\nwant %v", *client.HT.HeaderFields[i], hf)
		}
	}

	// test Emitted Set
	if !reflect.DeepEqual(client.ES.Header, expectedHeader) {
		log.Println(client.ES.Header)
		t.Errorf("\n got %v\nwant %v", client.ES.Header, expectedHeader)
	}

	// TODO: test Reference Set

	/**
	 * Second Request
	 */
	log.Println("========== Second Request ===============")

	buf = []byte{
		0x1b, 0x08, 0x6e, 0x6f,
		0x2d, 0x63, 0x61, 0x63,
		0x68, 0x65,
	}

	client.Decode(buf)

	expectedHeader = http.Header{
		"Method":        []string{"GET"},
		"Scheme":        []string{"http"},
		"Path":          []string{"/"},
		"Authority":     []string{"www.example.com"},
		"Cache-Control": []string{"no-cache"},
	}

	expectedHeaderFields = []HeaderField{
		HeaderField{"cache-control", "no-cache"},
		HeaderField{":authority", "www.example.com"},
		HeaderField{":path", "/"},
		HeaderField{":scheme", "http"},
		HeaderField{":method", "GET"},
	}

	// test Header Table
	if client.HT.Size() != 233 {
		t.Errorf("\n got %v\nwant %v", client.HT.Size(), 233)
	}

	// test Header Table
	for i, hf := range expectedHeaderFields {
		if !(*client.HT.HeaderFields[i] == hf) {
			t.Errorf("\n got %v\nwant %v", *client.HT.HeaderFields[i], hf)
		}
	}

	// test Emitted Set
	if !reflect.DeepEqual(client.ES.Header, expectedHeader) {
		t.Errorf("\n got %v\nwant %v", client.ES.Header, expectedHeader)
	}

	// TOOD: test Reference Set

	/**
	 * Third Request
	 */
	log.Println("========== Third Request ===============")

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

	client.Decode(buf)

	expectedHeader = http.Header{
		"Method":     []string{"GET"},
		"Scheme":     []string{"https"},
		"Path":       []string{"/index.html"},
		"Authority":  []string{"www.example.com"},
		"Custom-Key": []string{"custom-value"},
	}

	expectedHeaderFields = []HeaderField{
		HeaderField{"custom-key", "custom-value"},
		HeaderField{":path", "/index.html"},
		HeaderField{":scheme", "https"},
		HeaderField{"cache-control", "no-cache"},
		HeaderField{":authority", "www.example.com"},
		HeaderField{":path", "/"},
		HeaderField{":scheme", "http"},
		HeaderField{":method", "GET"},
	}

	// test Header Table
	if client.HT.Size() != 379 {
		t.Errorf("\n got %v\nwant %v", client.HT.Size(), 379)
	}

	// test Header Table
	for i, hf := range expectedHeaderFields {
		if !(*client.HT.HeaderFields[i] == hf) {
			t.Errorf("\n got %v\nwant %v", *client.HT.HeaderFields[i], hf)
		}
	}

	// test Emitted Set
	if !reflect.DeepEqual(client.ES.Header, expectedHeader) {
		t.Errorf("\n got %v\nwant %v", client.ES.Header, expectedHeader)
	}

	// TOOD: test Reference Set
}

func TestRequestWithHuffman(t *testing.T) {

	client := NewContext()

	/**
	 * First Request
	 */
	log.Println("========== First Request ===============")

	buf := []byte{
		0x82, 0x87, 0x86, 0x04,
		0x8b, 0xdb, 0x6d, 0x88,
		0x3e, 0x68, 0xd1, 0xcb,
		0x12, 0x25, 0xba, 0x7f,
	}

	expectedHeader := http.Header{
		"Method":    []string{"GET"},
		"Scheme":    []string{"http"},
		"Path":      []string{"/"},
		"Authority": []string{"www.example.com"},
	}

	expectedHeaderFields := []HeaderField{
		HeaderField{":authority", "www.example.com"},
		HeaderField{":path", "/"},
		HeaderField{":scheme", "http"},
		HeaderField{":method", "GET"},
	}

	client.Decode(buf)

	// test Header Table
	if client.HT.Size() != 180 {
		t.Errorf("\n got %v\nwant %v", client.HT.Size(), 180)
	}

	// test Header Table
	for i, hf := range expectedHeaderFields {
		if !(*client.HT.HeaderFields[i] == hf) {
			t.Errorf("\n got %v\nwant %v", *client.HT.HeaderFields[i], hf)
		}
	}

	// test Emitted Set
	if !reflect.DeepEqual(client.ES.Header, expectedHeader) {
		log.Println(client.ES.Header)
		t.Errorf("\n got %v\nwant %v", client.ES.Header, expectedHeader)
	}

	// TODO: test Reference Set

	/**
	 * Second Request
	 */
	log.Println("========== Second Request ===============")

	buf = []byte{
		0x1b, 0x86, 0x63, 0x65,
		0x4a, 0x13, 0x98, 0xff,
	}

	client.Decode(buf)

	expectedHeader = http.Header{
		"Method":        []string{"GET"},
		"Scheme":        []string{"http"},
		"Path":          []string{"/"},
		"Authority":     []string{"www.example.com"},
		"Cache-Control": []string{"no-cache"},
	}

	expectedHeaderFields = []HeaderField{
		HeaderField{"cache-control", "no-cache"},
		HeaderField{":authority", "www.example.com"},
		HeaderField{":path", "/"},
		HeaderField{":scheme", "http"},
		HeaderField{":method", "GET"},
	}

	// test Header Table
	if client.HT.Size() != 233 {
		t.Errorf("\n got %v\nwant %v", client.HT.Size(), 233)
	}

	// test Header Table
	for i, hf := range expectedHeaderFields {
		if !(*client.HT.HeaderFields[i] == hf) {
			t.Errorf("\n got %v\nwant %v", *client.HT.HeaderFields[i], hf)
		}
	}

	// test Emitted Set
	if !reflect.DeepEqual(client.ES.Header, expectedHeader) {
		t.Errorf("\n got %v\nwant %v", client.ES.Header, expectedHeader)
	}

	// TOOD: test Reference Set

	/**
	 * Third Request
	 */
	log.Println("========== Third Request ===============")

	buf = []byte{
		0x80, 0x85, 0x8c, 0x8b,
		0x84, 0x00, 0x88, 0x4e,
		0xb0, 0x8b, 0x74, 0x97,
		0x90, 0xfa, 0x7f, 0x89,
		0x4e, 0xb0, 0x8b, 0x74,
		0x97, 0x9a, 0x17, 0xa8,
		0xff,
	}

	client.Decode(buf)

	expectedHeader = http.Header{
		"Method":     []string{"GET"},
		"Scheme":     []string{"https"},
		"Path":       []string{"/index.html"},
		"Authority":  []string{"www.example.com"},
		"Custom-Key": []string{"custom-value"},
	}

	expectedHeaderFields = []HeaderField{
		HeaderField{"custom-key", "custom-value"},
		HeaderField{":path", "/index.html"},
		HeaderField{":scheme", "https"},
		HeaderField{"cache-control", "no-cache"},
		HeaderField{":authority", "www.example.com"},
		HeaderField{":path", "/"},
		HeaderField{":scheme", "http"},
		HeaderField{":method", "GET"},
	}

	// test Header Table
	if client.HT.Size() != 379 {
		t.Errorf("\n got %v\nwant %v", client.HT.Size(), 379)
	}

	// test Header Table
	for i, hf := range expectedHeaderFields {
		if !(*client.HT.HeaderFields[i] == hf) {
			t.Errorf("\n got %v\nwant %v", *client.HT.HeaderFields[i], hf)
		}
	}

	// test Emitted Set
	if !reflect.DeepEqual(client.ES.Header, expectedHeader) {
		t.Errorf("\n got %v\nwant %v", client.ES.Header, expectedHeader)
	}

	// TOOD: test Reference Set
}

func TestResponseWithoutHuffman(t *testing.T) {
	t.Skip()

	client := NewContext()

	/**
	 * First Request
	 */
	log.Println("========== First Request ===============")

	buf := []byte{
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

	expectedHeader := http.Header{
		"Status":        []string{"302"},
		"Cache-Control": []string{"private"},
		"Date":          []string{"Mon: []string{ 21 Oct 2013 201321 GMT"},
		"Location":      []string{"https//www.example.com"},
	}

	expectedHeaderFields := []HeaderField{
		HeaderField{"location", "https://www.example.com"},
		HeaderField{"date", "Mon, 21 Oct 2013 20:13:21 GMT"},
		HeaderField{"cache-control", "private"},
		HeaderField{":status", "302"},
	}

	client.Decode(buf)

	// test Header Table
	if client.HT.Size() != 222 {
		t.Errorf("\n got %v\nwant %v", client.HT.Size(), 222)
	}

	// test Header Table
	for i, hf := range expectedHeaderFields {
		if !(*client.HT.HeaderFields[i] == hf) {
			t.Errorf("\n got %v\nwant %v", *client.HT.HeaderFields[i], hf)
		}
	}

	// test Emitted Set
	if !reflect.DeepEqual(client.ES.Header, expectedHeader) {
		log.Println(client.ES.Header)
		t.Errorf("\n got %v\nwant %v", client.ES.Header, expectedHeader)
	}

	// TODO: test Reference Set
}
