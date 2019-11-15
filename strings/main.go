package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.EqualFold("Abc", "ABC"))
	ss := strings.Split("abc", ".")
	fmt.Println(ss[len(ss)-1])
}
