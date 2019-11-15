package main

import (
	"fmt"
	"time"
)

type Flag int64

func (f Flag) String() string {
	return fmt.Sprintf("0x%b", f)
}

func main() {
	fmt.Println("ok")
	fmt.Println(fmt.Sprintf("0x%b", 0xb))
	fmt.Println(Flag(1))
	time.Now()
}
