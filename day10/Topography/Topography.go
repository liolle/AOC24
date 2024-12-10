package topography

import (
	"fmt"
	"strconv"
)

type Topography struct {
  Map [][]int
  Rows int 
  Cols int
  Solution []Trailhead
}

type Trailhead struct {
  Row int 
  Col int
  Score int
}

var DIRECTIONS = []Trailhead{
  {0,-1,0},
  {1,0,0},
  {0,1,0},
  {-1,0,0},
}

func Includes(slice []Trailhead, value Trailhead) bool {
  for _,v := range(slice) {
    if v.Equals(value) {
      return true
    }
  }
  return false
}

func IncludesString(slice []string, value string) bool {
  for _,v := range(slice) {
    if v == value {
      return true
    }
  }
  return false
}

func (t *Trailhead) Equals(other Trailhead)bool{
  return t.Row == other.Row && t.Col == other.Col
}

func (topography Topography) Print(){

  for _,r := range(topography.Map){
    fmt.Println(r)
  }

  fmt.Println()

  for _,t := range(topography.Solution){
    fmt.Println(t)
  }
}

func (t *Topography) Inbound(x int,y int)bool{
  return x>=0 && y>=0 && x<t.Cols && y<t.Rows
}

func (top *Topography) Walk(start *Trailhead, row int, col int,found *[]string){
  if (top.Map[row][col] == 9){
    key := strconv.Itoa(start.Row) + strconv.Itoa(start.Col) + strconv.Itoa(row) + strconv.Itoa(col) 
    if(!IncludesString(*found,key)){
      start.Score += 1
      *found = append(*found,key)
    }
    return
  }

  for _,dir := range(DIRECTIONS){
    nx := dir.Row+row 
    ny := dir.Col+col 
    if top.Inbound(nx,ny) && top.Map[nx][ny] == top.Map[row][col] +1{
      top.Walk(start,nx,ny, found)
    }
  }
} 

func (top *Topography) Slide(start *Trailhead, row int, col int){
  if (top.Map[row][col] == 9){
    start.Score += 1
    return
  }

  for _,dir := range(DIRECTIONS){
    nx := dir.Row+row 
    ny := dir.Col+col 
    if top.Inbound(nx,ny) && top.Map[nx][ny] == top.Map[row][col] +1{
      top.Slide(start,nx,ny)
    }
  }
}

func (top *Topography) Scan() {
  visited := []string{}

  for i,r := range(top.Map){
    for j,c := range(r){
      if c == 0 {
        t := Trailhead{j,i,0}
        top.Walk(&t,i,j,&visited)
        top.Solution = append(top.Solution,t)
      }
    }
  }
}

func (top *Topography) Rate() {

  for i,r := range(top.Map){
    for j,c := range(r){
      if c == 0 {
        t := Trailhead{j,i,0}
        top.Slide(&t,i,j)
        top.Solution = append(top.Solution,t)
      }
    }
  }
}

