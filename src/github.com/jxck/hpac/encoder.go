package hpac

type Encoder struct {
	requestHeaderTable  HeaderTable
	responseHeaderTable HeaderTable
	referenceSet        ReferenceSet
}

func NewEncoder() Encoder {
	var encoder = Encoder{
		requestHeaderTable:  RequestHeaderTable,
		responseHeaderTable: ResponseHeaderTable,
		referenceSet:        ReferenceSet{},
	}
	return encoder
}
