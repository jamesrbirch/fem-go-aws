package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) Person {
	return Person{
		Name: name,
		Age:  age,
	}
}

// receiver - a method implemented on a type
// * - gives the reference to the allocated memory
// we can get the allocated memory address by using fmt.Println("address of item", &p.name)
func (p *Person) changeName(newName string) {
	p.Name = newName
}

func main() {
	// structs are go's answer to classes and object oriented design
	// favourite composition over inheritance
	myPerson := NewPerson("James", 42)
	myPerson.changeName("lp")

	// declare a memory pointer
	a := 7
	b := &a

	// dereference the memory allocation and store 9
	*b = 9
	fmt.Println(b)

	fmt.Printf("This is my person %+v\n", myPerson)

	mySlice := []int{
		1, 2, 3,
	}

	for index, _ := range mySlice {
		mySlice[index]++
	}

	fmt.Println(mySlice)

}
