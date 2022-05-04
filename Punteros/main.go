package main

import "fmt"

func main() {
	fruit := "🍎"
	var p *string
	p = &fruit
	*p = "🍍"
	fmt.Printf("Tipo %T, Valor:%s, Direccion: %v \n", fruit, fruit, &fruit)
	fmt.Printf("Tipo %p, valor %v, Desreferenciación: %s \n", p, p, *p)
}
