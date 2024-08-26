package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var global int


//Overview of New Atomic Operations Provided Since Go 1.19
//Go 1.19 introduced several types, each of which owns a set of atomic operation methods, to achieve the same effects made by the package-level functions listed in the last section.
//Among these types, Int32, Int64, Uint32, Uint64 and Uintptr are for integer atomic operations. The methods of the atomic.Int32 type are listed below. The methods of the other four types present in the similar way. 


func main() {

	var wg sync.WaitGroup

	var ops atomic.Int32

	for i := 0; i < 50; i++ {

		wg.Add(1)

		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				global = global + 1
				ops.Add(1)
			}
		}()

	}

	wg.Wait()

	fmt.Println("global :", global)
	fmt.Println("ops :", ops.Load())
}
