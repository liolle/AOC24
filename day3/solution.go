package main

import (
	"aoc24/mult"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func part1(buffer []byte)int {
  sum := 0
  lines := strings.Split(string(buffer),"\n") 
  for _,line := range(lines){
    patter := `mul\(([0-9]{1,3}),([0-9]{1,3})\)` 
    re := regexp.MustCompile(patter)
    matches := re.FindAllStringSubmatch(line,-1)

    for _,match := range(matches){
      if(len(matches)<3){continue}
      left,_ := strconv.Atoi(match[1]) 
      right,_ := strconv.Atoi(match[2])
      m := mult.Mult{Left:left,Right:right}
      sum +=  m.Mult()
    }
  }
  return sum
}

func getDos (buffer []byte)[]mult.Mult{
  input := []mult.Mult{}
  lines := []string{} 

  vals := strings.Split(string(buffer),"don't()")
  for idx,part := range(vals){

    if idx == 0 {
      lines = append(lines,part)
      continue
    }

    for ix,p := range(strings.Split(part,"do()")){

      if ix == 0 {
        continue
      }
      lines = append(lines,p)
    }
  }

  for _,line := range(lines){
    patter := `mul\(([0-9]{1,3}),([0-9]{1,3})\)` 
    re := regexp.MustCompile(patter)
    matches := re.FindAllStringSubmatch(line,-1)

    for _,match := range(matches){

      if(len(match)<3){continue}
      left,_ := strconv.Atoi(match[1]) 
      right,_ := strconv.Atoi(match[2])
      m := mult.Mult{Left:left,Right:right}
      input = append(input,m)
    }
  }

  return input
}

func part2(buffer []byte)int {
  input := getDos(buffer)
  sum := 0
  for _,mult := range(input){
    sum += mult.Mult() 
  }
  return sum
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

  fmt.Println(part1(buffer))
  fmt.Println(part2(buffer))
}
