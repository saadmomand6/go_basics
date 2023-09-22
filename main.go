package main

import (
	//"errors"
	"fmt"
	"runtime"
	"sync"
	//"time"
	//"time"
	//"sort"
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
// type MyStruct struct {
// 	x int
// 	y int
// }
// //add is a method that has access of all fields of struct
// func (o MyStruct) add() int {
// 	return o.x + o.y
// }

// func main() {
// 	myStruct := MyStruct{3, 2}
// 	result := myStruct.add()
// 	fmt.Println(result) // prints 5
// }
// ******//////methods//////

// ******for loop with range

// func main() {
// 	mylist := [] int {1,2,3,4,5};
// 	for i, j := range mylist{
// 		fmt.Println(i,j);
// 	}
// }

// ******//////for loop with range//////'

// ****** closure

// func company() func() int {
// 	a := 0;
// 	return func() int {
// 		a++;
// 		return a;
// 	}
// }
// func main() {
// 	 val  := company() //=> return func() int { a++; return a; }
// 	 fmt.Println(val())// here this function is called  which give val a = 1
// 	 fmt.Println(val()) // again the above process is executred and a = 2
// 	 fmt.Println(val()) // herre a = 3
// 	 fmt.Println(val())// herre a = 4

// 	 v := company() // now here the above function is called from scratch
// 	 //as it is assigned to a new variable
// 	 fmt.Println(v())
//}

// ******////// closure //////

// ****** Struct
// it is a collection of multiple data types
// type user struct {
// 	 name string
// 	 email string
// 	 age int
// 	 address address
// }
// type address struct {
// 	house string
// 	street string
// 	number int
// }
// func main(){
// 	var u user // here user is type
// 	fmt.Println(u);
// 	u = user{name:"saad",email:  "myemail",age:  22}
// 	fmt.Println(u);
// 	u = user{name :"saad",email:  "myemail",age:  22, address: address{
// 		house: "dbdfb",
// 		street: "uhuhu",
// 		number: 90232342,
// 	}}

// 	fmt.Println(u);
// }
// ****** Anoynimous Struct
// func main(){
// 	nameing := struct{
// 		name string
// 	}{name: "wawa"}
// 	fmt.Print(nameing)
// }
// ******////// Anoynimous Struct//////

// ****** Return and handle an error

// func main() {
//     // Get a greeting message and print it.
//     hi, _ := Hello("jnjn")
//     fmt.Println(hi)
// }
// // Hello returns a greeting for the named person.
// func Hello(name string) (string, error) {
//     // If no name was given, return an error with a message.
//     if name == "" {
//         return "", errors.New("empty name")
//     }
//     // If a name was received, return a value that embeds the name
//     // in a greeting message.
//     message := fmt.Sprintf("Hi, %v. Welcome!", name)
//     return message, nil
// }

// ******////// Return and handle an error //////

// ****** Slices
// func main() {
// array := [5]int{1, 2, 3, 4, 5,}
// slice := array[0:4]

// fmt.Println("Array: ", array)
// fmt.Println("Slice: ", slice)
// fmt.Println(len(slice))
// fmt.Println(cap(slice))
// slice = append(slice, 10)
// fmt.Println(slice)
// fmt.Println(len(slice))
// fmt.Println(cap(slice))
// slice = append(slice, 19)
// fmt.Println(slice)
// fmt.Println(len(slice))
// fmt.Println(cap(slice))

// slicebymake := make([]int, 4,7)
// fmt.Println(slicebymake)
// fmt.Println(len(slicebymake))
// fmt.Println(cap(slicebymake))

//}
// ******/////// Slices//////

// ****** Sorting Slices
// func main(){
// slice1:= []int {32,-89,2,-9,0,-64,87,10}
// fmt.Println(slice1)
// sort.Ints(slice1)
// fmt.Println(slice1)
// }
// ******/////// Sorting Slices//////

// ****** Panic and Recovery
// func Divide(a, b int) int {
// 	// Defer a function to handle recovery.
// 	defer func() {
// 		if r := recover(); r != nil {
// 			// Recovered from a panic. Print an error message.
// 			fmt.Println("Recovered from panic:", r)
// 		}
// 	}()

// 	if b == 0 {
// 		// Divide by zero is not allowed. Panic with an error message.
// 		panic("Division by zero")
// 	}

// 	// Perform the division.
// 	return a / b
// }

// func main() {
// 	// Call the Divide function with various inputs.
// 	result1 := Divide(6, 3)
// 	fmt.Println("Result1:", result1)

// 	result2 := Divide(8, 0) // This will panic.
// 	fmt.Println("Result2:", result2) // This line won't be reached due to the panic.

// 	result3 := Divide(10, 2)
// 	fmt.Println("Result3:", result3)
// }

// ******///////Panic and Recovery//////

// ****** ConCureency
// We define two functions, printNumbers and printLetters, which print numbers and letters respectively with a small delay between each print.
// In the main function, we start both printNumbers and printLetters as goroutines using the go keyword. This allows them to run concurrently.
// We use time.Sleep to give the goroutines some time to execute before the main function exits.
// When you run this program, you'll see numbers and letters printed concurrently, indicating concurrent execution.
// func printNumbers() {
// 	for i := 1; i <= 5; i++ {
// 		fmt.Printf("%d ", i)
// 		time.Sleep(1 * time.Second)
// 	}
// }

// func printLetters() {
// 	for _, char := range "ABCDE" {
// 		fmt.Printf("%c ", char)
// 		time.Sleep(1 * time.Second)
// 	}
// }

// func main() {
// 	go printNumbers() // Start printing numbers concurrently.
// 	go printLetters() // Start printing letters concurrently.

// 	// Sleep for a while to allow goroutines to execute.
// 	time.Sleep(9 * time.Second)

// 	fmt.Println("\nMain function is done.")
// }

// ******///////ConCureency//////

// ***** parallelism

// We use runtime.GOMAXPROCS to specify the maximum number of CPU cores to use for parallelism. In this case, we set it to 2 to utilize two CPU cores.
// We define two functions, printNumbers and printLetters, which print numbers and letters, respectively.
// In the main function, we use a sync.WaitGroup to wait for the goroutines to finish.
// We start two goroutines, one for each function. These goroutines can execute in parallel on the available CPU cores.
// The wg.Wait() call waits for both goroutines to complete before allowing the main function to exit.
func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
}

func printLetters() {
	for _, char := range "ABCDE" {
		fmt.Printf("%c ", char)
	}
}

func main() {
	// Set the maximum number of CPU cores to be used.
	runtime.GOMAXPROCS(2) // Using 2 cores for parallelism.

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		printNumbers()
	}()

	go func() {
		defer wg.Done()
		printLetters()
	}()

	wg.Wait()

	fmt.Println("\nMain function is done.")
}

// ******///////parallelism//////

// ***** Channels
// syntax
// var channel_name = chan int
// channel_name := make(chan type)

// to send
// chan1 <- chan2
// to recieve
// var := <- chan1

// ******///////  Channels  //////
