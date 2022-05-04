package main

import "fmt"

func main() {
	/*for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for {
		fmt.Println(i)
		i++
	}
	nums := []uint8{2, 4, 6, 8}
	for i, v := range nums {
		fmt.Println("Indice ", i, "valor ", v)
	}

	sports := map[string]string{"basketball": "🏀", "soccer": "⚽️"}
	for key, v := range sports {
		fmt.Println("key", key, "value", v)
	}*/
	hello := "hello"
	for _, v := range hello {
		fmt.Println(string(v))
	}
}
