package main

import (
	"fmt"
	"os"
	"slices"
)

func part1(board Board)(out int) {
  board.Build(board.Plan)
  board.TurnOn()
  out = board.Count()
  return out
}

func part2(board Board)(out string) {
  cnt := 0
  total := 0
  affected_outputs := []string{} 
  seen := map[string]bool{}
  for key := range board.Plan {
    valid,blame,_ := board.Valide(key)
    total ++
    if !valid {
      cnt++
      if _,has := seen[blame]; !has{
        affected_outputs = append(affected_outputs,blame)
        seen[blame] = true
      }  
    }
  } 

  slices.Sort((affected_outputs))
  fmt.Printf("[Invalid Gates]: %d/%d\n",cnt,total)

  for i := range affected_outputs {
    if i == 0 {
      out +=affected_outputs[i] 
      continue
    }
    out += fmt.Sprintf(",%s",affected_outputs[i])
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
  b2 := ExtractInput(string(buffer))
  fmt.Println(part1(b1))
  fmt.Println("-----------------------")
  fmt.Println(part2(b2))
}
