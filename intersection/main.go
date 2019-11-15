package main

import (
    "fmt"
    "github.com/juliangruber/go-intersect"
)

func main() {
  a := []int{1, 2, 3}
  b := []int{2, 3, 4}
  fmt.Println(intersect.Simple(a, b))
}

