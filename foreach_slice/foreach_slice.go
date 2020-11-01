package main

import "fmt"

func main() {
a := []string{"Foo", "Bar", "Test"}
  for i, s := range a {
    fmt.Println("i =", i, " | s =", s)
  }
}
