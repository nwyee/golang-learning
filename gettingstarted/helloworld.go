package gettingstarted

import "fmt"

// create own type
type jamesBond int

func HandsOnEx() {
	fmt.Println("[*] Welcome From Getting Started ")
	printEx()
	fundInt()
	fundConstants()
	fundControls()
	fundArrays()
	funcDefer()
	funcPanicRecover()
}

func funcPanicRecover() {
	fmt.Println("[-] Panic & Recover")
	defer func() {
		// recover() return value passed to panic call
		str := recover()
		fmt.Println("recover() ", str)
	}()
	panic("PANIC")
}

func funcDefer() {
	fmt.Println("[-] Defer Function Usage ")
	fmt.Println("Actually second() func called first place.")
	defer second()
	first()
}

func first() {
	fmt.Println("1st Sample Method")
}

func second() {
	fmt.Println("2nd Sample Method")
}

func fundArrays() {
	fmt.Println("[-] Arrays, Slices and Maps ")
	var x [5]float64
	x[4] = 100
	x[0] = 90
	x[1] = 80
	x[3] = 70
	x[2] = 50
	fmt.Println(x)

	var total float64 = 0
	/*
		A single _ (underscore) is used to tell the compiler that we don't need this.
	*/
	for _, value := range x {
		total += value
	}

	fmt.Println("Average : ", total/float64(len(x)))

	fmt.Println("*[-] Slice")
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	slice3 := make([]int, 3)
	slice4 := make([]int, 3, 9) // len(slice4)=0, cap(slice4)=5
	copy(slice3, slice2)
	fmt.Println(slice1, slice2, slice3, slice4)

	fmt.Println("*[-] Maps")
	/**
	Map is unordered collection of key-value pairs.
	- also known as associative array, hash table or a dictionary
	*/
	// xMap is a map of "string"s to "int"s.
	// xMap must be initialized before used. if not,  runtime error: assignment to entry in nil map
	xMap := make(map[string]int)
	xMap["key-1"] = 10
	xMap["key-2"] = 20
	fmt.Println(xMap)
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"
	fmt.Println(elements)
	fmt.Println("delete Nitrogen")
	delete(elements, "N")
	fmt.Println(elements)
	/**
	Technically a map returns the zero value for the value type
	(which for strings is the empty string).
	*/
	fmt.Println("element with key 'UNKNOWN' : ", elements["UNKNOWN"])
	// Better way to check
	if name, ok := elements["UNKNOWN"]; ok {
		fmt.Println(name, ok)
	} else {
		fmt.Println("Sorry there is no key with 'UNKNOWN' ")
	}
	moreMapEx()
}

func moreMapEx() {
	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": map[string]string{
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gas",
		},
	}

	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}

}

func fundControls() {
	fmt.Println("[-] ControlsStructures ")
	i := 1
	fmt.Println("Print 1 to 10 Using For Loop")
	fmt.Println("Print even / odd Using If / else")
	var flag string
	for i <= 10 {
		if i%2 == 0 {
			flag = "even"
		} else {
			flag = "odd"
		}
		fmt.Print(i, flag, "\t")
		i += 1
	}
	fmt.Println()
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
