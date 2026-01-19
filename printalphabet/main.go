package main

import "github.com/01-edu/z01"

func main() {
	for i := 'A'; i <= 'Z'; i++ {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
