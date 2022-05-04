package main

import (
	"fmt"
	"os"
)

func main() {
	/*defer fmt.Println(3)
	fmt.Println(1)
	fmt.Println(2)

	a := 5
	defer fmt.Println("Defer:", a)

	a = 10
	fmt.Println(a)
	*/
	file, err := os.Create("hello.txt")
	if err != nil {
		fmt.Printf("Ocurrió un error al crear el archivo %v", err)
		return
	}
	defer file.Close()

	_, err = file.Write([]byte("Hello EDteam"))
	if err != nil {
		fmt.Printf("Ocurrió un error al escribir el archivo %v", err)
		return
	}
}
