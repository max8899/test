package main

import "fmt"


func main() {
    fmt.Println("Hello, world")

    var a, b float64

    a = float64(3)

    b = float64(100)

    c := float64(a/b)

    fmt.Println(fmt.Sprintf("%.100f", c))

}
