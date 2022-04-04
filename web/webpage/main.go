package main
import (
	"text/template"
	"log"
	"os"
)

func main() {
	template, err := template.ParseFiles("text.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	template.Execute(os.Stdout, "Error Occured")
}
