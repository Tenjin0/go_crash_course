package employee

import "fmt"

type person struct {
	FirstName string
	LastName  string
}
type employee struct {
	person
	TotalLeaves int
	LeavesTaken int
}

func (e employee) LeavesRemaining() {

	fmt.Printf(
		"%s %s has %d leaves remaining\n",
		e.FirstName,
		e.LastName,
		e.TotalLeaves-e.LeavesTaken,
	)
}

func New(firstName string, lastName string, totalLeaves int, leavesTaken int) employee {
	e := employee{person{firstName, lastName}, totalLeaves, leavesTaken}
	return e
}
