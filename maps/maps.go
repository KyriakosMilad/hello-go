package maps

import "fmt"

func Maps() {
	// declare map
	egyptianClubsFoundedYear := make(map[string]int)
	egyptianClubsFoundedYear = map[string]int{
		"Zamalek SC":         1911,
		"Al Ahly SC":         1907,
		"Ismailly SC":        1924,
		"Al Masry SC":        1920,
		"Ittihad of Alex SC": 1914,
	}

	// shorter way:
	// egyptianClubsFoundedYear := map[string]int{
	// 	"Zamalek SC": 1911,
	// 	"Al Ahly SC": 1907,
	// 	"Ismailly SC": 1924,
	// 	"Al Masry SC": 1920,
	// 	"Ittihad of Alex SC": 1914,
	// }

	fmt.Printf("%v, %T \n", egyptianClubsFoundedYear, egyptianClubsFoundedYear)

	// you can get value from map using the key:
	println(egyptianClubsFoundedYear["Zamalek SC"])

	// notice if the key doesn't exist value wil be 0, if you want to make sure if key exists and it's value doesn't equal 0 you can do this:
	zamalekSCFoundedYear, ok := egyptianClubsFoundedYear["Zamalek SC"]
	pyramidsFCFoundedYear, ok := egyptianClubsFoundedYear["Pyramids FC"]

	// if record exists: ok will be true, if not: ok will be false
	println(zamalekSCFoundedYear, ok)  // 1911 true
	println(pyramidsFCFoundedYear, ok) // 0 false

	// add to map
	egyptianClubsFoundedYear["Pyramids FC"] = 2008
	println(egyptianClubsFoundedYear)

	// delete from map
	delete(egyptianClubsFoundedYear, "Pyramids FC")
	println(egyptianClubsFoundedYear)

	// get length of map
	println(len(egyptianClubsFoundedYear))
}
