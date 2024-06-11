package main

import (
	"femBasics/imports"
	"fmt"
)

func main() {
	newTicket := imports.Ticket{
		Id:    123,
		Event: "FEM course",
	}

	newTicket.PrintEvent()

	fmt.Println(newTicket)
}
