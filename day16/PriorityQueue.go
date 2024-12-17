package main

var DIRECTION = []Vect{
  {0,1},
  {1,0},
  {0,-1},
  {-1,0},
}

type Node struct {
  Pos Vect
  Direction int
  Priority int
  Index int
  Path []Vect
}

func normalize(val int, limit int)int{
  val = val%limit
  if val <0 {val+=limit}
  return val 
}

func (n *Node) move()Node{
  path := make([]Vect,len(n.Path))
  copy(path,n.Path)
  path = append(path,n.Pos)
  return Node{
    Pos:n.Pos.Add(DIRECTION[n.Direction]),
    Direction:n.Direction,
    Priority:n.Priority +1,
    Index:n.Index,
    Path:path,
  }
}

func (n *Node) rotate(r int)Node{
  return Node{
    Pos:n.Pos,
    Direction:normalize(n.Direction+r,4),
    Priority:n.Priority +1000,
    Index:n.Index,
    Path:n.Path,
  }
}

type Nodes []*Node

func (ns Nodes) Len() int { return len(ns) }

func (ns Nodes) Less(i, j int) bool {
  return ns[i].Priority < ns[j].Priority
}

func (ns Nodes) Swap(i, j int) {
  ns[i], ns[j] = ns[j], ns[i]
  ns[i].Index = i 
  ns[j].Index = j
}

func (ns *Nodes) Push(x interface{}) {
  n := len(*ns)
  item := x.(*Node)
  item.Index = n 
  *ns = append(*ns, item)
}

func (ns *Nodes) Pop() interface{} {
  old := *ns
  n := len(old)
  item := old[n-1]
  item.Index = -1 
  *ns = old[0 : n-1]
  return item
}
