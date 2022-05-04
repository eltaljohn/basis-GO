package main

import "fmt"

func main() {
	//emoji := "😸"
	//if emoji == "🌵" {
	//	fmt.Println("Es un cactus")
	//} else if emoji == "🙂" {
	//	fmt.Println("Es una carita")
	//} else {
	//	fmt.Printf("El emoji es %s", emoji)
	//}

	/*if emoji := "🌵"; emoji == "🌵" {
		fmt.Println("Es un cactus")
	} else if emoji == "🙂" {
		fmt.Println("Es un cactus")
	} else {
		fmt.Printf("El emoji es %s", emoji)
	}*/
	// fmt.Println(emoji)
	emoji := "🍎"
	/*switch emoji {
	case "🐈", "🐕‍":
		fmt.Println("Es un animal")
	case "🍌", "🍎":
		fmt.Println("es una fruta")
	default:
		fmt.Println("no es un animal ni una fruta")
	}*/
	switch {
	case emoji == "🐈" || emoji == "🐕‍":
		fmt.Println("Es un animal")
	case emoji == "🍌" || emoji == "🍎":
		fmt.Println("es una fruta")
	default:
		fmt.Println("no es un animal ni una fruta")
	}
}
