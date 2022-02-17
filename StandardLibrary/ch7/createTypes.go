package main

import (
	"fmt"
	"reflect"
)

type employee struct {
	employeeId int
	firstName  string
	lastName   string
}

func main() {
	employees := make([]employee, 3)

	employees = append(employees, employee{0, "Lloyd", "Christmas"})
	employees = append(employees, employee{1, "Harry", "Dune"})
	employees = append(employees, employee{0, "Sea", "Bass"})
	// fmt.Printf("the type is %v \n", reflect.TypeOf(employees))
	// fmt.Printf("the value is %v \n", reflect.ValueOf(employees))
	// fmt.Printf("the kind of type is %v \n", reflect.ValueOf(employees).Kind())
	// fmt.Printf("the Name of type is %v \n", reflect.TypeOf(employees).Name()) // this returns blank...slice doesn't have a name
	// fmt.Printf("the Name of type is %v \n", reflect.TypeOf(employees[3]).Name())

	eType := reflect.TypeOf(employees)
	newEmployeeList := reflect.MakeSlice(eType, 0, 0)
	newEmployeeList = reflect.Append(newEmployeeList, reflect.ValueOf(employee{3, "Mary", "Swanson"}))
	fmt.Printf("First List %v\n", employees)
	fmt.Printf("new List %v\n", newEmployeeList)
}
