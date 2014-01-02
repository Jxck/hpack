package hpack

// A header set is a potentially ordered group of header fields that are encoded jointly.
// A complete set of key-value pairs contained in a HTTP request or response is a header set.
type HeaderSet []*HeaderField
