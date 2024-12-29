package main

import (
	"fmt"
	"os"
)

func part1(secrets []Secret)(sum int64){
  
  for i := range secrets {
    secrets[i].Search(2000)
    sum += secrets[i].Current
  }
  return sum 
}
func part2(secrets []Secret,limit int)(sequence string,val int) {
  cache := map[string]int{}
  val = 0
  for i := range secrets {
    secrets[i].Scan(cache,limit)
  }

  for key := range cache {
    if cache[key] >val {
      val = cache[key]
      sequence = key
    }
  }

  return sequence,val 
}

func main()  {
  // Open the file
  filePath := "input.txt"
  //filePath := "sample.txt"
  //filePath := "sample2.txt"
  buffer, err := os.ReadFile(filePath)
  if err != nil {
    fmt.Println("Error opening file:", err)
    return
  }

  s1 := ExtractIntput(string(buffer))
  s2 := ExtractIntput(string(buffer))
  fmt.Println(part1(s1))
  fmt.Println("-----")
  fmt.Println(part2(s2,2001))
  
}
