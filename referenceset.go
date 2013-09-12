package hpack

type ReferenceSet map[string]string

func NewReferenceSet() ReferenceSet {
	return ReferenceSet{}
}

func (r ReferenceSet) Add(key, value string) {
	r[key] = value
}

func (r ReferenceSet) Del(key string) {
	delete(r, key)
}
