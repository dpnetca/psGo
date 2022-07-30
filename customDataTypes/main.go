package main

import (
	"fmt"

	"github.com/dpnetca/psGo/customerDataTypes/organization"
)

func main() {
	/* old
	p := organization.PersonV1{FirstName: "John", LastName: "Smith"}
	fmt.Println(p.ID())
	fmt.Println(p.FirstName)

	var p2 organization.Identifiable = organization.PersonV1{FirstName: "Jane", LastName: "Doe"}
	fmt.Println(p2.ID())
	// fmt.Println(p2.FirstName)  // This doesn't work....
	*/

	p3 := organization.NewPerson("Jack", "Frost", organization.NewSocialSecurityNumber("123-45-6789"))
	// err := p3.SetTwitterHandler("broken")
	err := p3.SetTwitterHandler("@good")
	if err != nil {
		fmt.Printf("An error occured setting twitter handler: %s\n", err)
	}
	fmt.Println(p3.ID())
	fmt.Println(p3.Country())
	fmt.Println(p3.FullName())
	fmt.Println(p3.TwitterHandler())
	fmt.Println(p3.TwitterHandler().RedirectUrl())

	p4 := organization.NewPerson("Jack", "Frost", organization.NewEurpoeanUnionIdentifier("123", "GER"))
	fmt.Println(p4.ID())
	fmt.Println(p4.Country())

}
