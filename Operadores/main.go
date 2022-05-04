package main

import "fmt"

func main() {
	// Operadores aritméticos: (), *, /, %, +, -
	var a int8 = 3 + 4 + 2*3
	fmt.Println(a)
	// Operadores de asignación: =, +=, -=, *=. /=, %=
	var b = 10
	b -= 2
	fmt.Println(b)
	// declaración post-incremento y pre-decremento: ++ , --
	// (No son una expresión sino una declaración)
	var c = 3
	c--
	fmt.Println(c)
	// Operadores de comparación: > ,< ,== ,!= ,>= ,<=
	fmt.Println(4 != 4)
	// Operadores lógicos &&, ||
	var age = 14
	fmt.Println("adulto: ", age >= 18 && age <= 60)
	fmt.Println("niño o viejo: ", age < 18 || age > 60)
	// Unario
	fmt.Println(!(4 == 4))
}
