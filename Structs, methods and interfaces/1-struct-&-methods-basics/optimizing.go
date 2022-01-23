package main

import (
	"fmt"
	"unsafe"
)

type memoryUnoptimized struct {
	myBool  bool    // 1 byte
	myFloat float64 // 8 bytes
	myInt   int32   // 4 bytes
}

// Create another struct and re-arrange the fields of memory to optimize the memory allocation

func optimizing() {
	m := memoryUnoptimized{}
	fmt.Println("Bytes allocated for Memory struct:", unsafe.Sizeof(m))
}
