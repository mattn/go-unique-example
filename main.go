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
}

const N = 10000000

func getAlloc() uint64 {
	var m runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m)
	return m.Alloc
}

func test1() {
	before := getAlloc()
	a := make([]string, N)
	for i := 0; i < len(a); i++ {
		a[i] = ss[i%len(ss)]
	}
	log.Printf("test1: %v allocated", getAlloc()-before)
}

func test2() {
	before := getAlloc()
	a := make([]unique.Handle[string], N)
	for i := 0; i < len(a); i++ {
		a[i] = unique.Make(ss[i%len(ss)])
	}
	log.Printf("test2: %v allocated", getAlloc()-before)
}

func main() {
	test1()
	test2()
}
