package main

import (
	"fmt"
	//"time"
)

// ******simple function calling and const

// const(
//     z = iota
//     f
//     s
//     t
//     )
// func main() {
//   a := 10
//   b := 120
//   fmt.Println("Hello World!")
//   fmt.Println(test(a,b))
//   fmt.Println(z, f, s, t)
// }

// func test(number1 int, number2 int) (int, error){
//     var number3 int
//     number3 = number1 + number2
//     return number3, nil
// }
// **//////simple function calling and const///////

// ****** if else
// const(
//         z = iota
//         f
//         s
//         t
//         )
// func main() {
//   if z == 3 {
//     fmt.Println("z is true")
//   }else if f == 3 {
//     fmt.Println("f is true")
//   }else if s == 2 {
//     fmt.Println("s is true")
//   }else{
//     fmt.Println("t is true")
//   }
//   fmt.Println(z, f, s, t)
// }
// ******////// if else//////

// ****** switch case
// func main() {
//     i := 2
// fmt.Print("Write ", i, " as \n"   )
// switch i {
// case 1:
//     fmt.Println("one")
// case 2:
//     fmt.Println("two")
// case 3:
//     fmt.Println("three")
// }
// switch time.Now().Weekday() {
// case time.Saturday, time.Sunday:
//     fmt.Println("It's the weekend")
// default:
//     fmt.Println("It's a weekday")
// }
// }
// ******////// switch case//////


// ******pointers
// Before we start there are two important operators which we will use in pointers i.e. 
// * Operator also termed as the dereferencing operator used to declare pointer variable and access the value stored in the address.
// & operator termed as address operator used to returns the address of a variable or to access the address of a variable to a pointer.
// func main() {
//     // taking a normal variable
//     var x int = 5748
//     // declaration of pointer
//     var p *int
//     // initialization of pointer
//     p = &x
//     // displaying the result
//     fmt.Println("Value stored in x = ", x)
//     fmt.Println("Address of x = ", &x)
//     fmt.Println("Value stored in variable p = ", p)
// }
// ******////// pointers//////

// ******methods
type MyStruct struct {
	x int
	y int
}
//add is a method that has access of all fields of struct
func (o MyStruct) add() int {
	return o.x + o.y
}

func main() {
	myStruct := MyStruct{3, 2}
	result := myStruct.add()
	fmt.Println(result) // prints 5
}
// ******//////methods//////