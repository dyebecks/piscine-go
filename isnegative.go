package main

import piscine "github.com/01-edu/z01"

func IsNegative(nb int) {
	if nb < 0 {
		z01.printLn("T")
	} else {
		z01.printLn("F")
	}
	z01.PrintRune('\n')
}
