package main

import "fmt"

func main() {
	// food := [3]string{"π", "π", "π­"}
	// fmt.Println(food)

	// Slices
	/*set := [7]string{"π¦", "π", "π", "π¦", "π¦", "π«", "π"}
	animals := set[:5]
	fly := set[3:]
	fly[0] = "π¦"
	fmt.Println("Array", set)
	fmt.Println("Animales", animals)
	fmt.Println("Voladores", fly)
	fmt.Println("All", set[:])*/

	/*food := [5]string{"π­", "π", "π", "π", "π"}
	fruits := food[1:3]
	fruits = append(fruits, "π", "π", "π")

	// Array[4]{"π", "π", "π", "π"}
	// new Array[8]{"π", "π","π", "π", "π"}

	fmt.Println("food", food)
	fmt.Println("fruits", fruits)
	fmt.Println("tamaΓ±o", len(fruits))
	fmt.Println("capacidad", cap(fruits))*/

	// fruits := []string{"π", "π"}
	/*fruits := make([]string, 0, 3)
	fruits = append(fruits, "π", "π", "π", "π")
	fmt.Println("fruits", fruits)
	fmt.Println("tamaΓ±o", len(fruits))
	fmt.Println("capacidad", cap(fruits))*/

	fruits := []string{}
	fmt.Println("fruits", fruits)
	fmt.Println("tamaΓ±o", len(fruits))
	fmt.Println("capacidad", cap(fruits))
	fmt.Println(fruits == nil)
}
