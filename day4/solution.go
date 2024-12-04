package main

import (
	"aoc24/Point"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Extract 2D array from input

func arrayFromBuffuer(buffer []byte)[][]string{
  arr := [][]string{}
  lines := strings.Split(string(buffer),"\n")

  for _,line := range(lines){
    if line == "" {continue}
    arr = append(arr,strings.Split(line,""))
  }
  return arr
}

func Includes(slice []Point.Point, value Point.Point) bool {
  for _,v := range(slice) {
    if v.Equals(value) {
      return true
    }
  }
  return false
}


var directions = []Point.Point{
  {Row: -1,Col: -1,Direction:"NW"},
  {Row: 0,Col: -1,Direction:"W"},
  {Row: 1,Col: -1,Direction:"SW"},
  {Row: 1,Col: 0,Direction:"S"},
  {Row: 1,Col: 1,Direction:"SE"},
  {Row: 0,Col: 1,Direction:"E"},
  {Row: -1,Col: 1,Direction:"NE"},
  {Row: -1,Col: 0,Direction:"N"},
}

// can be improved by using backtracking
func search(array [][]string, idx int, target string, start Point.Point,visited []Point.Point,cnt *int) {

  if idx == len(target)-1 && string(target[idx])== array[start.Row][start.Col]  {
    *cnt++
    return 
  } 
  if idx>= len(target)-1 || string(target[idx]) != array[start.Row][start.Col] || Includes(visited,start) {return}  

  visited = append(visited,start)

  for _,dir := range(directions){
    p := start.GetAdd(dir) 
    if (p.Outbound(len(array),len(array[0])) || Includes(visited,p) || (start.Direction != "" && start.Direction != p.Direction) ){continue}
    v:= make([]Point.Point,len(visited))
    copy(v,visited)
    search(array,idx +1,target,p,v,cnt) 
  }

}

var diagonals = []Point.Point{
  {Row: -1,Col: 1,Direction:"NE"},
  {Row: 1,Col: -1,Direction:"SW"},
  {Row: -1,Col: -1,Direction:"NW"},
  {Row: 1,Col: 1,Direction:"SE"},
}
func validate(arr [][]string, point Point.Point) bool{
  if(arr[point.Row][point.Col] != "A" ) {return false}
  for _,dir := range(diagonals){
    if point.GetAdd(dir).Outbound(len(arr),len(arr[0])) {return false}
  }

  validation_string := ""

  for _,dir := range(diagonals){
    p :=  point.GetAdd(dir)  
    validation_string += arr[p.Row][p.Col] 
  }

  pattern := `(SM|MS)(SM|MS)`
  regex := regexp.MustCompile(pattern)
  
  return regex.MatchString(validation_string) 
}

func part1(buffer []byte)int {
  input := arrayFromBuffuer(buffer)
  g_cnt := 0

  for i := 0; i < len(input); i++ {
    for j := 0; j < len(input[i]); j++ {
      start := Point.Point{Row: i,Col: j,Direction:""} 
      visited := []Point.Point{}
      search(input,0,"XMAS",start,visited,&g_cnt)
    }
  }

  return g_cnt
}
func part2(buffer []byte)int {
  input := arrayFromBuffuer(buffer)
  g_cnt := 0

  for i := 0; i < len(input); i++ {
    for j := 0; j < len(input[i]); j++ {
      p := Point.Point{Row: i,Col: j,Direction:""} 
      if validate(input,p){g_cnt++}
    }
  }

  return g_cnt
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
