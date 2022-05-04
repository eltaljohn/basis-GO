package main

import (
	"fmt"
	"strings"
)

func main() {
	//hello("John", "Bolanos")
	//emoji := "🐶"
	//chance(&emoji)
	//fmt.Println(emoji)

	//total := sum(2, 3)
	//fmt.Println(total)
	minusc, mayusc := Convert("BolAnoS")
	fmt.Println(minusc, mayusc)
}

func Convert(text string) (string, string) {
	min := strings.ToLower(text)
	may := strings.ToUpper(text)
	return min, may
}

//func sum(num1, num2 int) int {
//	return num1 + num2
//}

//func chance(value *string) {
//	*value = "🙂"
//}

//func hello(firstName string, lastName string) {
//	fmt.Printf("Hello %s %s", firstName, lastName)
//}
