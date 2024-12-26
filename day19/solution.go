package main

import (
  "fmt"
  "os"
)

func part1(l Layout)int {
  valid_cnt := 0
  
  for i := range l.Desighs {
    initial := Node{
      Prev:nil,
      Desigh: l.Desighs[i],
      Score:0,
    }
    found := l.Solve(initial)
    fmt.Println(l.Desighs[i],fmt.Sprintf("%d / %d",i+1,len(l.Desighs)),found>0)
    valid_cnt += min(1,found)
  }
  
  return valid_cnt
}

func part2(l Layout)int {
  valid_cnt := 0
  
  for i := range l.Desighs {
    initial := Node{
      Prev:nil,
      Desigh: l.Desighs[i],
      Score:0,
    }
    found := l.Solve(initial)
    fmt.Println(l.Desighs[i],found)
    valid_cnt += found
  }

  return valid_cnt
}

func main()  {
  // Open the file
  filePath := "input.txt"
  //filePath := "sample.txt"
  //filePath := "edge.txt"
  buffer, err := os.ReadFile(filePath)
  if err != nil {
    fmt.Println("Error opening file:", err)
    return
  }

  l1 := ExtractInput(string(buffer))

  fmt.Println(part1(l1))
  fmt.Println(part2(l1))
}
