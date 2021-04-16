package variables

import "fmt"

//  notice:
//  naming variable for in-package scope first letter small
//  naming variable for out-package scope first letter capital
//  all variables must be used
// can't redeclare variables but can shadow them

func Variables() {
	// ways to declare variable
	var i int
	i = 3

	j := 22.5

	var k string = "str"

	fmt.Printf("var value: %v, var type: %T \n", i, i)
	fmt.Printf("var value: %v, var type: %T \n", j, j)
	fmt.Printf("var value: %v, var type: %T \n", k, k)
}
