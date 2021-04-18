package controls

import (
	"fmt"
)

func Controls() {
	fmt.Println("first")
	defer fmt.Println("defer -> second") // use defer to delay operation execution to be after the function execution
	fmt.Println("third")

	fmt.Println("one")
	//panic("something went wrong!") // use panic to stop the program, notice any defer operation will be executed before panic
	fmt.Println("two")

	// you can recover from panic using recover() function with defer like this
	//fmt.Println("one")
	//defer func() { // this called anonymous function will know about later in functions package
	//	if err := recover(); err != nil {
	//		log.Fatal(err) // use this continue the app knowing u can handle this error
	//		panic("stop permanently")
	//	}
	//}()
	//panic("something went wrong!") // use panic to stop the program, notice any defer operation will be executed before panic
	//fmt.Println("two")
}
