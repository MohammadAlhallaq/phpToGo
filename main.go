package main

import (
	"fmt"
	"sync"
)

func addProgrammer(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("another routine")
}

func main() {
	wg := sync.WaitGroup{}

	wg.Add(1)

	go addProgrammer(&wg)

	wg.Wait()
}
