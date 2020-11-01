package main

import "fmt"

func main() {
array := []string{"Foo", "Bar", "Test"}
  for index, value := range array {
    fmt.Println("index =", index, " | value =", value)
  }
}
