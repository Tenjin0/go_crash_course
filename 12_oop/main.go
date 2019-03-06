package main

import "github.com/Tenjin0/go_crash_course/12_oop/employee"

func main() {

	// e := employee.Employee{
	// 	TotalLeaves: 30,
	// 	LeavesTaken: 20,
	// 	Person: employee.Person{
	// 		FirstName: "Sam",
	// 		LastName:  "Dupont",
	// 	},
	// }
	e := employee.New("Sam", "Dupont", 30, 20)
	e.LeavesRemaining()

}
