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

	// for y := 0; y <= 1000; y++ {
	// 	if y != 0 && y%3 == 0 && y%7 == 0 && y%9 == 0 {
	// 		fmt.Println(y)
	// 		continue
	// 		// break
	// 	}
	// 	// else {
	// 	// 	fmt.Println("N")
	// 	// }
	// }

	///////////////////////////////// Switch Statement /////////////////////////////////
	// ans := 2

	// switch ans {
	// case 1:
	// 	fmt.Println("case one match")
	// case 2:
	// 	fmt.Println("case two match")
	// default:
	// 	fmt.Println("case no match")
	// }

	// switch {
	// case ans <= 1:
	// 	fmt.Println("case one match")
	// case ans > 1:
	// 	fmt.Println("case two match")
	// default:
	// 	fmt.Println("case no match")
	// }

	///////////////////////////////////////////////// Arrays ////////////////////////////////////////////////
	// var arr [5]int
	// arr[1] = 400
	// arr[2] = 900
	// fmt.Println(arr) // [0 400 900 0 0]
	// sum := 0
	// arr1 := [3]int{2, 3, 5}
	// fmt.Println(arr1)      // [2 3 5]
	// fmt.Println(len(arr1)) // 3
	// for i := 0; i < len(arr1); i++ {
	// 	sum += arr1[i]
	// }
	// fmt.Println(sum) // 10

	// // 2d arrays
	// arr2D := [2][3]int{{1, 2, 3}, {3, 4, 3}}
	// fmt.Println(arr2D)       // [[1 2 3] [3 4 3]]
	// fmt.Println(arr2D[1][2]) // 3

	///////////////////////////////////////////////// Slices ////////////////////////////////////////////////

	// var x [5]int = [5]int{1, 2, 3, 4, 5}
	// fmt.Println(x) // [1 2 3 4 5]

	// // slice
	// var s []int = x[:]
	// fmt.Println(s) // [1 2 3 4 5]

	// s1 := x[1:3]
	// fmt.Println(s1) // [2 3]

	// // length
	// fmt.Println(len(s1)) // 2
	// // capacity
	// fmt.Println(cap(s1)) // 4

	// // reslice slice
	// fmt.Println(s1[:cap(s1)]) // [2 3 4 5]
	// fmt.Println(s1[1:])       // [3]
	// fmt.Println(s1[2:])       // []

	// slice
	// var a []int = []int{3, 4, 5, 6, 7, 8}
	// fmt.Println(a)          // [3 4 5 6 7 8]
	// fmt.Println(a[:3])      // [3 4 5]
	// fmt.Println(cap(a[:3])) // 6

	// // append to slice
	// a = append(a, 9)
	// fmt.Println(a)
	// b := append(a, 9)
	// fmt.Println(b) // [3 4 5 6 7 8 9]

	// // create slice with make function
	// z := make([]int, 5)
	// fmt.Println(z)      // [0 0 0 0 0]
	// fmt.Printf("%T", z) // []int

	/////////////////////////////////////////// Range & Slice/Array Examples ////////////////////////////////////
	var slic []int = []int{1, 2, 3, 4, 4, 5, 6, 3, 5, 6, 7, 8, 9}

	// for i := 0; i < len(slic); i++ {
	// 	fmt.Println(slic[i])
	// }

	// range

	// for index, element := range slic {
	// 	fmt.Printf("%d: %d \n", index, element)
	// }

	// _ is a anonymous variable
	// for i, ele := range slic {
	// 	for j, ele1 := range slic {
	// 		if ele == ele1 && j > i {
	// 			fmt.Println(ele) // 3 4 5 6
	// 		}
	// 	}
	// }
	for i, ele := range slic {
		for j := i + 1; j < len(slic); j++ {
			ele2 := slic[j]
			if ele2 == ele {
				fmt.Println(ele) // 3 4 5 6
			}
		}
	}
}
