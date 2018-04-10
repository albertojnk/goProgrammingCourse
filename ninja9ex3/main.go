package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	inc := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := inc
			runtime.Gosched()
			v++
			inc = v
			fmt.Println(inc)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value: ", inc)
}
