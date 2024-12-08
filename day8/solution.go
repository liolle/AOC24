package main

import (
	roof "aoc24/Roof"
	"fmt"
	"os"
	"strings"
)

func ExtractInput (buffer string) roof.Roof{
  
  frequencys := map[string][]roof.Position{}
  antennas := []roof.Position{}

  rows := 0
  cols := 0

  lines := strings.Split(buffer,"\n")
  rows = len(lines)-1

  for y,line := range(lines[:len(lines)-1]){

    chars := strings.Split(line,"")
    cols = len(chars)

    for x,char := range(chars){
      if char == "." {continue}
      antennas = append(antennas,roof.Position{X: x,Y: y})
      if _,exist := frequencys[char]; exist {
        frequencys[char] = append(frequencys[char],roof.Position{X: x,Y: y})
      }else {
        frequencys[char] = []roof.Position{{X: x,Y: y}}
      }
    }
  }

  return roof.Roof{
    Frequencys: frequencys,
    Antennas: antennas,
    Rows: rows,
    Cols: cols,
    Antinodes: []roof.Position{},
  }
}

func part1(roof roof.Roof)int {
  roof.ComputeAntinodes()
  fmt.Println(roof.Antinodes)
  roof.Print()
  return len(roof.Antinodes)
}

func part2(roof roof.Roof)int {
  roof.ComputeAntinodesHamonic()
  fmt.Println(roof.Antinodes)
  roof.Print()
  return len(roof.Antinodes) + len(roof.Antennas)
}

func main()  {
  // Open the file
  filePath := "input.txt"
  //filePath := "sample.txt"
  //filePath := "edge1.txt"
  buffer, err := os.ReadFile(filePath)
  if err != nil {
    fmt.Println("Error opening file:", err)
    return
  }

  r1 := ExtractInput(string(buffer))
  r2 := ExtractInput(string(buffer))

  fmt.Println(part1(r1))
  fmt.Println(part2(r2), " (Antinodes + Antennas)")
}
