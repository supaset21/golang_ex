package main

import (
	"hello/employee"
	"hello/ice"
)

func main() {
	e := employee.Employee{
		FirstName:   "Ice",
		LastName:    "Icce",
		TotalLeaves: 30,
		LeavesTaken: 20,
	}

	e.LeavesRemaining()

	i := ice.Ice{
		FirstName:   "Ice",
		LastName:    "Icce",
		TotalLeaves: 30,
		LeavesTaken: 20,
	}

	i.LeavesRemaining()
}
