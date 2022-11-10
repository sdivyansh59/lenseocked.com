package main

import (
	"os"
	"text/template"
)

type User struct {
	Name string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := User {
		Name : "John Smith",
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

	data.Name = "Jon Calhoun"
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

}