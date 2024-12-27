package main

import (
  "fmt"
  "os"
)

func part1(r Race,cheat_duration int,min_time_save int )int {
  r.BFS()
  cnt := 0
  for key := range r.Path{
    node := r.Path[key]
    gain :=r.Cheat(node,cheat_duration)
    for key := range gain {
      val := gain[key]
      if val>=min_time_save{
        cnt++
      }
    }
  }
  return cnt
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

  r1 := ExtractInput(string(buffer))
  fmt.Println(part1(r1,2,100))
  fmt.Println(part1(r1,20,100))
}
