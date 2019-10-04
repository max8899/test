package main

import "fmt"


func main() {
    fmt.Println("Hello, world")

    var a, b float32

    a = float32(1)

    b = float32(3)

    c := float32(a/b)

    fmt.Println(fmt.Sprintf("%v", c))

}
