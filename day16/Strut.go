package main

var DIRECTION = []Pos{
  {0, 1,0},
  {1, 0,1},
  {0, -1,2},
  {-1, 0,3},
}

type Node struct {
  Pos       Pos
  Score  int
  Prev *Node
}

type Pos struct {
  row,col,dir int
}

type Ps struct {
  row,col int
}

func (p *Pos) Extract()Ps{
  return Ps{p.row,p.col}
}

func normalize(val int, limit int) int {
  val = val % limit
  if val < 0 {
    val += limit
  }
  return val
}

func (p *Pos) Add(other Pos)Pos{
  return Pos{p.row +other.row, p.col +other.col,other.dir}
}

func (p *Pos) Set(other Pos){
  p.row = other.row
  p.col = other.col
  p.dir = other.dir
}

func (p *Pos) AddS(other Pos){
  p.row += other.row
  p.col += other.col
}


func (n *Node) move() Node {
  return Node{
    Pos:    n.Pos.Add(DIRECTION[n.Pos.dir]),
    Score:  n.Score + 1,
    Prev:   n,
  }
}

func (n *Node) rotate(r int) Node {
  return Node{
    Pos:    n.Pos.Add(DIRECTION[normalize(r+n.Pos.dir,4)]),
    Score:  n.Score + 1001,
    Prev:   n,
  }
}
