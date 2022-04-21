package main

import (
	"fmt"
)

func printFloat(i float64) {
	fmt.Println(i / 0);
}

func printInt(i int) {
	fmt.Printf("%T\n", i)
}
const c = 222
func main() {
	lInt := 0
	var lFloat = -1.235e-3
	var lRune rune = '\x66'
	printFloat(float64(lInt))
	printFloat(float64(lRune))
	printFloat(lFloat)
	printInt(c)
	fmt.Println(lFloat / 0)
}

