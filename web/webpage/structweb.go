package main
import (
	"text/template"
	"log"
	"os"
)

type Person struct {
	type Person struct {
	Name string
	Hobbies [] string
}

const data = "Person Name is {{.Name}}.{{range .Hobbies}} He likes {{.}} {{end}}"

func main() {
	person := Person {
	Name : "Shahzaib",
	Hobbies: []string {"Playing Cricket ", "programming"},	
	}
	
	template1 := template.New("test")
	template1, err := template1.Parse(data)
	if err != nil {
		log.Fatal(err)
		return
	}

	template1.Execute(os.Stdout, person)
}