package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func main() {

	var dilbert Employee
	position := &dilbert.Position
	*position = "Senior " + *position // promoted, for outsourcing to Elbonia
	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"

	fmt.Println(employeeOfTheMonth)
}
