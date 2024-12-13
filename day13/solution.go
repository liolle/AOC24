package main

import (
  "container/heap"
  "fmt"
  "os"
  "regexp"
  "strconv"
  "strings"
)

func ExtractInput(buffer string,offeset int64 )[]Machine{
  machines := []Machine{} 
  groups := strings.Split(buffer,"\n\n") 

  xb_pattern := `X\+([0-9]+)`
  yb_pattern := `Y\+([0-9]+)`

  xp_pattern := `X=([0-9]+)`
  yp_pattern := `Y=([0-9]+)`

  for _,g := range groups {
    lines := strings.Split(g,"\n")

    ax,_ := strconv.Atoi(regexp.MustCompile(xb_pattern).FindStringSubmatch(lines[0])[1])
    ay,_ := strconv.Atoi(regexp.MustCompile(yb_pattern).FindStringSubmatch(lines[0])[1])

    bx,_ := strconv.Atoi(regexp.MustCompile(xb_pattern).FindStringSubmatch(lines[1])[1])
    by,_ := strconv.Atoi(regexp.MustCompile(yb_pattern).FindStringSubmatch(lines[1])[1])

    px,_ := strconv.Atoi(regexp.MustCompile(xp_pattern).FindStringSubmatch(lines[2])[1])
    py,_ := strconv.Atoi(regexp.MustCompile(yp_pattern).FindStringSubmatch(lines[2])[1])

    machines = append(machines, Machine{
      A: Vect{int64(ax),int64(ay)},
      B: Vect{int64(bx),int64(by)},
      Prize: Vect{int64(px)+offeset,int64(py)+offeset},
    })
  }

  return machines
}

type Machine struct {
  A Vect
  B Vect
  Prize Vect
}

type Vect struct {
  Row int64
  Col int64
}

type State struct {
  Pos Vect 
  Priority int64 
  Index    int64  
  A int64 
  B int64
  Steps int64
}

type MState []*State

func (pq MState) Len() int { return len(pq) }

func (pq MState) Less(i, j int) bool {
  return pq[i].Priority < pq[j].Priority
}

func (pq MState) Swap(i, j int) {
  pq[i], pq[j] = pq[j], pq[i]
  pq[i].Index = int64(i)
  pq[j].Index = int64(j)
}

func (pq *MState) Push(x interface{}) {
  n := len(*pq)
  item := x.(*State)
  item.Index = int64(n)
  *pq = append(*pq, item)
}

func (pq *MState) Pop() interface{} {
  old := *pq
  n := len(old)
  item := old[n-1]
  item.Index = -1 
  *pq = old[0 : n-1]
  return item
}

func (m *Machine) IsValidState(s State)bool{
  return s.Pos.Row <= m.Prize.Row && s.Pos.Col <= m.Prize.Col && s.A < 100 && s.B < 100  
}

func (m *Machine) Found(vect Vect)bool{
  return vect.Row == m.Prize.Row && vect.Col == m.Prize.Col
}
func (m *Machine) Search() int64{

  pq := make(MState,0)
  visited := map[string]int64{}
  heap.Init(&pq)
  heap.Push(&pq,&State{Pos:Vect{int64(0),int64(0)},Priority:int64(0),Steps:int64(0)})

  for pq.Len() >0{
    next := heap.Pop(&pq).(*State)
    key := strconv.FormatInt(next.Pos.Row,10)+"."+strconv.FormatInt(next.Pos.Col,10) 
    _,include := visited[key]
    if include  {continue}

    visited[key] = next.Steps

    if m.Found(next.Pos){return next.Priority} 
    if !m.IsValidState(*next){continue}

    heap.Push(&pq,&State{Pos:Vect{next.Pos.Row+m.A.Row,next.Pos.Col+m.A.Col},Priority:next.Priority+3, A:next.A+1, B:next.B, Steps:next.Steps +1})
    heap.Push(&pq,&State{Pos:Vect{next.Pos.Row+m.B.Row,next.Pos.Col+m.B.Col},Priority:next.Priority+1, A:next.A, B:next.B+1, Steps:next.Steps +1})
  }

  return -1
}

func part1(machines []Machine)int64 {
  sum := int64(0)
  for _,m := range machines {
    res := m.Search()
    if res >0{
      sum += res 
    }
  }
  return sum 
}

func (m *Machine) Solve() (int, int, error) {
  a := m.A.Row
  b:= m.B.Row
  c := m.Prize.Row
  d := m.A.Col
  e := m.B.Col
  f := m.Prize.Col

  det := a*e - b*d

  if det == 0 {
    return 0, 0, fmt.Errorf("determinant = 0")
  }

  A := (c*e - b*f) / det
  B := (a*f - c*d) / det
  return int(A), int(B), nil
}

func part2(machines []Machine)int64 {
  sum :=0
  for _,m := range machines {
    x,y,err := m.Solve()
    if err != nil{
      continue
    }
    vx := x*int(m.A.Row)+y*int(m.B.Row)
    vy := x*int(m.A.Col)+y*int(m.B.Col)
    if vx != int(m.Prize.Row) || vy != int(m.Prize.Col) {
      continue
    } 
    sum += x*3 + y
  }
  return int64(sum)
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

  m1 := ExtractInput(string(buffer),0)
  m2 := ExtractInput(string(buffer),10000000000000)

  fmt.Println(part1(m1))
  fmt.Println(part2(m2))
}
