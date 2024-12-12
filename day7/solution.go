package main

import (
	bridge "aoc24/Bridge"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DFS(b *bridge.Bridge,idx int,sum int,operators string,third_op bool)  {


  last := string(operators[len(operators)-1])

  if last == "+" {sum+=b.Values[idx]}
  if last == "*" {sum*=b.Values[idx]}
  if last == "|" {
    left := strconv.Itoa(sum)
    if sum == 0 {left = ""}
    right := strconv.Itoa(b.Values[idx]) 
    sum,_ = strconv.Atoi(left +right)
  }


  if sum == b.Target && idx == len(b.Values)-1 {
    op := operators[1:]
    if !Includes(b.Solutions, op){
      b.Solutions = append(b.Solutions,op)
    }
    return
  }

  if idx >= len(b.Values) || sum > b.Target {
    return
  }

  if idx < len(b.Values)-1{

    DFS(b,idx+1,sum,operators + "*",third_op)
    DFS(b,idx+1,sum,operators + "+",third_op)
    DFS(b,idx+1,sum,operators + "|",third_op)

  }

}

func ExtractInput(buffer string) []bridge.Bridge{

  bridges := []bridge.Bridge{}
  lines := strings.Split(buffer,"\n")

  for _,line := range(lines[:len(lines)-1]) {

    parts := strings.Split(line,": ")
    values := []int{}
    left,_ := strconv.Atoi(parts[0])

    for _,v := range(strings.Split(parts[1]," ")){
      value,_ := strconv.Atoi(v)
      values = append(values,value)
    }

    bridges = append(bridges,bridge.Bridge{
      Target:left,
      Values:values,
      Solutions: []string{},
    })

  }

  return bridges
}

func Includes(slice []string, value string) bool {
  for _,v := range(slice) {
    if v == value {
      return true
    }
  }
  return false
}

func part1(bridges []bridge.Bridge) int{
  sum := 0

  for _,b := range(bridges){
    DFS(&b,0,1,"*",false)
    DFS(&b,0,0,"+",false)
    fmt.Println(b.Solutions)
    if len(b.Solutions)>0{sum+=b.Target}
  }

  return sum
}

func part2(bridges []bridge.Bridge) int{
  sum := 0

  for _,b := range(bridges){
    DFS(&b,0,1,"*",true)
    DFS(&b,0,0,"+",true)
    DFS(&b,0,0,"|",true)

    fmt.Println(b.Solutions)
    if len(b.Solutions)>0{sum+=b.Target}
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

  bridges := ExtractInput(string(buffer))
  fmt.Println(part1(bridges))
  fmt.Println(part2(bridges))
}
