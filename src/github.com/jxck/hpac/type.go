package hpac

type Frame interface {
}

// Indexed Header Representation
type IndexedHeader struct {
	Flag1 uint8
	Index uint8
}

// Literal Header without Indexing - New Name
type NewNameWithoutIndexing struct {
	Flag1       uint8
	Flag2       uint8
	Flag3       uint8
	Index       uint32
	NameLength  uint32
	NameString  string
	ValueLength uint32
	ValueString string
}

// Literal Header without Indexing - Indexed Name
type IndexedNameWithoutIndexing struct {
	Flag1       uint8
	Flag2       uint8
	Flag3       uint8
	Index       uint32
	ValueLength uint32
	ValueString string
}

// Literal Header with Incremental Indexing - Indexed Name
type IndexedNameWithIncrementalIndexing struct {
	Flag1       uint8
	Flag2       uint8
	Flag3       uint8
	Index       uint32
	ValueLength uint32
	ValueString string
}

// Literal Header with Incremental Indexing - New Name
type NewNameWithIncrementalIndexing struct {
	Flag1       uint8
	Flag2       uint8
	Flag3       uint8
	Index       uint8
	NameLength  uint32
	NameString  string
	ValueLength uint32
	ValueString string
}

// Literal Header with Substitution Indexing - Indexed Name
type IndexedNameWithSubstitutionIndexing struct {
	Flag1            uint8
	Flag2            uint8
	Index            uint32
	SubstitutedIndex uint32
	ValueLength      uint32
	ValueString      string
}

// Literal Header with Substitution Indexing - New Name
type NewNameWithSubstitutionIndexing struct {
	Flag1            uint8
	Flag2            uint8
	Index            uint8
	NameLength       uint32
	NameString       string
	SubstitutedIndex uint32
	ValueLength      uint32
	ValueString      string
}
