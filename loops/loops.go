package loops

import "fmt"

func Loops() {
	// initialise inline
	for i := 0; i <= 3; i++ { // variables to work with; condition to determine when to stop; what happens after
		fmt.Println(i)
	}

	// you can work with multiple variables like this
	for i, j := 0, 0; i < 2 && j > 0; i, j = i+1, j+6 {
		fmt.Println(i, j)
	}

	// use variable created before
	k := 0
	for ; k < 6; k++ { // leave the first statement empty and don't forget to put semicolon
		fmt.Println(k)
	}

	// skip current element and continue to next one
	for i := 0; i <= 3; i++ { // variables to work with; condition to determine when to stop; what happens after
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}

	// there is no while loop in go but you can do it like this
	a := 0
	for { // well stop when a equals four
		fmt.Println(a)
		if a > 3 {
			break
		}
		a++
	}

	// break multiple loops using labels
MainLoop:
	for o := 0; o < 4; o++ {
		for p := 0; p == 2; p++ {
			if (p + o) == 4 {
				//break // if you do it like this it will break the nearest loop (p loop)
				break MainLoop // this well break all inside the label MainLoop
			}
		}
	}

	// loop through maps, arrays, slices, string letters
	s := []string{"a", "b", "c"}

	for key, value := range s {
		println(key, value)
	}

	// you can ignore value or key if you are not going to use them by replacing one of them by _ like this
	for _, value := range s { // only value available
		println(value)
	}
}
