package main_test

import (
	"log"
	"runtime"
	"testing"
	"unique"
)

var ss = []string{
	"fooooooooooooooooooooo",
	"baaaaaaaaaaaaaaaaaaaar",
	"baaaaaaaaaaaaaaaaaaaaz",
}

func getAlloc() uint64 {
	var m runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m)
	return m.Alloc
}

func BenchmarkWithoutUnique(b *testing.B) {
	b.ResetTimer()

	log.Println(getAlloc())
	a := make([]string, b.N)
	for i := 0; i < len(a); i++ {
		a[i] = ss[i%len(ss)]
	}
	log.Println(getAlloc())
}

func BenchmarkWithUnique(b *testing.B) {
	b.ResetTimer()

	log.Println(getAlloc())
	a := make([]unique.Handle[string], b.N)
	for i := 0; i < b.N; i++ {
		a[i] = unique.Make(ss[i%len(ss)])
	}
	log.Println(getAlloc())
}
