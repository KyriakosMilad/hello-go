package conditions

import "fmt"

func Conditions() {
	// if condition
	a := 5
	if a == 5 {
		fmt.Println(a)
	}

	// and
	if a == 5 && a < 10 { // true: a == 5 and a < 10
		fmt.Println(a)
	}

	// or
	if a != 5 || a < 10 { // true a != 5 BUT a < 10
		fmt.Println(a)
	}

	// else if and else
	if a == 5 { // true
		fmt.Println("if", a)
	} else if a < 10 { // true (this will overlay the the above conditions)
		fmt.Println("else if", a)
	} else { // only if all above conditions doesn't meet
		fmt.Println("conditions doesn't meet")
	}

	egyptianClubsFoundedYear := map[string]int{
		"Zamalek SC":         1911,
		"Al Ahly SC":         1907,
		"Ismailly SC":        1924,
		"Al Masry SC":        1920,
		"Ittihad of Alex SC": 1914,
	}

	// initializer if (initialize and check variable inline)
	if zamalekSCFoundedYear, ok := egyptianClubsFoundedYear["Zamalek SC"]; ok {
		fmt.Println(zamalekSCFoundedYear)
	}

	// switch case
	b := 2
	switch b {
	case 1: // if b == 1 (false)
		fmt.Println("case equals one")
	case 2: // if b == 1 (true)
		fmt.Println("case equals two")
	default: // if none of the above conditions met
		fmt.Println("conditions doesn't meet")
	}

	// or
	switch b {
	case 1, 3, 99: // if b == 1 or b == 3 or b == 99 (true)
		fmt.Println("case equals one or three or ninety nine")
	default: // if none of the above conditions met
		fmt.Println("conditions doesn't meet")
	}

	// initialise inline
	switch c := 50 + 5; c {
	case 49: // if c == 49 (false)
		fmt.Println("case equals forty nine")
	case 55: // if c == 55 (true)
		fmt.Println("case equals fifty five")
	default: // if none of the above conditions met
		fmt.Println("conditions doesn't meet")
	}

	k := 109
	// notice: switch breaks whenever condition is met (doesn't continue the rest of the cases)
	// to make it continue through use the word 'fallthrough'
	switch {
	case k == 109: // if k == 109 (true)
		fmt.Println("case equals one hundred and nine")
		fallthrough // tells go continue and don't break
	case k < 150: // if k < 150 (false)
		fmt.Println("case lower than 1one hundred and fifty")
	default: // if none of the above conditions met
		fmt.Println("conditions doesn't meet")
	}

	// break
	switch f := 9 - 5; f {
	case 4: // if f == 4 (true)
		fmt.Println("case equals four")
		break // tells go don't continue
		fmt.Println("case equals four again")
	case 1: // if f == 1 (false)
		fmt.Println("case equals one")
	default: // if none of the above conditions met
		fmt.Println("conditions doesn't meet")
	}
}
