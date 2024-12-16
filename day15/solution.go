package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ExtractInput(buffer string)Warehouse{

  parts := strings.Split(buffer,"\n\n")
  r_position := Vect{1,1} 
  rr_position := Vect{1,1} 
  warehouse_map:= [][]string{}
  warehouse_map2:= [][]string{}
  moves := []Vect{}
  rows := strings.Split(parts[0],"\n")

  for i := range rows{
    rw := strings.Split(rows[i],"")
    rww := []string{}
    for j := range rw {
      switch string(rows[i][j]) {
      case "O":
        rww = append(rww,"[","]")
        break
      case "@":
        rww = append(rww,"@",".")
        break
      default:
        rww = append(rww,string(rows[i][j]),string(rows[i][j]))
        break
      }
      if rw[j] == "@" {
        r_position.Row = i 
        r_position.Col = j
        rr_position.Row = 2*i 
        rr_position.Col = 2*j
      }
    }
    warehouse_map = append(warehouse_map,rw)
    warehouse_map2 = append(warehouse_map2,rww)
  }

  reg := regexp.MustCompile("[<>v^]")
  matches := reg.FindAllString(parts[1], -1)

  for i := range matches {
    switch matches[i] {
    case "<":
      moves = append(moves,Vect{0,-1})
      break  
    case ">":
      moves = append(moves,Vect{0,1})
      break  
    case "v":
      moves = append(moves,Vect{1,0})
      break  
    case "^":
      moves = append(moves,Vect{-1,0})
      break  
    default:
      break
    }
  }

  return Warehouse{
    R_Position: r_position,
    Map: warehouse_map,
    Moves: moves,
  }
}

func ExtractInput2(buffer string)Warehouse{

  parts := strings.Split(buffer,"\n\n")
  r_position := Vect{1,1} 
  warehouse_map:= [][]string{}
  moves := []Vect{}
  rows := strings.Split(parts[0],"\n")

  for i := range rows{
    rw := []string{}
    for j := range strings.Split(rows[i],"") {
      switch string(rows[i][j]) {
      case "O":
        rw = append(rw,"[","]")
        break
      case "@":
        rw = append(rw,"@",".")
        break
      default:
        rw = append(rw,string(rows[i][j]),string(rows[i][j]))
        break
      }
      if string(rows[i][j]) == "@" {
        r_position.Row = i 
        r_position.Col = 2*j
      }
    }
    warehouse_map = append(warehouse_map,rw)
  }

  reg := regexp.MustCompile("[<>v^]")
  matches := reg.FindAllString(parts[1], -1)

  for i := range matches {
    switch matches[i] {
    case "<":
      moves = append(moves,Vect{0,-1})
      break  
    case ">":
      moves = append(moves,Vect{0,1})
      break  
    case "v":
      moves = append(moves,Vect{1,0})
      break  
    case "^":
      moves = append(moves,Vect{-1,0})
      break  
    default:
      break
    }
  }

  return Warehouse{
    R_Position: r_position,
    Map: warehouse_map,
    Moves: moves,
  }
}

type Vect struct {
  Row int
  Col int 
}

func (v *Vect) Add(other Vect)Vect{
  return Vect{v.Row +other.Row, v.Col +other.Col}
}

func (v *Vect) Equals(other Vect)bool{
  return v.Row == other.Row && v.Col  == other.Col
}

func (v *Vect) AddS(other Vect){
  v.Row +=other.Row 
  v.Col +=other.Col
}

func (v *Vect) Set(other Vect){
  v.Row = other.Row
  v.Col = other.Col
}

type Warehouse struct {
  R_Position Vect
  Map [][]string
  Moves []Vect
}

type MoveType struct {
  From Vect
  To Vect
  Copy bool
} 

type MoveGroup []MoveType

func (w *Warehouse) InBound(vect Vect) bool{
  return vect.Row >= 0 && vect.Col >= 0 && vect.Row < len(w.Map) && vect.Col < len(w.Map[0])
}

