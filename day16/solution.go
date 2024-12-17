package main

import (
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func ExtractInput(buffer string)Maze{
  lines := strings.Split(buffer,"\n")
  out := [][]string{}
  pos := Vect{0,0}

  for i := range lines {
    ll := strings.Split(lines[i],"")
    for j := range ll{
      if ll[j] == "S" {
        pos.Set(Vect{i,j})
      }
    } 
    out = append(out,ll)
  } 

  return Maze{out,pos}
}

type Maze struct {
  Map [][]string
  Start Vect
}

func (m *Maze) IsValidState(n Node)bool{
  return m.Map[n.Pos.Row][n.Pos.Col] != "#" 
}

func (m *Maze) Found(n Node)bool{
  return m.Map[n.Pos.Row][n.Pos.Col] == "E"
}

func (m *Maze) Print(seets []Vect){

  mp := [][]string{}
  
  for _,row := range m.Map{
    mp = append(mp,row)
  }

  for i := range seets{
    mp[seets[i].Row][seets[i].Col] = "O"
    
  }

  for i := range mp {
    fmt.Println(mp[i])
  }
}

func (m *Maze) Solve()(int,[]Vect){
  nodes := make(Nodes,0)
  visited := map[string]int{}
  heap.Init(&nodes)
  heap.Push(&nodes,&Node{Pos:m.Start,Direction:0,Priority:0})

  best_path_points := map[string]Vect{}
  min_Score := math.MaxInt32

  for len(nodes) >0 {
    node := heap.Pop(&nodes).(*Node)
    key := strconv.FormatInt(int64(node.Pos.Row),10)+"."+strconv.FormatInt(int64(node.Pos.Col),10)+"."+ strconv.FormatInt(int64(node.Direction),10) 
    _,exist := visited[key]
    if exist  {continue}
    visited[key] = 1
    if m.Found(*node){

        fmt.Println(node.Priority)
      if node.Priority <= min_Score {
        for i := range node.Path {
          key := strconv.FormatInt(int64(node.Path[i].Row),10)+"."+strconv.FormatInt(int64(node.Path[i].Col),10)
          _,exist := best_path_points[key]
          if exist {continue}
          best_path_points[key] =node.Path[i] 
        }
        min_Score = node.Priority

      }else{
        out := []Vect{} 
        for key :=range best_path_points {
          out = append(out, best_path_points[key])
        }
        return min_Score,out
      }
    }

    if !m.IsValidState(*node){continue}

    r2 := node.rotate(1)
    r1 := node.rotate(-1)
    mv := node.move()

    heap.Push(&nodes,&r2)
    heap.Push(&nodes,&r1)
    heap.Push(&nodes,&mv)
  }

  if min_Score == math.MaxInt32 {
    return -1,[]Vect{}
  }else {
    out := []Vect{} 
    for key :=range best_path_points {
      out = append(out, best_path_points[key])
    }
    return min_Score,out
  }
}


func part1(m Maze)int {
  min_path,best_sites := m.Solve() 
  for i := range best_sites{
    fmt.Println(best_sites[i])
  }

  m.Print(best_sites)

  return min_path 
}
func part2()int {
  return 0
}

func main()  {
  // Open the file
  //filePath := "input.txt"
  filePath := "sample.txt"
  buffer, err := os.ReadFile(filePath)
  if err != nil {
    fmt.Println("Error opening file:", err)
    return
  }

  m1 := ExtractInput(strings.Trim(string(buffer),"\n"))

  fmt.Println(part1(m1))
}
