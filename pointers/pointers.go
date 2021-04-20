package pointers

import "fmt"

// *int (pointer type) means pointer to integer value
// &a   means address of value a in the memory
// *&a  means get the value that this address (&a) holds in the memory

func Pointers() {
	var a int = 5
	var b *int = &a
	fmt.Println(a, b, *b)

	a = 3
	fmt.Println(a, b, *b)

	// each element in array, slice, map, struct has it's own address in the memory
	array := [3]int{1, 2, 3}
	fmt.Println(array, &array[0])

	// notice you can't make math of address in go, if you really need to do this check this package "unsafe" https://golang.org/pkg/unsafe/
	//slice := []int{1,2,3}
	//fmt.Println(slice, &slice[2] - 4) // wrong

	var ms *myStruct
	ms = new(myStruct) // you can also create struct without new method like this &myStruct{foo: "hi"}
	fmt.Println(ms, *ms)
	(*ms).foo = "hi"
	// ms.foo = "hi" // you can always do it like this too and go will understand what are you trying to do
	fmt.Println(ms, *ms)
}

type myStruct struct {
	foo string
}
