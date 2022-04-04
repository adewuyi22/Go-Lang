package main
import (
	"fmt"
) 

func main () {
	fmt.Println("start")
	defer fmt.Println("This was deferred")
	panic("something bad happen")
	fmt.Println("end")
}
