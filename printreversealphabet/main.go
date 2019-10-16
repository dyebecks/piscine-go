package main

import "github.com/01-edu/z01"
import "fmt"

func main() {
	for i := 122; i >= 97; {
		z01.PrintRune(i)
		i--
	}
	fmt.Println()
}