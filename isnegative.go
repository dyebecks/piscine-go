package main

import piscine ".."

func IsNegative(nb int) {
	if nb<0 {piscine.printLn("T")}
	else {piscine.printLn("F")}
}


func main() {
	piscine.IsNegative(1)
	piscine.IsNegative(0)
	piscine.IsNegative(-1)
	
}
