package main

import "github.com/01-edu/z01"

func main() {
	for i := 48; i <= 57; {
		z01.PrintRune(rune(i))
		i++
	}
	z01.PrintRune('\n')
}
