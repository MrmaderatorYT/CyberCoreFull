package world

import "fmt"

type InMemory struct {
	data map[string]string
}

func newInMemory() *InMemory {
	return &InMemory{data: make(map[string]string)}
}

func world() {
	// create a new in-memory object
	inMem := newInMemory()
	// add some data
	inMem.data["key1"] = "value1"
	inMem.data["key2"] = "value2"
	// print the data
	fmt.Println(inMem.data)
}
