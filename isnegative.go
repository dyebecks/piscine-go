package main

import piscine "github.com/01-edu/z01"

func IsNegative(nb int) {
	if nb < 0 {
		z01.printRune("T")
	} else {
		z01.printRune("F")
	}
	z01.PrintRune('\n')
}
