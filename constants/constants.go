package constants

import "fmt"

// 	naming constant is like variable:
//  naming constant for in-package scope first letter small
//  naming constant for out-package scope first letter capital

//  all constants must be used

//  can't redeclare constants but can shadow them

// iota:-

const (
	a = iota // const a int = 0
	b = iota // const b int = 1
	c = iota // const c int = 2
	d = iota // const d int = 3
)

// every const or var block hat it owns iota value, notice in the next block iota starts at 0

const (
	h = iota // const h int = 0
	i = iota // const i int = 1
	j = iota // const j int = 2
	k = iota // const k int = 3
)

// you can make sum of iota value

const (
	x = iota + 5 // const x int = 5 (0 + 5)
	y = iota     // const y int = 1
	z = iota + 8 // const z int = 10 (2 + 8)
)

func Constants() {
	// declare constant
	const g int = 3
	const f = 3.44

	// wrong way to declare constant:-
	// const myConst float64 = math.Sin(1.57)
	// it's wrong because constant's value must be determined before runtime, and functions being called at runtime

	fmt.Printf("const value: %v, const type: %T \n", i, i)
	fmt.Printf("const value: %v, const type: %T \n", a, a)
	fmt.Printf("const value: %v, const type: %T \n", d, d)
	fmt.Printf("const value: %v, const type: %T \n", h, h)
	fmt.Printf("const value: %v, const type: %T \n", k, k)
	fmt.Printf("const value: %v, const type: %T \n", x, x)
	fmt.Printf("const value: %v, const type: %T \n", y, y)
	fmt.Printf("const value: %v, const type: %T \n", z, z)
}
