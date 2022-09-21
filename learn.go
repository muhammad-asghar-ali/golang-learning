package main

import "fmt"

func main() {
	///////////////////////////////// Variables & Data Types ///////////////////////////////////////
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

	///////////////////////////////// Assignment Expression & Implicit vs Explicit ///////////////////////////////
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

	////////////////////////////////////// Printing to Console & fmt //////////////////////////////////////////
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

	//////////////////////////////////// Sprint ////////////////////////////////////////

	// var out string = fmt.Sprintf("hello %T %05d \n", 34, 34)
	// print(out)

	////////////////////////////////////// Console Input (Bufio Scanner) & Type Conversion /////////////////////////////////
	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Printf("Type the year you were born: ")
	// scanner.Scan()
	// input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	// fmt.Printf("you age are %d years old.", 2022-input)

	//////////////////////////////////// Arithmetic Operators  /////////////////////////////////////////
	// var num1 int = 9
	// var num2 int = 4
	// sum := num1 + num2
	// diff := num1 - num2
	// mul := num1 * num2
	// div := num1 / num2
	// mod := num1 % num2
	// m := 2*(num1*num2) - num1/2

	// fmt.Printf("Sum: %d \n", sum)
	// fmt.Printf("Diff: %d \n", diff)
	// fmt.Printf("Mul: %d \n", mul)
	// fmt.Printf("Div: %d \n", div) // 2
	// fmt.Printf("Mod: %d \n", mod)
	// fmt.Printf("Arithmetic expression: %d \n", m)

	////////////////////////////////////////// Conditions & Boolean Expressions ////////////////////////////////////
	// a := 6
	// b := 7
	// fmt.Printf("%t \n", a < b)  // true
	// fmt.Printf("%t \n", a > b)  // false
	// fmt.Printf("%t \n", a == b) // false
	// fmt.Printf("%t \n", a <= b) // true
	// fmt.Printf("%t \n", a >= b) // false
	// fmt.Printf("%t \n", a != b) // true

	///////////////////////////////////////// Chained Conditionals (AND, OR, NOT) ////////////////////////////////

	// x := 4
	// y := 9

	// fmt.Printf("%t \n", x < y || x > y)   //true
	// fmt.Printf("%t \n", x < y && x > y)   // false
	// fmt.Printf("%t \n", x == y || x == y) // false
	// fmt.Printf("%t \n", x == y && x == y) // false
	// fmt.Printf("%t \n", x <= y || x >= y) // true
	// fmt.Printf("%t \n", x <= y && x >= y) // false
	// fmt.Printf("%t \n", x != y && x != y) // true
	// fmt.Printf("%t \n", x != y || x != y) // true

	// fmt.Printf("%t \n", false || true && false)    // false
	// fmt.Printf("%t \n", (false || true) && false)  // false
	// fmt.Printf("%t \n", false || !(true && false)) // true

	// fmt.Printf("%t \n", true || false && true)    // true
	// fmt.Printf("%t \n", true || (false && true))  // true
	// fmt.Printf("%t \n", true || !(false && true)) // true

	///////////////////////////////////////////////// If, Else If, Else ////////////////////////////////////////////

	// age := 16
	// if age < 17 {
	// 	fmt.Println("you can not ride")
	// 	fmt.Printf("wait %d years to ride", 18-age)
	// } else if age >= 14 && age <= 17 {
	// 	fmt.Println("you can ride with parent")
	// } else {
	// 	fmt.Println("you can ride")
	// }

	///////////////////////////////////////////////// For Loops & While Loops ////////////////////////////////////

	// for loop
	// x := 3
	// for x < 5 {
	// 	fmt.Println(x)
	// 	x++
	// }

	// for y := 0; y <= 5; y++ {
	// 	fmt.Println(y)
	// }

	for y := 0; y <= 1000; y++ {
		if y != 0 && y%3 == 0 && y%7 == 0 && y%9 == 0 {
			fmt.Println(y)
			continue
			// break
		}
		// else {
		// 	fmt.Println("N")
		// }
	}
}
