package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}

func ZipString(str string) string {
	if len(str) == 0 {
		return ""
	}
	result := ""
	count := 1
	for i := 0; i < len(str); i++ {
		if i+1 < len(str) && str[i] == str[i+1] {
			count++
		} else {
			result += strconv.Itoa(count) + string(str[i])
			count = 1
		}
	}
	return result
}
