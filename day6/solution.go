package main

import (
	maze "aoc24/Maze"
	"fmt"
	"os"
)

func part1(maze maze.Maze)int{
  can_walk := true

  for can_walk {
    can_walk = maze.Walk()
  }

  return maze.VisitedCells +1 
}

func part2( m maze.Maze) int{

  can_walk := true
  cycle := false 
  path := []maze.Position{}

  for can_walk {
    can_walk = m.Travel(&path,&cycle)
  }

  count := 0
  for _,p := range(path[1:]){
    m.Clear()

    cycle := false 
    can_walk = true
    pt := [] maze.Position{} 
    m.AddObstacle(p.Row,p.Col)
    for can_walk && !cycle {
      can_walk = m.Travel(&pt,&cycle)
    }

    if cycle {
      if !maze.IncludesS(m.NewObstacles,p) && !p.EqualsS(m.StartPoint){
        fmt.Println(p)
        m.NewObstacles = append(m.NewObstacles,p)
      }
      count++
    }

    m.RemoveObstacle(p.Row,p.Col)
  }

  m.Clear()
  m.PrintP2()
  return len(m.NewObstacles)
}

func main()  {
  filePath := "input.txt"
  //filePath := "sample.txt"
  //filePath := "edge1.txt"
  buffer, err := os.ReadFile(filePath)
  if err != nil {
    fmt.Println("Error opening file:", err)
    return
  }

  m1 := maze.NewMaze(string(buffer))
  m2 := maze.NewMaze(string(buffer)) 
  fmt.Println(part1(m1))
  fmt.Println(part2(m2))
}
