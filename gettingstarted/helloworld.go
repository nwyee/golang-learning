package gettingstarted

import "fmt"

// create own type
type jamesBond int

func HandsOnEx() {
	fmt.Println("[*] Welcome From Getting Started ")
	printEx()
	fundInt()
	fundConstants()
}

/**
  https://golang.org/pkg/fmt/
*/
func printEx() {
	fmt.Println("[-] fmt print")

	var x jamesBond
	x = 7
	s := fmt.Sprintf("00%v\t%T", x, x)
	fmt.Println("hello world\n", s)

	s = "Hello, 世界"
	// the uninterpreted bytes of the string or slice
	fmt.Printf("%s\n", s)
	// a double-quoted string safely escaped with Go syntax
	fmt.Printf("%q\n", s)
	// Unicode format
	for i, v := range s {
		fmt.Printf("%#U \t %d", v, i)
	}

}

/*
 Integer types :
    - uint8, uint16, uint32, uint64
    - int8, int16, int32, int64
Float types :
	- float32, float64
Generally we should stick with float64 when working with the floating point numbers.

*/
func fundInt() {
	fmt.Println("[-] Fundamental Integer, Float, String")
	fmt.Println("1 + 1 =", 1+1)
	// [.0] is a floating point number
	fmt.Println("1 + 1 =", 1.0+1.0)
	// String
	/*
		indexed is starting at 0 .
		2nd character will be printed out
		however, you will see '101' instead of 'e'
	*/
	fmt.Println("Hello World"[1])
}

/*
Constants
	- a simple, unchanging value
	- only exist at compile time
	- there are TYPED and UNTYPED constants
	- UNTYPED
			- do not have a fixed type
			- can be implicitly converted by the compiler

*/
func fundConstants() {
	fmt.Println("[-] TYPED / UNTYPED constants")
	// UNTYPED
	const x = 40
	// TYPED
	const typedY float64 = 43.2

	type hotDog int
	type hotCat float64

	fmt.Println(x)
	fmt.Println(typedY)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", typedY)

	fmt.Printf("%T\n", hotDog(x))
	// hotDog(typeY) will be error, so let me comment out
	// Cannot convert an expression of the type 'float64' to the type 'hotDog'
	//fmt.Printf("%T\n", hotDog(typedY))

	// UNTYPED constant easier to use
	// no need to conversion to use
	fmt.Printf("%T\n", hotCat(x))
	fmt.Printf("%T\n", hotCat(typedY))

}
