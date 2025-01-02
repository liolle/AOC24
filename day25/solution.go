package main

import (
  "fmt"
  "os"
)

func part1(bag Bag)(out int) {
  _,keys := bag.CountValidKey() 
  seen := map[string]bool{}
  for i := range keys {
    if _,has := seen[keys[i]];!has {
      out ++
      seen[keys[i]] = true
    }
  }
  return out  
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

  b1 := ExtractInput(string(buffer))
  fmt.Println(part1(b1))
}
