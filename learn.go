package main

import (
	"fmt"
)

func main() {
	///////////////////////////////// Variables & Data Types
	// var s string = "String"
	// var number uint = 260
	// // overflow the number
	// // var number uint8 = 260
	// // deafult number type is int
	// var i int = -12
	// var num uint = 12
	// var f float64 = 12.4
	// var b bool = true

	// // println means print on new line
	// fmt.Println(s)
	// fmt.Println(number)
	// fmt.Println(i)
	// fmt.Println(num)
	// fmt.Println(f)
	// fmt.Println(b)
	// fmt.Println(number + num)

	// fmt.Println("Hello World")

	///////////////////////////////// Assignment Expression & Implicit vs Explicit
	// // explicit define

	// var ex uint = 16

	// // implicit define
	// var numb = 20000
	// // := expression assignment opreator
	// imp := 32
	// // printf means print format
	// fmt.Printf("%T \n", numb)
	// fmt.Println(ex)
	// fmt.Printf("%T", imp)

	////////////////////////////////////// Printing to Console & fmt
	// var nu uint64
	// var bl bool
	// var st string

	// fmt.Println(nu)
	// fmt.Println(bl)
	// fmt.Println(st)

	// fmt.Printf("hello %% %T %v", 10, 10) // hello % int 10
	// var bl = true
	// fmt.Printf("hello %T %t \n", bl, bl)

	// // binary, decimla, otcal, hexadecimal
	// var integer = 1025
	// var integer1 = 3245
	// fmt.Printf("hello %T %b \n", integer, integer)
	// fmt.Printf("hello %T %o \n", integer, integer)
	// fmt.Printf("hello %T %d \n", integer, integer)
	// fmt.Printf("hello %T %x \n", integer, integer)
	// fmt.Printf("hello %T %x \n", integer1, integer1)
	// fmt.Printf("hello %T %X \n", integer1, integer1)

	// // floating point number
	// var floating = 23.444445656789321
	// fmt.Printf("hello %T %e \n", floating, floating) // 2.344445e+01
	// fmt.Printf("hello %T %f \n", floating, floating) //  23.444446
	// fmt.Printf("hello %T %F \n", floating, floating) //  23.444446
	// fmt.Printf("hello %T %g \n", floating, floating) //  23.44444565678932 for large exponent

	// // String and widt and preiciion
	// var string1 = "john"
	// fmt.Printf("hello %T %s \n", string1, string1)           // hello string john
	// fmt.Printf("hello %T %q \n", string1, string1)           // hello string "john"  (with quotation)
	// fmt.Printf("hello %T %9q \n", string1, string1)          // hello string    "john"
	// fmt.Printf("hello %T %-9q is cool \n", string1, string1) // hello string "john"    is cool

	// var floating1 = 23.444445656789321
	// fmt.Printf("hello %T %.2f \n", floating1, floating1) // hello float64 23.44
	// fmt.Printf("hello %T %.f \n", floating1, floating1)  // hello float64 23

	// /////// padding
	// var integer2 = 23
	// fmt.Printf("hello %T %05d \n", integer2, integer2) // hello int 00023

	//////////////////////////////////// Sprint

	// var out string = fmt.Sprintf("hello %T %05d \n", 34, 34)
	// print(out)

	////////////////////////////////////// Console Input (Bufio Scanner) & Type Conversion
	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Printf("Type the year you were born: ")
	// scanner.Scan()
	// input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	// fmt.Printf("you age are %d years old.", 2022-input)

	//////////////////////////////////// Arithmetic Operators & Math
	var num1 float64 = 9
	var num2 float64 = 4
	sum := num1 + num2
	diff := num1 - num2
	mul := num1 * num2
	div := num1 / num2

	fmt.Printf("Sum: %g \n", sum)
	fmt.Printf("Diff: %g \n", diff)
	fmt.Printf("Mul: %g \n", mul)
	fmt.Printf("Div: %g \n", div) // 2.5

}
