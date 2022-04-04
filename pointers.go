package main 
import (
	"fmt"
)

//Declaring a pointer function 
func mypointer() {

	p := 5
	a := p
	fmt.Println(p, a) 
	
	p = 8
	fmt.Println(p, a) 
	
}

func main() {
mypointer()
	
} 