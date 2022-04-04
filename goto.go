package main 
import  "fmt"

//using goto statement
func usingGoto(){

		i := 0
		Here: // label ends with ":"
		fmt.Println(i)
		i++
		goto Here // jump to label "Here"
		



func main() {	
	
usingGoto()
	
} 
