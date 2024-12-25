package main

import (
  "fmt"
  "math"
  "os"
)

func part1(c Computer)[]int {
  c.Compute()
  return  c.OutputBuffer
}

func part2 (c Computer)int{
  start := int(math.Pow(8,15))

  B := 0 
  C := 0 
  idx := 0 
  for i:=0;idx<16 && i<1000000;i++{
    start+=1<<int(math.Max(0,38-float64(idx*3))) 
    A := start
    c.reset(B,C)
    c.A = A    
    c.Compute()
    ix := int(math.Max(0,float64(len(c.Target)-1-idx ))) 

    if c.Target[ix] == c.OutputBuffer[ix]{
      idx++
    } 
  }
  fmt.Println(c.OutputBuffer,start)

  return start
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
  c1 := ExtractInput(string(buffer))
  fmt.Println(part1(c1))
  fmt.Println(part2(c1))
}
