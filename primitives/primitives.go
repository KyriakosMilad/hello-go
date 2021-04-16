package primitives

import "fmt"

// notice:-
// every declared variable without value, it's value is 0 (= false)
// can't make sum of different types

// integer types:-
// int8   = -128 : 127
// int16  = -32,768 : 32,767
// int32  = -2,147,483,648 : 2,147,483,647
// int64  = -9,223,372,036,854,775,808 : 9,223,372,036,854,775,807

// unsigned integer types:
// uint8  =  0 : 255
// uint16 =  0 : 65,535
// uint32 =  0 : 4,294,967,295

// floating number types:
// float32 = ±1.18E-38 : ±3.4E38
// float64 = ±2.23E-308 : ±1.8E308

// complex number:
// complex64  = float32 + float32
// complex128 = float64 + float64

// complex number functions:
// imag()    : get the imaginary part from a complex number
// real()    : get the real part from a complex number
// complex() : create a complex number

// string types:
// UTF-8:
// you can get specific digit's value using 0-index like an array, example: string[0] will give first digit's value from the variable string
// you can't change specific digit value using the array index, you have to change the variable's value
// Rune (UTF-32): https://stackoverflow.com/questions/19310700/what-is-a-rune

func Primitives() {
	// to make sum of variables must all be of the same type and it will give value of the same type
	a := 10 // binary: 1010
	b := 3  // binary: 0011

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b) // notice dividing integer by integer will give integer not float
	fmt.Println(a % b)

	// comparison
	println("comparison:")
	fmt.Println(a & b)  // means record 1 if both digits equals 1: (0010 = 2)
	fmt.Println(a | b)  // means record 1 if one of  the digits equals 1: (1011 = 11)
	fmt.Println(a ^ b)  // means record 1 if if one digit equals 1 and the other doesn't: (1001 = 9)
	fmt.Println(a &^ b) // means record 1 if both digits doesn't equal 1: (0100 = 8)

	// shifting
	println("shifting 8:")
	i := 8              // 2^3
	fmt.Println(i << 3) // 2^3 * 2^3 = 2^6 = 64
	fmt.Println(i >> 3) // 2^3 / 2^3 = 2^0 = 1

	println("shifting 4:")
	j := 4              // 2^2
	fmt.Println(j << 3) // 2^2 * 2^3 = 2^5 = 32
	fmt.Println(j >> 3) // 2^2 / 2^3 = 4/8 = 0.5 = 0 (because we are shifting integer the result value will be of the same type)

	// boolean
	t := true
	var f bool = false

	println("boolean:")
	fmt.Printf("type: %T, value: %v \n", t, t)
	fmt.Printf("type: %T, value: %v \n", f, f)
}
