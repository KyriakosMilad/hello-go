package slices

import "fmt"

// slice capacity vs slice length: https://tour.golang.org/moretypes/11
// The length of a slice is the number of elements it contains.
// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

func Slices() {
	// declare slice
	var grades = make([]string, 2, 100) // type, length, capacity
	// you can change value of slice element using index:
	grades[0] = "khkhkhkhkhkhkhkh"
	grades[1] = "a7a"

	fmt.Printf("%v, %T \n", grades, grades)
	fmt.Printf("grades slice length: %v, %T \n", len(grades), len(grades))
	fmt.Printf("grades slice capacity: %v, %T \n", cap(grades), cap(grades))

	// push to slice
	grades = append(grades, "wOw")

	fmt.Printf("%v, %T \n", grades, grades)
	fmt.Printf("grades slice length: %v, %T \n", len(grades), len(grades))
	fmt.Printf("grades slice capacity: %v, %T \n", cap(grades), cap(grades))

	femaleStudents := []string{"Shery", "Nefertiti", "Mirna", "Marina", "Isis", "Teti"}
	maleStudents := []string{"Abanoub", "Bakhom", "Fam", "Ramses", "Kamus", "Ahmus"}

	// [index:index-1]
	a := maleStudents[:]   // slice of all elements
	b := maleStudents[2:]  // slice from 3rd elements to end
	c := maleStudents[:4]  // slice first 4 elements
	d := maleStudents[2:5] // slice from 3rd element to 5th element
	e := maleStudents[4]   // get the 5th element

	// remove from slice
	// first element
	femaleStudents = femaleStudents[1:]
	// last element
	femaleStudents = femaleStudents[:len(femaleStudents)-1]
	// remove element using index (we want to remove 3rd element (Marina) - which it's index is 2 - from ["Nefertiti", "Mirna", "Marina", "Isis"])
	femaleStudents = append(femaleStudents[:2], femaleStudents[3:]...)

	fmt.Printf("%v, %T \n", femaleStudents, femaleStudents)
	fmt.Printf("femaleStudents slice length: %v, %T \n", len(femaleStudents), len(femaleStudents))
	fmt.Printf("femaleStudents slice capacity: %v, %T \n", cap(femaleStudents), cap(femaleStudents))

	fmt.Printf("%v, %T \n", maleStudents, maleStudents)
	fmt.Printf("%v, %T \n", a, a)
	fmt.Printf("%v, %T \n", b, b)
	fmt.Printf("%v, %T \n", c, c)
	fmt.Printf("%v, %T \n", d, d)
	fmt.Printf("%v, %T \n", e, e)
}
