package main

import (
  "fmt"
  "os"
)

func part1()int {
}
func part2()int {
}

func main()  {
  // Open the file
  filePath := "input.txt"
  //filePath := "sample.txt"
  buffer, err := os.ReadFile(filePath)
  if err != nil {
    fmt.Println("Error opening file:", err)
    return
  }

}
