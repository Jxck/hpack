package hpack

/*
  HT 1 0 1
     2 1 2
     3 2 3
     4 3 4
  ST 5 0 1
     6 1 2
     7 2 3
     8 3 4

  i = 8
  T[8] = ST[3] = ST[8-4-1]


HT
1 0
2 1
3 2
4 3
5 0
6 1
7 2
8 3
. .
58 53
*/
var StaticHeaderTable = []*HeaderField{
	/* 0*/ {":authority", ""},
	/* 1*/ {":method", "GET"},
	/* 2*/ {":method", "POST"},
	/* 3*/ {":path", "/"},
	/* 4*/ {":path", "/index.html"},
	/* 5*/ {":scheme", "http"},
	/* 6*/ {":scheme", "https"},
	/* 7*/ {":status", "200"},
	/* 8*/ {":status", "500"},
	/* 9*/ {":status", "404"},
	/*10*/ {":status", "403"},
	/*11*/ {":status", "400"},
	/*12*/ {":status", "401"},
	/*13*/ {"accept-charset", ""},
	/*14*/ {"accept-encoding", ""},
	/*15*/ {"accept-language", ""},
	/*16*/ {"accept-ranges", ""},
	/*17*/ {"accept", ""},
	/*18*/ {"access-control-allow-origin", ""},
	/*19*/ {"age", ""},
	/*20*/ {"allow", ""},
	/*21*/ {"authorization", ""},
	/*22*/ {"cache-control", ""},
	/*23*/ {"content-disposition", ""},
	/*24*/ {"content-encoding", ""},
	/*25*/ {"content-language", ""},
	/*26*/ {"content-length", ""},
	/*27*/ {"content-location", ""},
	/*28*/ {"content-range", ""},
	/*29*/ {"content-type", ""},
	/*30*/ {"cookie", ""},
	/*31*/ {"date", ""},
	/*32*/ {"etag", ""},
	/*33*/ {"expect", ""},
	/*34*/ {"expires", ""},
	/*35*/ {"from", ""},
	/*36*/ {"host", ""},
	/*37*/ {"if-match", ""},
	/*38*/ {"if-modified-since", ""},
	/*39*/ {"if-none-match", ""},
	/*40*/ {"if-range", ""},
	/*41*/ {"if-unmodified-since", ""},
	/*42*/ {"last-modified", ""},
	/*43*/ {"link", ""},
	/*44*/ {"location", ""},
	/*45*/ {"max-forwards", ""},
	/*46*/ {"proxy-authenticate", ""},
	/*47*/ {"proxy-authorization", ""},
	/*48*/ {"range", ""},
	/*49*/ {"referer", ""},
	/*50*/ {"refresh", ""},
	/*51*/ {"retry-after", ""},
	/*52*/ {"server", ""},
	/*53*/ {"set-cookie", ""},
	/*54*/ {"strict-transport-security", ""},
	/*55*/ {"transfer-encoding", ""},
	/*56*/ {"user-agent", ""},
	/*57*/ {"vary", ""},
	/*58*/ {"via", ""},
	/*59*/ {"www-authenticate", ""},
}
