package hpac

type ReferenceSet map[string]string

func NewReferenceSet() ReferenceSet {
	return ReferenceSet{}
}

func (r ReferenceSet) Add(key, value string) {
	r[key] = value
}

func (r ReferenceSet) Set(key, value string) bool {
	_, ok := r[key]
	if !ok {
		return ok
	}

	r[key] = value
	return true
}

func (r ReferenceSet) Del(key string) {
	delete(r, key)
}
