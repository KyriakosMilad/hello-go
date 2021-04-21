package functions

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func Functions() {
	// check if error print error if not print the output
	err, a := hello("Kyriakos")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(a)
	}

	i := 3
	changeNumber(&i)
	fmt.Println(i)

	// call struct method
	g := greeter{name: "Kamose"}
	g.greet()

	// call struct method
	g2 := greeter{name: "Kamose"}
	g2.emptyTheName()
	fmt.Println(g2.name)

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

// this will change the passed number to a random number using pointers

func changeNumber(number *int) {
	max := 99
	min := 1
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(max-min) + min
	*number = random
	fmt.Println("number changed to " + strconv.Itoa(random))
}

type greeter struct {
	name string
}

// this is a method for struct greeter
// you can't manipulate struct values because it's not a pointer

func (g greeter) greet() {
	if g.name == "" {
		panic("error name cannot be empty string")
	}
	fmt.Println("Hello from method, " + g.name)
}

// this is a method for struct greeter with pointer
// if you use struct with pointer when creating a method you can manipulate the struct values

func (g *greeter) emptyTheName() {
	g.name = ""
}
