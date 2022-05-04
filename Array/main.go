package main

import "fmt"

func main() {
	// food := [3]string{"🍔", "🍕", "🌭"}
	// fmt.Println(food)

	// Slices
	/*set := [7]string{"🦁", "🐎", "🐄", "🦋", "🦜", "🛫", "🚁"}
	animals := set[:5]
	fly := set[3:]
	fly[0] = "🦅"
	fmt.Println("Array", set)
	fmt.Println("Animales", animals)
	fmt.Println("Voladores", fly)
	fmt.Println("All", set[:])*/

	/*food := [5]string{"🌭", "🍓", "🍋", "🍔", "🍕"}
	fruits := food[1:3]
	fruits = append(fruits, "🍍", "🍈", "🍐")

	// Array[4]{"🍓", "🍋", "🍔", "🍕"}
	// new Array[8]{"🍓", "🍋","🍍", "🍈", "🍐"}

	fmt.Println("food", food)
	fmt.Println("fruits", fruits)
	fmt.Println("tamaño", len(fruits))
	fmt.Println("capacidad", cap(fruits))*/

	// fruits := []string{"🍓", "🍋"}
	/*fruits := make([]string, 0, 3)
	fruits = append(fruits, "🍓", "🍋", "🍍", "🍎")
	fmt.Println("fruits", fruits)
	fmt.Println("tamaño", len(fruits))
	fmt.Println("capacidad", cap(fruits))*/

	fruits := []string{}
	fmt.Println("fruits", fruits)
	fmt.Println("tamaño", len(fruits))
	fmt.Println("capacidad", cap(fruits))
	fmt.Println(fruits == nil)
}
