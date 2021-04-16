package arrays

import "fmt"

func Arrays() {
	// declare array
	var grades [2]string
	// you can set and change value of array element using index:
	grades[0] = "khkhkhkhkhkhkhkh"
	grades[1] = "a7a"

	grades[0] = "WTF"

	maleStudents := [2]string{"Mina", "Beshoy"}
	femaleStudents := [...]string{"Maria", "Merette"} // ... means leave the number of values to Go to count them

	arrayOfArrays := [2][3]int{{1, 2, 3}, {3, 2, 1}} // another way: [2][3]int{[3]int{1,2,3}, [3]int{3,2,1}}

	fmt.Printf("%v, %T \n", grades, grades)
	fmt.Printf("%v, %T \n", maleStudents, maleStudents)
	fmt.Printf("%v, %T \n", femaleStudents, femaleStudents)
	fmt.Printf("%v, %T \n", arrayOfArrays, arrayOfArrays)
	fmt.Printf("grades array length: %v, %T \n", len(grades), len(grades))
}
