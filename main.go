package main

import (
	//"errors"
	//"encoding/json"
	//"fmt"
	// "sync"
	// "time"
	//"strings"
	//"runtime"
	//"sort"
	// custom_package "gomodulename/custom"
	//"encoding/json"
	// "encoding/json"
	// "io"
	"gomodulename/database"
	"gomodulename/models"
	"net/http"

	"github.com/gin-gonic/gin"
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
// 	// Set the maximum number of CPU cores to be used.
// 	//runtime.GOMAXPROCS(2) // Using 2 cores for parallelism.

// 	var wg sync.WaitGroup
// 	wg.Add(2) //here 2 represent number of goroutines we are having

// 	go func() {
// 		defer wg.Done()
// 		printNumbers()
// 	}()

// 	go func() {
// 		defer wg.Done()
// 		printLetters()
// 	}()

// 	wg.Wait()

// 	fmt.Println("\nMain function is done.")
// }

// ******///////parallelism//////

// ***** Channels
// syntax
// var channel_name = chan int
// channel_name := make(chan type)

// to send
// chan1 <- chan2
// to recieve
// var := <- chan1
// Close a channel: close(channel).
// After closing, no value will be sent to the channel.

// EXAMPLE 1
// Prints a greeting message using values received in
// the channel
// func greet(c chan string) {

// 	name := <- c	// receiving value from channel
// 	fmt.Println("Hello", name)
// }

// func main() {

// 	// Making a channel of value type string
// 	text := make(chan string)

// 	// Starting a concurrent goroutine
// 	go greet(text)

// 	// Sending values to the channel c
// 	text <- "World"

// 	// Closing channel
// 	close(text)
// }

// EXAMPLE 1 ENDS

// EXAMPLE 2

// func main() {
// 	n := 3

// 	// This is where we "make" the channel, which can be used
// 	// to move the `int` datatype
// 	cname := make(chan int)

// 	// We still run this function as a goroutine, but this time,
// 	// the channel that we made is also provided
// 	go multiplyByTwo(n, cname)

// 	// Once any output is received on this channel, print it to the console and proceed
// 	fmt.Println(<-cname)
// }

// // This function now accepts a channel as its second argument...
// // here (out chan<- int) will give the value to chan
// func multiplyByTwo(num int, out chan<- int) {
// 	result := num * 2

// 	//... and pipes the result into it
// 	out <- result
// }

// EXAMPLE 2 ENDS
// ******///////  Channels  //////

// ***** Select
// if multiple cases are ready to proceed, then one of them can be selected randomly.
// The default statement in the select statement is used to protect select
// statement from blocking. This statement executes when there is no case statement is ready to proceed.
// func main() {
// 	chan1 := make(chan string)
// 	chan2 := make(chan int)

// 	go fun1(chan1)
// 	go fun2(chan2)

// 	select {
// 	case val1 := <-chan1:
// 		fmt.Println(val1)
// 	case val2 := <-chan2:
// 		fmt.Println(val2)
// 	}

// }
// func fun1(chanel1 chan string) {
// 	chanel1 <- "chnel1 executed"

// }
// func fun2(chanel2 chan int) {
// 	chanel2 <- 22
// }

// ******///////  Select  //////

// ***** Mutex
// '********** EXample 1
// type  safecounter struct{
// 	mu sync.Mutex
// 	val int
// }
// func (c *safecounter) Increment(){
// 		c.mu.Lock()
// 		defer c.mu.Unlock()
// 		c.val++
// }
// func (c *safecounter) GetValue() int {
// 		c.mu.Lock()
// 		defer c.mu.Unlock()
// 		return c.val
// 	}
// func main(){
// 	counter := safecounter{}
// 	for i := 0; i < 10; i++ {
// 		 go func() {
// 			 counter.Increment()
// 		}()
// 	}
// 	time.Sleep(1 * time.Second)
// 	fmt.Printf("Final counter value: %d\n", counter.GetValue())
// }
// '********** EXample 1 ends

// '********** EXample 2

// type Counter struct {
// 	mu sync.Mutex
// 	value int
// }

// func (c *Counter) Increment() {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	c.value++
// }

// func (c *Counter) GetValue() int {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	return c.value
// }

// func main() {
// 	counter := &Counter{}

// 	for i := 0; i < 10; i++ {
// 		go func() {
// 			counter.Increment()
// 		}()
// 	}

// 	time.Sleep(1 * time.Second) // Wait for goroutines to finish

// 	fmt.Printf("Final counter value: %d\n", counter.GetValue())
// }

// '********** EXample 2 ends

// ******///////  Mutex  //////

// ****** Interface
// syntax:
// type interface_name Interface{
// }
// Example 1
// type my_struct struct {
// 	first  int
// 	second int
// }

// type my_interface interface {
// 	//these are signatures which should must be implemented
// 	add(a, b int) int
// 	multi(a, b int) int
// }

// func (met my_struct) add(a, b int) int {
// 	return a + b + met.first + met.second
// }
// func (met my_struct) multi(a, b int) int {
// 	return a * b * met.first * met.second
// }
// func main() {
// 	var t my_interface
// 	t = my_struct{1, 2} //this 1,2 will be store in met.first and met.second
// 	fmt.Println(t.add(1, 2))             //this 1,2 will be store in a ,b
// 	fmt.Println(t.multi(1, 2))           //this 1,2 will be store in a ,b
// }

// Example 1 ends
// ******///////  Interface  //////

// ****** MAPS

// syntax
//var map_name map[key_type]value_type
//ex: "apples" : 1
//we can direct assign values:
// mp := make(map[string]int{
// 	"peach" : 2,
// 	"grapes" : 3,
// 	"apples" : 1,
// 	"bananas" : 2,
// })
// empty map
// mp := make(map[string]int)
// func main(){
// 	var map_name map[string]int = map[string]int{
// 		"peach" : 2,
// 		"grapes" : 3,
// 		"apples" : 1,
// 		"bananas" : 2,
// 	}
// fmt.Println(map_name)
// fmt.Println(map_name["apples"])
//to change value
// map_name["apples"] = 2
// fmt.Println(map_name)
// //to add new value
// map_name["oranges"] = 2
// fmt.Println(map_name)
// //to delete value
// delete(map_name, "peach")
// fmt.Println(map_name)
//to check does value exit or not and if doesnot exit then add that value
//val , ok := map_name["peach"]
// fmt.Println(val, ok)
// fmt.Println(len(map_name))
//for loop in map
// 	for k, v := range map_name {
// 		fmt.Println(k,v)
// 	}
// }

// ******///////  MAPS  //////

// ****** JSON Manupulating
// Example 1
// func main(){
// 	var info map[string]interface{}
// 	data := `{"name":"saad", "age":22, "number":234234234}`
// 	err := json.Unmarshal([]byte(data), &info)
// 	if err != nil{
// 		fmt.Println(err)
// 	}
// 	fmt.Println(info)
//  	fmt.Println(info["name"])
//  }
// Example 1 ends

// Example 2
// type Mystruct struct {
// 	Name string
// 	Age int
// 	Number int
// }
// func main(){
// 	var wrapdata Mystruct
// 	data := `{"name":"saad", "age":22, "number":234234234}`
// 	err := json.Unmarshal([]byte(data), &wrapdata)
// 	if err != nil{
// 		fmt.Println(err)
// 	}
// 	fmt.Println(wrapdata.Age)
// }
// Example 2 ends
// ******///////  JSON Manupulating  //////

// ******///////  CUSTOM PACKAGE  //////
// firstly you will need a mod file
// 1) go mod init .
// 2) go mod init gomodulename

// func main(){
// 	custom_package.PrintValues("kjnkbk")
// 	custom_package.Val = 100
// 	fmt.Println(custom_package.Val)
// }
// ******///////  CUSTOM PACKAGE   //////

// ******///////  GIN FRAMWORK //////
// func main(){
// 	router := gin.Default()
// 	// the get function requires a path and a function
// 	router.GET("/firstApi", func (c *gin.Context){
// 		c.JSON(http.StatusOK, gin.H{
// 			"data": "i am your first framework",
// 		  })
// 		})
// 	router.Run()
// }
// ******///////  GIN FRAMWORK //////

// ******///////  API understanding //////
type api struct{
	Name string `json:"name"`
	Email string `json: "email"`
}

 var data api

func main(){
	router := gin.Default()
	router.GET("/getapi", getcontroller)
	router.POST("/postapi", postcontroller)
	router.PUT("/putapi", putcontroller)
	router.DELETE("/deleteapi", deletecontroller)
	database.InitDB()
	models.AutoMigrateTableIfNotExist()
	router.Run()
}
func getcontroller(c *gin.Context){
c.JSON(http.StatusOK, gin.H{
	"data": data,
})
}
func postcontroller(c *gin.Context){
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "something went wrong",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
		"message": "posted sucessfully",
	})

}
func putcontroller(c *gin.Context){
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "something went wrong",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
		"message": "updated sucessfully",
	})
}
func deletecontroller(c *gin.Context){
	data = api{}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
		"message": "deleted sucessfully",
	})
}
// ******///////  API understanding //////

// ******/////// THIRD PARTY API understanding //////
// func main() {
// 	router := gin.Default()
// 	router.GET("/getapi", getcontroller)
// 	router.Run()
// }

// var url = "http://date.jsontest.com/"

// func getcontroller(c *gin.Context) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"message": err,
// 		})
// 	}
// 	//close response body
// 	defer resp.Body.Close()
// 	//read all response and convert into array of bytes
// 	data, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"message": err,
// 		})
// 	}
// 	var target map[string]interface{}
// 	// unmarshal byte data ko read krke target map me bind krdega
// 	err = json.Unmarshal(data, &target)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"message": err,
// 		})
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"data": target,
// 	})
// }

// ******/////// THIRD PARTY API understanding //////
