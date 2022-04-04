package main
import (
      "fmt" 
      "strings" 
)

func prefix() {
  
//THIS WILL DETECT THE FIRST ALPHABET AND RETURN TRUE
  str := "AMERICA"
  fmt.Println("Value returned for prefix is :", strings.HasPrefix(str, "AM"))

}

func suffix() {
  
  //THIS WILL DETECT THE LAST ALPHABET AND RETURN TRUE
    str := "CANADA"
    fmt.Println("Value returned for suffix is: ", strings.HasSuffix(str, "DA"))
  
  }
  



func main(){
  prefix()
  suffix()
}