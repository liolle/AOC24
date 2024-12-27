package main

import (
	"math"
	"strings"
)

func ExtractInput(buffer string)Race{
  track := [][]string{} 
  start := Pos{0,0}
  end := Pos{0,0}
  lines := strings.Split(strings.Trim(buffer,"\n"),"\n")

  for i := range lines {
    row := strings.Split(lines[i],"")
    for col := range row {
      if row[col] == "S" {
        start.row = i
        start.col = col
      }

      if row[col] == "E" {
        end.row = i
        end.col = col
      }
    }
    track = append(track,row)
  } 

  return Race{
    Track:track,
    Start:start,
    Path:map[Pos]Node{},
    End:end,
  }
}

func (r *Race) InBound(node Node) bool{
  return  node.Pos.row>=0 && node.Pos.col >=0 && node.Pos.row<len(r.Track) && node.Pos.col<len(r.Track[0]) 
}


func (r *Race) IsValidPos(node Node) bool{
  if !r.InBound(node) {return false}
  return  r.Track[node.Pos.row][node.Pos.col] != "#" || node.CheatLimit > node.Score  
}

func (r *Race) BFS(){

  start := r.Start 
  end := r.End 
  nodes := []Node{{Pos:start,Score:0,Prev:nil,CheatLimit:0}}
  visited := map[Pos]int{}

  for len(nodes)>0 {
    node := nodes[0]
    nodes = nodes[1:]

    if node.Pos == end {
      for p := &node;p!=nil;p=p.Prev{
        r.Path[p.Pos] = *p
      }
      return 
    }

    dirs := []Pos{
      DIRECTION[0],
      DIRECTION[1],
      DIRECTION[2],
      DIRECTION[3],
    }

    for i := range dirs{
      nd := dirs[i] 
      n_node := node.move(nd)      
      if !r.IsValidPos(n_node) {continue}
      if _,has := visited[n_node.Pos]; has{continue}

      visited[n_node.Pos]=n_node.Score
      nodes = append(nodes,n_node)
    }
  }
}

func (r *Race) Cheat(node Node, cheat_duration int)[]int{
  max_gain := []int{}

  for row := range r.Track {
    for col,val := range r.Track[row]{
      dist := math.Abs(float64(node.Pos.row) - float64(row)) + math.Abs(float64(node.Pos.col) - float64(col))
      if dist >float64(cheat_duration) {continue}
      if val == "#" {continue} 
      s1 :=r.Path[Pos{row,col}].Score 
      s2 :=r.Path[node.Pos].Score  
      gain := s1 - s2 - int(dist)   
      if gain<0 {continue}
      max_gain = append(max_gain,gain)
    }
  }

  return max_gain
}

type Race struct {
  Track [][]string
  Start Pos
  Path map[Pos]Node
  End Pos
}

var DIRECTION = []Pos{
  {0, 1},
  {1, 0},
  {0, -1},
  {-1, 0},
}

type Pos struct {
  row,col int
}

func (p Pos) Add(other Pos)Pos{
  return Pos{p.row +other.row, p.col +other.col}
}

type Node struct {
  Pos       Pos
  Score  int
  CheatLimit int
  Prev *Node
}

func (n Node) move(direction Pos) Node {
  return Node{
    Pos:    n.Pos.Add(direction),
    Score:  n.Score + 1,
    Prev:   &n,
    CheatLimit: n.CheatLimit,
  }
}
