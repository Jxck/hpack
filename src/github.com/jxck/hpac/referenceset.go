package hpac

type ReferenceSet map[string]string

func (r ReferenceSet) Add(key, value string) {
	r[key] = value
}

func (r ReferenceSet) Set(key, value string) {
	r[key] = value
}

func (r ReferenceSet) Del(key string) {
	delete(r, key)
}
