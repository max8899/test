package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("9" >= fmt.Sprintf("%d", time.Now().Unix()))

}
