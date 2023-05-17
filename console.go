package main

import "fmt"

func main() {

	// fmt.Printf("Hello %T %v", 10, 10)
	// fmt.Printf("Hello %t", true)

	// //----> Binary representation of the number
	// fmt.Printf("Number : %b", 1025)
	// //----> Octal representation of the number
	// fmt.Printf("Number : %o", 897)
	// //----> Decimal representation of the number
	// fmt.Printf("Number : %d", 897)
	// //hexadecimal representation of the number
	// fmt.Printf("Number : %x", 897)

	// //------> Scientific notation
	// fmt.Printf("Number : %e", 2.34454534334343)
	// //------> Decimal with no exponent
	// fmt.Print("Number : %g", 3.3534243252)

	//-----> Strings
	fmt.Printf("Number : %.2f is cool ", 3.373535326463757)
}

//FMT cheat sheet

//General
// %v (value in default format)
// %T (type)
// %% (literal %)

//Boolean
//%t(true or false)

//Intger

// %b (base 2)
// %o (base 8)
// %d (base 10)
// %x (base 16)

//Floating points
//%e (scientific notation)
//%f / %F(decimal no exponent)
// %g (for large exponents)

//Strings
// %s (default)
// %q(double quoted string)

// \n for new line
// /t for tab or space
