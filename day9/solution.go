package main

import (
	disk "aoc24/Disk"
	"fmt"
	"os"
	"strconv"
)

func ExtractInput(buffer string) disk.Disk{

  m := []disk.Layout{}

  for i := 0; i < len(buffer)-1; i += 2 {
    left := buffer[i]
    right := ""
    if i <len(buffer)-1 {right = string(buffer[i+1])}

    ln,_ := strconv.Atoi(string(left))
    rn,_ := strconv.Atoi(right)

    m = append(m,disk.Layout{Files: ln,FreeSpaces: rn})

  }

  return disk.Disk{
    Map: m,
    Compilation: []int{},
  }

}

func part1(disk disk.Disk)int {
  sum := 0
  disk.Compile()
  disk.Fill()
  disk.Print()
  for idx,f := range(disk.Compilation) {
    if f == -1 {break}
    sum += idx * f
  }
  return sum
}

func part2(disk disk.Disk)int {
  sum := 0
  disk.Compile()
  disk.FullFill()
  disk.Print()
  for idx,f := range(disk.Compilation) {
    if f == -1 {continue}
    sum += idx * f
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

  d1 := ExtractInput(string(buffer))
  d2 := ExtractInput(string(buffer))
  fmt.Println(part1(d1))
  fmt.Println(part2(d2))
}
