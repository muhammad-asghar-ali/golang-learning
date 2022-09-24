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
	// var slic []int = []int{1, 2, 3, 4, 4, 5, 6, 3, 5, 6, 7, 8, 9}

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
	// for i, ele := range slic {
	// 	for j := i + 1; j < len(slic); j++ {
	// 		ele2 := slic[j]
	// 		if ele2 == ele {
	// 			fmt.Println(ele) // 3 4 5 6
	// 		}
	// 	}
	// }

	///////////////////////////////////////////// Maps //////////////////////////////////////

	// var mp map[string]int = map[string]int{
	// 	"apple":  1,
	// 	"pear":   2,
	// 	"oanger": 3,
	// }

	// mp1 := map[string]int{
	// 	"apple":  1,
	// 	"pear":   2,
	// 	"oanger": 3,
	// }

	// mp2 := make(map[string]int)
	// fmt.Println(mp)  // map[apple:1 oanger:3 pear:2]
	// fmt.Println(mp1) // map[apple:1 oanger:3 pear:2]
	// fmt.Println(mp2) // map[]

	// // access values
	// value := mp1["apple"]
	// fmt.Println(value) // 1
	// // change value
	// mp1["apple"] = 900
	// fmt.Println(mp1) // map[apple:900 oanger:3 pear:2]
	// // add values
	// mp1["new add"] = 800
	// fmt.Println(mp1) // map[apple:900 new add:800 oanger:3 pear:2]

	// // delete value
	// delete(mp1, "new add")
	// fmt.Println(mp1) // map[apple:900 oanger:3 pear:2]

	// // check the value
	// value, ok := mp1["apple"]
	// fmt.Println(value, ok) // 900 true

	// val, ok := mp1["john"]
	// fmt.Println(val, ok) // 0 false

	// // length of map
	// fmt.Println(len(mp1)) // 3

	///////////////////////////////////////////// Functions //////////////////////////////////////
	// test()      // function calling
	// test1(12)   // 12
	// add(12, 10) // 22
	// ans := mul(12, 3)
	// fmt.Println(ans) // 36

	// sum, sub := math(2, 3)
	// fmt.Println(sum, sub) // 5 -1

	// mul, div := math1(8, 2)
	// fmt.Println(mul, div) // 16 4

	/////////////////////////////////////////// Advanced Function Concepts & Function Closures ///////////////////////////

	// function as pass as a refenernce to variables
	// ref := test
	// ref() // function calling

	// re := test1
	// re(2) // 2

	// test := func() {
	// 	fmt.Println("inside main function")
	// }
	// test() // inside main function

	// test1 := func(x int) {
	// 	fmt.Println(x)
	// }
	// test1(3) // 3

	// test2 := func(x int) int {
	// 	return x * 2
	// }(8)
	// fmt.Println(test2) // 16

	// test3 := func(x int) int {
	// 	return x * 2
	// }

	// callback(test3) // 14

	// //  Function Closures
	// returnFunction("hello")()
	// a := returnFunction("Good bye")
	// a()

	/////////////////////////////////////// Pointers & Derefrence Operator (& and *) //////////////////////////

	// v := 7
	// fmt.Println(&v) // 0xc000012088

	// c := &v
	// fmt.Println(c)  // 0xc000012088
	// fmt.Println(*c) // 7
	// *c = 10
	// fmt.Println(*c) // 10
	// fmt.Println(v)  // 10

	// toChage := "Hello"
	// var pointer *string = &toChage
	// fmt.Println(pointer, &pointer)  // 0xc000088220 0xc000006028
	// fmt.Println(*pointer, &pointer) // Hello 0xc000006028
	// change1(toChage)
	// fmt.Println(toChage)
	// change(&toChage)
	// fmt.Println(toChage)

	///////////////////////////////////////////////// Structs and Custom Types ////////////////////////////////
	// var p Point = Point{1, 2, false}
	// var p1 Point = Point{3, -6, true}
	// p2 := Point{x: 2}
	// fmt.Println(p)         // {1, 2, false}
	// fmt.Println(p1)        // {3 -6, true}
	// fmt.Println(p.x, p1.x) // 1 3
	// fmt.Println(p2)

	// // pass to as a pointer
	// p3 := &Point{x: 3}
	// fmt.Println(p3) // &{3 0 false}
	// ChangeX(p3)
	// fmt.Println(p3) // &{100 0 false}

	// p4 := Point{x: 3}
	// fmt.Println(p4) // {3 0 false}
	// ChangeNoP(p4)
	// fmt.Println(p4) // {3 0 false}

	// // nested struct
	// c := Circle{4.3, p3}
	// fmt.Println(c) // {4.3 0xc000016228}
	// c1 := Circle{4.3, &Point{4, 5, true}}
	// fmt.Println(c1)               // {4.3 0xc000016228}
	// fmt.Println(c1.point)         // &{4 5 true}
	// fmt.Println(c1.point.isGrade) // true

	//////////////////////////////////////// Struct Methods ///////////////////////////////////////////////
	s1 := Student{"john", []int{60, 70, 80, 80}, 19}
	// fmt.Println(s1) // {john [60 70 80 80] 19}
	// s1.setAge(22)
	// fmt.Println(s1) // {john [60 70 80 80] 22}
	// s1.setAgeNo(21)
	// fmt.Println(s1) //{john [60 70 80 80] 19}

	avg := s1.getAverage()
	fmt.Println(avg) // 72.5

	max := s1.getMaxGrade()
	fmt.Println(max) // 80
}

type Student struct {
	name  string
	grade []int
	age   int
}

func (s *Student) setAge(age int) {
	s.age = age
}

func (s Student) setAgeNo(age int) {
	s.age = age
}

func (s Student) getAverage() float32 {
	sum := 0
	for _, v := range s.grade {
		sum += v
	}
	return float32(sum) / float32(len(s.grade))
}

func (s Student) getMaxGrade() int {
	max := 0
	for _, v := range s.grade {
		if max < v {
			max = v
		}
	}
	return max
}

func ChangeX(pt *Point) {
	pt.x = 100
}

func ChangeNoP(pt Point) {
	pt.x = 100
}

func change(str *string) {
	*str = "changed"
}

func change1(str string) {
	str = "changed"
}

func test() {
	fmt.Println("function calling")
}

func test1(x int) {
	fmt.Println(x)
}

func add(x int, y int) {
	fmt.Println(x + y)
}

func mul(a, b int) int {
	return a * b
}

// return multi values
func math(a, b int) (int, int) {
	return a + b, a - b
}

// return values with labels
func math1(a, b int) (mul int, div int) {
	defer fmt.Println("call when the return") // print when the function returns
	mul = a * b
	div = a / b
	fmt.Println("Before return")
	return
}

//  Function Closures
func returnFunction(x string) func() {
	return func() { fmt.Println(x) }
}

func callback(myFunc func(int) int) {
	fmt.Println(myFunc(7))
}

///////////////////////////////////////////////// Structs
type Point struct {
	x       int
	y       int
	isGrade bool
}

type Circle struct {
	redius float64
	point  *Point
}
