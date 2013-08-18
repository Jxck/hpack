package main

import (
	. "github.com/jxck/hpac"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)
	RequestHeaderTable.Add(Header{"hoge", "fuga"})

	log.Println(RequestHeaderTable)

	ref := ReferenceSet{}
	ref.Add("hoge", "fuga")
	log.Println(ref)
	ref.Set("hoge", "piyo")
	log.Println(ref)
	ref.Del("hoge")
	log.Println(ref)
}
