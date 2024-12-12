package main

import (
	"fmt"
	"os"
	"strings"
)

func ExtractInput(buffer string) Garden{
  out := [][]string{}
  lines := strings.Split(buffer,"\n")
  r := len(lines)
  c := 0

  for _,line := range(lines){
    c = len(line)
    row := []string{}

    for _,char := range(line){
      row = append(row,string(char))
    }
    out = append(out,row)
  }

  return Garden{
    Map:out,
    Rows:r,
    Cols:c,
    Sections: [][]Cell{},
  }
}

type Garden struct {
  Map [][]string
  Rows int
  Cols int
  Sections [][]Cell 
}

type Cell struct{
  Row int 
  Col int 
}

func (c Cell) GetAdd(other Cell) Cell{
  return Cell{
    c.Row + other.Row,
    c.Col + other.Col,
  }
}

var DIRECTION = []Cell{
  {-1,0},
  {0,1},
  {1,0},
  {0,-1},
}

func (c *Cell) Equals(other Cell)bool{
  return c.Col == other.Col && c.Row ==other.Row
}

func Includes(slice []Cell, value Cell) bool {
  for _,v := range(slice) {
    if v.Equals(value) {
      return true
    }
  }
  return false
}

func (g *Garden) Split (){
  visited := []Cell{}
  for r,row := range(g.Map){
    for c := range row {
      cell :=Cell{r,c} 
      if !Includes(visited,cell){
        group := []Cell{}
        g.DFS(cell,&visited,&group)
        g.Sections = append(g.Sections,group)
      }
    }
  }
}

func (g *Garden) InBound(cell Cell)bool  {
  return cell.Row >= 0 && cell.Col >=0 && cell.Row < g.Rows && cell.Col < g.Cols 
}

func (g *Garden) DFS (current Cell,visited *[]Cell,group *[]Cell){
  if Includes(*visited,current){return}

  *visited = append(*visited,current)
  *group = append(*group,current)

  for _,cell := range(DIRECTION){
    n_cell := Cell{cell.Row+current.Row,cell.Col+current.Col} 
    if g.InBound(n_cell) && g.Map[n_cell.Row][n_cell.Col] == g.Map[current.Row][current.Col] {
      g.DFS(n_cell,visited,group)
    }
  }
}

func (g * Garden) ComputeScore(cells []Cell) int{
  border_cnt := 0
  for _,cell := range(cells){
    for _,dir := range(DIRECTION){
      n_cell := Cell{cell.Row+dir.Row,cell.Col+dir.Col} 
      if !g.InBound(n_cell) || !(g.Map[n_cell.Row][n_cell.Col] == g.Map[cell.Row][cell.Col]){
        border_cnt++
      }
    }
  } 
  return border_cnt
}

//Outer corner Pain
var CORNER_DIRECTION = [][]Cell{
  {{-1,0},{0,-1},{-1,-1}},
  {{-1,0},{0,1},{-1,1}},
  {{1,0},{0,1},{1,1}},
  {{1,0},{0,-1},{1,-1}},
} 

func (g Garden) CornerCount(cell Cell) int{
  cnt := 0

  //Inner corner 
  for _,dir := range(CORNER_DIRECTION){
    c1 := cell.GetAdd(dir[0]) 
    c2 := cell.GetAdd(dir[1])     
    c3 := cell.GetAdd(dir[2])

    if !g.InBound(c1) || !g.InBound(c2) {continue}

    char_cell := g.Map[cell.Row][cell.Col] 
    char_c1 := g.Map[c1.Row][c1.Col] 
    char_c2 := g.Map[c2.Row][c2.Col] 
    char_c3 := g.Map[c3.Row][c3.Col] 

    if char_c1 != char_cell || char_c2 != char_cell  {continue}
    if  char_c3 == char_cell {continue}
    cnt++
  }

  //Outer corner
  for _,dir := range(CORNER_DIRECTION){
    c1 := cell.GetAdd(dir[0]) 
    c2 := cell.GetAdd(dir[1])     
    char_cell := g.Map[cell.Row][cell.Col] 

    if !g.InBound(c1) && !g.InBound(c2) {
      cnt++
      continue
    }else if !g.InBound(c1) {
      char_c2 := g.Map[c2.Row][c2.Col] 
      if char_c2 != char_cell {
        cnt++
      }
    }else if !g.InBound(c2) {
      char_c1 := g.Map[c1.Row][c1.Col] 
      if char_c1 != char_cell {
        cnt++
      }
    }else {
      char_c1 := g.Map[c1.Row][c1.Col] 
      char_c2 := g.Map[c2.Row][c2.Col] 
      if char_c1 != char_cell && char_c2 != char_cell  {
        cnt++
      }
    }
  }

  return cnt
}

func (g * Garden) ComputeScoreSide(cells []Cell) int{
  corner_cnt := 0
  for _,cell := range(cells){
    corner_cnt += g.CornerCount(cell)
  } 
  return corner_cnt
}

func part1(g Garden)int {
  sum := 0
  g.Split()

  for _,section :=range(g.Sections){
    sum += len(section) * g.ComputeScore(section)  
  }
  return sum
}

func part2(g Garden)int {
  sum := 0
  g.Split()
  
  for _,section :=range(g.Sections){
    sum += len(section) * g.ComputeScoreSide(section)    
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

  g1 := ExtractInput(strings.Trim(string(buffer),"\n"))
  for _,row := range(g1.Map){
    fmt.Println(row)
  }
  fmt.Println(part1(g1))
  fmt.Println(part2(g1))
}
