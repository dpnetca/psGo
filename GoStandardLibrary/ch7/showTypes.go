package main

import (
	"fmt"
	"reflect"
)

func main() {
	type employee struct {
		employeeId int
		firstName  string
		lastName   string
	}

	type customer struct {
		customerId int
		firstName  string
		lastName   string
		company    string
	}

	newEmployee := employee{0, "Fred", "Flinstone"}
	fmt.Printf("Our new employee is %s %s with an id of %d\n", newEmployee.firstName, newEmployee.lastName, newEmployee.employeeId)

	fmt.Printf("the type is %v \n", reflect.TypeOf(newEmployee))
	fmt.Printf("the value is %v \n", reflect.ValueOf(newEmployee))
	fmt.Printf("the kind of type is %v \n", reflect.ValueOf(newEmployee).Kind())

	fmt.Printf("\n\n*******************\n\n")

	newCustomer := customer{234, "Barney", "Rubble", "Slate Rock and Gravel"}

	addPerson(newEmployee)
	addPerson(newCustomer)
}

func addPerson(p interface{}) bool {
	if reflect.ValueOf(p).Kind() == reflect.Struct {
		v := reflect.ValueOf(p)
		switch reflect.TypeOf(p).Name() {
		case "employee":
			emplSQL := "INSERT INTO employee (employeeID, firstName, lastName VALUES (?,?,?)"
			fmt.Printf("SQL: %s\n", emplSQL)
			fmt.Printf("Added %v\n", v.Field(1))
		case "customer":
			cusSQL := "INSERT INTO customer (customerId, firstName, lastName, company VALUES (?,?,?,?)"
			fmt.Printf("SQL: %s\n", cusSQL)
			fmt.Printf("Added %v\n", v.Field(1))
		default:
			fmt.Println("unhandled type")
		}
		return true
	} else {
		return false

	}
}
