package main

import (
	"log"
	"runtime"
	"unique"
)

var ss = []string{
	"fooooooooooooooooooooo",
	"baaaaaaaaaaaaaaaaaaaar",
	"baaaaaaaaaaaaaaaaaaaaz",
	"zooooooooooooooooooooo",
	"ohoooooooooooooooooooo",
}

const N = 10000000

func getAlloc() uint64 {
	var m runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m)
	return m.Alloc
}

func test_without_unique_package() {
	before := getAlloc()
	a := make([]string, N)
	for i := 0; i < len(a); i++ {
		a[i] = ss[i%len(ss)]
	}
	log.Printf("test_without_unique_package: %v allocated", getAlloc()-before)
	runtime.KeepAlive(a)
}

func test_with_unique_package() {
	before := getAlloc()
	a := make([]unique.Handle[string], N)
	for i := 0; i < len(a); i++ {
		a[i] = unique.Make(ss[i%len(ss)])
	}
	log.Printf("test_with_unique_package: %v allocated", getAlloc()-before)
	runtime.KeepAlive(a)
}

func main() {
	test_without_unique_package()
	runtime.GC()
	test_with_unique_package()
}
