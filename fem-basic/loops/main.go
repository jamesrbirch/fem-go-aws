package main

import (
	"fmt"
)

func main() {
	// Fixed size array
	animals := [2]string{}
	animals[0] = "dog"
	animals[1] = "cat"

	fmt.Println(animals)

	animals2 := [2]string{
		"dog",
		"cat",
	}
	fmt.Println(animals2)

	// Slice
	// Acts more like a dynamic array
	animalsSlice := []string{
		"dog",
		"cat",
	}

	animalsSlice = append(animalsSlice, "moose")

	// this uses the range of indexes we want to delete
	// animalsSlice = slices.Delete(animalsSlice, 0, 1)

	fmt.Println(animalsSlice)

	for i := 0; i < len(animalsSlice); i++ {
		fmt.Printf("this is my animal %s\n", animalsSlice[i])
	}

	// iterate a slice using ranges
	for index, value := range animalsSlice {
		fmt.Printf("this is my index %d and this is my animal %s\n", index, value)
	}

	// while
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

}