func (w *Warehouse) MoveTo( from Vect, to Vect){
  char := w.Map[from.Row][from.Col] 
  w.Map[from.Row][from.Col] =  w.Map[to.Row][to.Col] 
  w.Map[to.Row][to.Col] =  char
  w.R_Position.Set(to)
}

func (w *Warehouse) MoveToGroup(grp MoveGroup){

  //fmt.Println(grp)
  seen := map[string]int{}
  for i := range grp {
    j := len(grp) -i-1
    key := strconv.Itoa(grp[j].From.Row)+"-"+strconv.Itoa(grp[j].From.Col) + strconv.Itoa(grp[j].To.Row)+"-"+strconv.Itoa(grp[j].To.Col)  
    _,exist := seen[key]
    seen[key] = 1
    if exist {continue}

    from, to := grp[j].From,grp[j].To
    char := w.Map[from.Row][from.Col] 
    if grp[j].Copy { w.Map[from.Row][from.Col] =  w.Map[to.Row][to.Col]} 
    if !grp[j].Copy { w.Map[from.Row][from.Col] = "."} 
    w.Map[to.Row][to.Col] =  char
    w.R_Position.Set(to)
  }
}

func (w *Warehouse) Move(from Vect,dir Vect)bool{
  n_pos := from.Add(dir) 

  switch w.Map[n_pos.Row][n_pos.Col] {
  case ".":
    w.MoveTo(from,n_pos)
    return true
  case "O":
    if w.Move(n_pos,dir){
      w.MoveTo(from,n_pos)
      return true
    }
    break
  case "#":
  default:
    break
  }
  return false 
}

func (w *Warehouse) Push(from Vect,dir Vect,move_group *MoveGroup)bool{
  to := from.Add(dir) 
(*move_group) = append(*move_group,MoveType{
        From:from,
        To:to,
        Copy:true,
      })

  switch w.Map[to.Row][to.Col] {
  case ".":

    return true
  case "#":
    return false 
  case "]":
    v := w.Push(to,dir,move_group) 
    if !v {return false}
    if dir.Col != 0 {

      return true
    }else if w.Push(to.Add(Vect{0,-1}),dir,move_group)  {

      return true
    }
    break
  case "[":
    v := w.Push(to,dir,move_group) 
    if !v {return false}
    if dir.Col != 0 {
      
      return true
    }else if w.Push(to.Add(Vect{0,1}),dir,move_group)  {

      return true
    }
    return false 

  default:
    break
  }
  return false 
}


func (w *Warehouse) ComputeCoor()int{
  sum := 0
  for i := range w.Map {
    for j := range w.Map[i] {

      if regexp.MustCompile("[O\\[]").MatchString(w.Map[i][j]) {sum += 100*i+j}
    }
  }
  return sum
}

func (w *Warehouse) Print(){
  for i := range w.Map{
    for j := range w.Map[i]{
      fmt.Print(w.Map[i][j])
    }
    fmt.Println()
  }
}

func part1(w Warehouse)int {
  w.Print()
  for i := range w.Moves {
    w.Move(w.R_Position,w.Moves[i])
  }
  w.Print()
  return w.ComputeCoor()
}

func part2(w Warehouse)int {
  w.Print()

  for i := range w.Moves {
    move_group := MoveGroup{}
    valid := w.Push(w.R_Position,w.Moves[i],&move_group)
    if valid {
      w.MoveToGroup(move_group)
    }
  }
  w.Print()
  return w.ComputeCoor()
}

func main()  {
  // Open the file
  filePath := "input.txt"
  //filePath := "sample.txt"
  //filePath := "edge.txt"
  buffer, err := os.ReadFile(filePath)
  if err != nil {
    fmt.Println("Error opening file:", err)
    return
  }

  w1 := ExtractInput(strings.Trim(string(buffer),"\n"))
  w2 := ExtractInput2(strings.Trim(string(buffer),"\n"))

  fmt.Println(part1(w1))
  fmt.Println(part2(w2))
}
