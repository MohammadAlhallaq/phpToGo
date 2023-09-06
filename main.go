package main

import "fmt"

func main() {

	slice1 := []string{"a", "b", "c", "d"}
	slice2 := append(slice1[:1], "v", "f", "g", "l")
	fmt.Println(slice1, slice2)

}
