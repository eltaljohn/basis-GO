package main

import "fmt"

func main() {
	type course struct {
		Name      string
		Professor string
		Country   string
	}

	db := course{Name: "Bases de datos", Professor: "Alexys", Country: "Colombia"}
	git := course{"Git", "Beto", "Bolivia"}
	css := course{Professor: "Alvaro"}

	fmt.Printf("%+v\n", db.Name)
	fmt.Printf("%+v\n", git.Country)
	fmt.Printf("%+v\n", css.Professor)

	p := &db
	p.Professor = "Alvaro"
	fmt.Printf("%+v\n", db)
	fmt.Printf("%+v\n", p)
}
