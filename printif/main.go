package main

import (
	"fmt"
)

func main() {
	fmt.Print(PrintIf("abcdefz"))
	fmt.Print(PrintIf("abc"))
	fmt.Print(PrintIf(""))
	fmt.Print(PrintIf("10"))
}

func PrintIf(str string) string {
	if str == "" || len(str) >= 3 {
		return "G\n"
	}
	return "Invalid Input\n"
}
