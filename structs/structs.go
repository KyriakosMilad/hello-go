package structs

import (
	"fmt"
	"reflect"
)

// naming
// just like variables first letter capital (Pascal) means it's exported for out package use
// and first letter small (camelCase) means it's for in package scope
// not just the struct name but also the the fields in the struct should be named the same way

// create struct

type Vehicle struct {
	wheels    int `required:"true" max:"100"` // this called tag look how to use it (line 48)
	motorType string
}

func Structs() {
	// use struct as a type
	var v Vehicle = Vehicle{
		wheels:    4,
		motorType: "IDK",
	}

	// notice: you can create a struct and then add data to it
	//v := Vehicle{}
	//v.wheels = 4
	//v.motorType = "IDK"

	// positional syntax (not recommended)
	//v := Vehicle{4, "IDK"}

	// anonymous struct
	//v := struct{wheels int}{3}

	// get struct data
	fmt.Println(v)

	// get specific field
	fmt.Println(v.wheels)

	// change field value
	v.wheels = 2
	fmt.Println(v.wheels)

	// use field tags (common cases: validate data)
	t := reflect.TypeOf(Vehicle{})
	field, _ := t.FieldByName("wheels")
	fmt.Println(field.Tag)

	// go structs inheritance
	type Car struct {
		Vehicle
		doors int
	}

	c := Car{}
	c.wheels = 4
	c.motorType = "IDK"
	c.doors = 3

	fmt.Println(c)
}
