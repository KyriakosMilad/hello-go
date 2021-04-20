package functions

import (
	"fmt"
)

func Functions() {
	// check if error print error if not print the output
	err, a := hello("Kyriakos")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(a)
	}

	// call struct method
	g := greeter{name: "Kamose"}
	g.greet()

	// anonymous function
	func() {
		fmt.Println("hi from anonymous function")
	}()
}

// function name (param 1 param 1 type, param 2 param 2 type) (return 1 type, return 2 type)
// you can also pass pointers to the function not just values itself

func hello(name string) (error, string) {
	if name == "" {
		return fmt.Errorf("error name cannot be empty string"), ""
	}
	return nil, "Hello from function, " + name
}

type greeter struct {
	name string
}

// this is a method for struct greeter

func (g *greeter) greet() {
	if g.name == "" {
		panic("error name cannot be empty string")
	}
	fmt.Println("Hello from method, " + g.name)
}
