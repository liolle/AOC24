package main

import (
  "fmt"
  "math"
  "os"
  "strings"
)

func ExtractInput(buffer string)Maze{
  lines := strings.Split(buffer,"\n")
  out := [][]rune{}
  start := Pos{0,0,0}
  end := Pos{0,0,0}

  for i := range lines {
    ll := strings.Split(lines[i],"")
    rn := []rune{}
    for j := range ll{
      char := []rune(ll[j])[0]
      if char == 'S' {
        start.Set(Pos{i,j,0})
      }else if char == 'E' {
        end.Set(Pos{i,j,0})
      }
      rn = append(rn,char)
    } 
    out = append(out,rn)
  } 
  return Maze{out,start,end}
}

type Maze struct {
  Map [][]rune
  Start Pos
  End Pos
}

func (m *Maze) IsValidState(n Node)bool{
  return m.Map[n.Pos.row][n.Pos.col] != '#' 
}

func (m *Maze) Found(n Node)bool{
  row,col := n.Pos.row,n.Pos.col
  char := m.Map[row][col] 
  return char == 'E' 
}

func (m *Maze) Print(seets []Pos){
  mp := [][]rune{}

  for _,row := range m.Map{
    mp = append(mp,row)
  }

  for i := range seets{
    mp[seets[i].row][seets[i].col] = 'O'
  }

  for i := range mp {
    for j := range mp[i]{
      if mp[i][j] =='.'{
        mp[i][j] = ' '
      }
    }
  }

  for i := range mp {
    for _,r := range mp[i]{
      fmt.Printf("%s ",string(r))
    }
    fmt.Println()
  }
}



func (m *Maze) Solve()(min_score int, best_seets []Node){
  nodes := []Node{{Pos:m.Start,Score:0,Prev:nil}}
  visited := map[Pos]int{}
  min_score = math.MaxInt
  found_node := []Node{}

  for len(nodes) >0 {
    node := nodes[0]
    nodes = nodes[1:]
    if node.Score > min_score {
      continue
    } 

    if m.Found(node){
      found_node = append(found_node,node)
      if node.Score <= min_score {
        min_score = node.Score
      }
      continue
    }

    mv := node.move()
    r2 := node.rotate(1)
    r1 := node.rotate(-1)
    dirs := []Node{r2,r1,mv}

    for i := range dirs{
      nd := dirs[i] 

      if !m.IsValidState(nd) {continue}
      if prev,has := visited[nd.Pos];has && prev < nd.Score{continue}

      visited[nd.Pos]=nd.Score
      nodes = append(nodes,nd)
    }
  }

  for i := range found_node{
    if found_node[i].Score == min_score {
      best_seets = append(best_seets,found_node[i])
    }
  }

  return min_score,best_seets 
}

func CountDistinct(arr []Pos)int{
  mp := map[Ps]int{}

  for i := range arr {
    _,exist := mp[arr[i].Extract()]

    if exist  {
      mp[arr[i].Extract()]++
    }else{
      mp[arr[i].Extract()]=1
    }
  }

  return len(mp)
}

func part1(m Maze)int {
  min_path,_ := m.Solve() 
  return min_path 
}

func part2(m Maze)int {
  _,best_seets := m.Solve() 
  seets := []Pos{}
  for i := range best_seets {
    for j := &best_seets[i];j != nil;j = (*j).Prev{
      seets = append(seets,j.Pos)
    }
  }
  m.Print(seets)
  return CountDistinct(seets) 
}

func main()  {
  // Open the file
  filePath := "input.txt"
  //filePath := "sample.txt"
  //filePath := "sample2.txt"
  buffer, err := os.ReadFile(filePath)
  if err != nil {
    fmt.Println("Error opening file:", err)
    return
  }

  m1 := ExtractInput(strings.Trim(string(buffer),"\n"))
  m2 := ExtractInput(strings.Trim(string(buffer),"\n"))

  fmt.Println(part1(m1))
  fmt.Println(part2(m2))
}
