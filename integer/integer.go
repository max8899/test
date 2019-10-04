package main

import (
    "fmt"
    "reflect"
)


func main() {
    fmt.Println("Hello, world")
    fmt.Println("0", reflect.TypeOf(0).Name())
    fmt.Println("2 ** 8 ", reflect.TypeOf(256).Name())
    fmt.Println("2 ** 32 ", reflect.TypeOf(4294967296).Name())
    fmt.Println("2 ** 64 ", reflect.TypeOf(1844674407370955161).Name())
}
