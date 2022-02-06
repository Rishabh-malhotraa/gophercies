package main

import "fmt"

type Profile struct {
	Name  string
	Major string
	Language
}

type Language struct {
	Compiled    string
	Interpreted string
}

func main() {
	var x = 1

	a := Profile{"Rishabh", "COMPSCI", Language{"C++", "JAVASCRIPT"}}
	fmt.Println(a.Compiled, a.Name, x)
}
