package main

import (
	"errors"
	"fmt"
)

func main() {
	/*content, err := ioutil.ReadFile("./hello.txt")
	if err != nil {
		fmt.Printf("Ocurrió un error: %v", err)
		return
	}

	fmt.Println(string(content))*/
	result, err := division(10, 0)
	if err != nil {
		fmt.Printf("Ocurrió un error: %v", err)
		return
	}
	fmt.Println(result)
}

func division(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("no puedes dividir por cero")
	}
	return dividendo / divisor, nil
}
