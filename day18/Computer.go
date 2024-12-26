package main

import (
  "fmt"
  "strconv"
  "strings"
)

func ExtractInput(buffer string)Computer{
  trimed_buffer := strings.Trim(buffer,"\n")
  Obstables,rows,cols := map[Pos]int{},0,0
  ObstablePoll := []Pos{}
  lines := strings.Split(trimed_buffer,"\n")

  j:=0
  for i := range lines {
    parts := strings.Split(lines[i],",")
    row,_ := strconv.Atoi(parts[1])
    col,_ := strconv.Atoi(parts[0])
    pos :=Pos{row:row,col:col} 
    ObstablePoll = append(ObstablePoll,pos)
    j++
  }

  return Computer{
    Obstables: Obstables,
    ObstableIdx:0,
    ObstablePoll:ObstablePoll,
    Rows: rows,
    Cols: cols,
    Path:map[Pos]int{},
    Limit:j,
    End:nil,
  } 
}

type Computer struct {
  ObstableIdx int
  ObstablePoll []Pos
  Obstables map[Pos]int
  Path map[Pos]int
  Rows int 
  Cols int
  Limit int 
  End *Node
}

func (c *Computer) setDimension(rows int, cols int,limit int){
  c.Rows = rows
  c.Cols = cols
  c.Limit = limit
}

func (c Computer) IsValidPos(pos Pos) bool{
  inBound := pos.row>=0 && pos.col >=0 && pos.row<c.Rows && pos.col<c.Cols  
  if !inBound {return false}
  limit,exist := c.Obstables[pos]
  if exist {
    return limit>= c.Limit 
  }
  return true 
}

var DIRECTION = []Pos{
  {0, 1},
  {1, 0},
  {0, -1},
  {-1, 0},
}

func (c *Computer) Search(inital Node) {
  c.End = nil
  start := inital.Pos
  end := Pos{c.Rows-1,c.Cols-1}
  nodes := []Node{{Pos:start,Score:0,Prev:nil}}
  visited := map[Pos]int{}

  for key := range c.Path {
    delete(c.Path,key)
  }

  for len(nodes) >0 {
    node := nodes[0]
    nodes = nodes[1:]

    if node.Pos == end {
      c.End = &node

      for nd := &node;nd != nil;nd=(*nd).Prev {
        c.Path[nd.Pos] = nd.Score
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
      if !c.IsValidPos(n_node.Pos) {continue}
      if _,has := visited[n_node.Pos]; has{continue}

      visited[n_node.Pos]=n_node.Score
      nodes = append(nodes,n_node)
    }
  }
}

func (c Computer) Print(){
  for i :=0 ;i< c.Rows;i++ {
    line := ""
    for j := 0;j<c.Cols;j++{
      if limit,exist := c.Obstables[Pos{i,j}];exist && limit<c.Limit {
        line+="#"
      }else{
        line+="."
      }
    }
    fmt.Println(line)
  }
}

func (c Computer) PrintWithPath(path map[Pos]int){
  for i :=0 ;i< c.Rows;i++ {
    line := ""
    for j := 0;j<c.Cols;j++{
      pos := Pos{i,j}
      if limit,exist := c.Obstables[pos];exist && limit<c.Limit {
        if c.ObstableIdx<len(c.ObstablePoll) && pos == c.ObstablePoll[c.ObstableIdx]{
          line += "\033[31m #"
        }else {
          line+="\033[0m #"
        }
      }else if _,exist := path[pos];exist {
        line+="\033[0m O"
      }else{
        line+="\033[0m ."
      }
    }
    fmt.Println(line)
  }
}

type Node struct {
  Pos       Pos
  Score  int
  Prev *Node
}

type Pos struct {
  row,col int
}

func (p Pos) Add(other Pos)Pos{
  return Pos{p.row +other.row, p.col +other.col}
}

func (p *Pos) Set(other Pos){
  p.row = other.row
  p.col = other.col
}

func (p *Pos) AddS(other Pos){
  p.row += other.row
  p.col += other.col
}

func (n Node) move(direction Pos) Node {
  return Node{
    Pos:    n.Pos.Add(direction),
    Score:  n.Score + 1,
    Prev:   &n,
  }
}
