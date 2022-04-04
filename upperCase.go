package main 

import (
	"fmt"
	"strings"
)

func UPPERCASE() {
	country := "Nigeria"
	fmt.Println(strings.ToUpper(country))
}

func lowercase() {
	state := "Lagos"
	fmt.Println(strings.ToLower(state))
}


func main() {
UPPERCASE()
lowercase()

}