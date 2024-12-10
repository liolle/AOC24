package main

import (
  "fmt"
  "os"
)

func part1()int {
  return 0
}
func part2()int {
  return 0
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
