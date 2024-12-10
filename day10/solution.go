package main

import (
	topography "aoc24/Topography"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func Extract(buffer string)topography.Topography{

  top := [][]int{}
  lines := strings.Split(buffer, "\n")

 for _,line := range(lines[:len(lines)-1]){
    Row := []int{}
    
    for _,r := range(line){
      num,_ := strconv.Atoi(string(r))
      Row = append(Row,num)
    }

    top = append(top,Row)
  }
  return topography.Topography{
    Map: top,
    Rows: len(top),
    Cols: len(top[0]),
    Solution: []topography.Trailhead{},
  }

}

func part1(t topography.Topography)int {
  sum := 0
  t.Scan()

  for _,elem := range(t.Solution) {sum += elem.Score }

  return sum
}

func part2(t topography.Topography)int {
  sum := 0
  t.Rate()

  for _,elem := range(t.Solution) {sum += elem.Score}

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

  t1 := Extract(string(buffer))
  t2 := Extract(string(buffer))

  t1.Print()

  fmt.Println(part1(t1))
  fmt.Println(part2(t2))
}
