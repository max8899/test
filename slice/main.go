package main

import "fmt"

func main() {
    fmt.Println("Hello, world")
    slice := []int{1, 2, 3, 4, 5}

    fmt.Println("4:5", slice[4:5])
    fmt.Println("4:4", slice[4:4])
    fmt.Println(slice[0:1])
    fmt.Println(slice[3:])
}
