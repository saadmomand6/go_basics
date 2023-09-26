package custom_package

import "fmt"
//function or variables declared with small letters can't be accessed outside the package

func PrintValues(s string){
	fmt.Println(s)
}
func printVal(v string){
	fmt.Println(v)
}
var Val int
var val int