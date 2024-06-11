package imports

import "fmt"

type Ticket struct {
	Id    int
	Event string
}

func (t Ticket) PrintEvent() {
	fmt.Println(t.Event)
}
