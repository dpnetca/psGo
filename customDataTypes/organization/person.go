package organization

import (
	"errors"
	"fmt"
	"strings"
)

type Identifiable interface {
	ID() string
}

type Citizen interface {
	Identifiable
	Country() string
}

/* old
type PersonV1 struct {
	FirstName string
	LastName  string
}

func (p PersonV1) ID() string {
	return "V1-1234"
}
*/
// this is a type alias
// type TwitterHandler = string
/*
alias can be used interchangably with base type
but can build functions to extend an alias

alias copys the fields and the methods
definition only copy fields but allows you to define own methods
*/

// this is a type declaration (or type definition?)
type TwitterHandler string

func (th TwitterHandler) RedirectUrl() string {
	cleanHandler := strings.TrimPrefix(string(th), "@")
	return fmt.Sprintf("https://www.twitter.com/%s", cleanHandler)
}

type socialSecurityNumber string

func NewSocialSecurityNumber(value string) Citizen {
	return socialSecurityNumber(value)
}

func (ssn socialSecurityNumber) ID() string {
	return string(ssn)
}

func (ssn socialSecurityNumber) Country() string {
	return "USA"
}

type eurpoeanUnionIdentifier struct {
	id      string
	country string
}

func NewEurpoeanUnionIdentifier(id, country string) Citizen {
	return eurpoeanUnionIdentifier{
		id:      id,
		country: country,
	}
}

func (eui eurpoeanUnionIdentifier) ID() string {
	return eui.id
}

func (eui eurpoeanUnionIdentifier) Country() string {
	return fmt.Sprintf("EU: %s", eui.country)
}

type Name struct {
	first string
	last  string
}

func (n *Name) FullName() string {
	return fmt.Sprintf("%s %s", n.first, n.last)
}

type Employee struct {
	Name
}

// lower case fields are private and not directly accessible
type PersonV2 struct {
	// firstName      string
	// lastName       string
	// name           Name // this would be accessed as p.name.FullName
	Name           // this embeds the methods/fields directly so I can do p.FullName still
	twitterHandler TwitterHandler
	Citizen
}

// use a function to create the object
func NewPerson(firstName, lastName string, citizen Citizen) PersonV2 {
	return PersonV2{
		Name:    Name{firstName, lastName},
		Citizen: citizen,
		// firstName: firstName,
		// lastName:  lastName,

	}
}

// use setters and getters to access fields
// func (p *PersonV2) FullName() string {
// 	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
// }

// this would overwrite the emded type "Idenfiable"
func (p *PersonV2) ID() string {
	// return "V2-12345"
	return fmt.Sprintf("Persons Identifier: %s", p.Citizen.ID())
}

// needs to be pointer based to set the data
// func (p *PersonV2) SetTwitterHandler(handler string) error {
func (p *PersonV2) SetTwitterHandler(handler TwitterHandler) error {
	if len(handler) == 0 {
		p.twitterHandler = handler
		// } else if !strings.HasPrefix(handler, "@") {
	} else if !strings.HasPrefix(string(handler), "@") {
		return errors.New("twitter handler must start with an @ symbol")
	}

	p.twitterHandler = handler
	return nil
}

// func (p *PersonV2) TwitterHandler() string {
func (p *PersonV2) TwitterHandler() TwitterHandler {
	return p.twitterHandler
}
