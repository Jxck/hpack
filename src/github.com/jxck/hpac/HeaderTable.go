package hpac

type Header struct {
	HeaderName, HeaderValue string
}

type HeaderTable []Header

func (ht *HeaderTable) Add(header Header) {
	*ht = append(*ht, header)
}

var RequestHeaderTable = HeaderTable{
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

var ResponseHeaderTable = []Header{
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
