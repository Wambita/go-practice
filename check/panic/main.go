// Write your answer here, and then test your code.
// Your job is to implement the PopulateData() function.

package main

import (
	"fmt"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const (
	showExpectedResult = false
	showHints          = false
)

const (
	SuccessMsg string = "Operation successful!"
	PanicMsg   string = "Operation panicked!"
)

// Person is a name and age tuple.
type Person struct {
	Name string
	Age  int
}

// Persons saves a map of Person.
type Persons struct {
	m map[string]Person
}

// NewPersons initializes the Persons type.
func NewPersons() *Persons {
	return &Persons{
		m: make(map[string]Person),
	}
}

// PanickingAdd adds a new Person to the Persons receiver and
// panics in the case of invalid ages.
func (p *Persons) PanickingAdd(np Person) {
	if np.Age < 0 || np.Age > 100 {
		panic(fmt.Sprintf("add:invalid age %d for name %s", np.Age, np.Name))
	}
	p.m[np.Name] = np
}

// PopulateData should add data to the Persons type without panicking.
func PopulateData(data []Person) (result string) {
	result = SuccessMsg
	// Your code here
	defer func() {
		if r := recover(); r != nil {
			result = PanicMsg
		}
	}()
	p := NewPersons()
	for _, np := range data {
		p.PanickingAdd(np)
	}
	return
}
