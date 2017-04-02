// HelloStructs project main.go
package main

import (
	"fmt"
)

func main() {
	pRoberto := Person{Name: "Roberto", Age: 37}
	fmt.Printf("Person name: %s Age: %d\n", pRoberto.Name, pRoberto.Age)
	MakePersonOld(pRoberto) // won't make it older, cause is passing value
	fmt.Printf("Person name: %s Age: %d\n", pRoberto.Name, pRoberto.Age)
	MakePersonOldRef(&pRoberto) // this time it should work, passing reference to object
	fmt.Printf("Person **older** name: %s Age: %d\n", pRoberto.Name, pRoberto.Age)
}

type Person struct {
	Name string
	Age  int
}

func MakePersonOld(person Person) {
	person.Age += 100
}

func MakePersonOldRef(person *Person) {
	person.Age += 100
}
