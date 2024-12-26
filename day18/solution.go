package main

import (
  "fmt"
  "os"
)

func part1(c Computer)int {
  node := Node{
    Pos: Pos{0,0},
    Score: 0,
    Prev:nil,
  }

  for i := 0;i<c.Limit;i++{
    c.Obstables[c.ObstablePoll[i]] = i
  }

  c.ObstableIdx = c.Limit
  c.Search(node)
  c.PrintWithPath(c.Path)
  return len(c.Path)-1
}
func part2(c Computer)int {
  node := Node{
    Pos: Pos{0,0},
    Score: 0,
    Prev:nil,
  }
  c.Search(node)

  for c.End != nil {
    if c.ObstableIdx >= len(c.ObstablePoll){break}
    if(c.ObstableIdx >= c.Limit){break}
    nex_obs := c.ObstablePoll[c.ObstableIdx]
    _,hit := c.Path[nex_obs]
    for !hit && c.ObstableIdx<len(c.ObstablePoll)-1{
      _,h :=c.Path[nex_obs] 
      hit = h
      if hit {break}
      c.Obstables[nex_obs] = c.ObstableIdx
      c.ObstableIdx++
      nex_obs = c.ObstablePoll[c.ObstableIdx]
    }

    if c.ObstableIdx >= len(c.ObstablePoll){break}
    if(c.ObstableIdx >= c.Limit){break}

    nd := c.End
    for  ;nd !=nil && nd.Pos != nex_obs;nd = nd.Prev{
      delete(c.Path,nd.Pos)
    }
    c.Obstables[nex_obs] = c.ObstableIdx
    if nd != nil {
      delete(c.Path,nd.Pos)
      nd = nd.Prev
    }
    c.Search(node)
  }

  c.PrintWithPath(c.Path)
  fmt.Println(c.ObstableIdx,c.ObstablePoll[c.ObstableIdx])
  return c.ObstableIdx
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

  c1 := ExtractInput(string(buffer))
  c2 := ExtractInput(string(buffer))
  c1.setDimension(71,71,1024)
  c2.setDimension(71,71,3450)
  //c1.setDimension(7,7,12)
  //c2.setDimension(7,7,25)

  fmt.Println(part1(c1))
  fmt.Println(part2(c2))
}
