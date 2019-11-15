package main

import "fmt"

type A struct {
    Name string
}

func main() {
   m := make(map[string]int)
   fmt.Println(m)
   m["a"] = 1
   m["b"] = 2

   m1 := make(map[string]A, 1)
   fmt.Println(fmt.Sprintf("%+v", m1))
   delete(m , "c")
   fmt.Println(m)
}

