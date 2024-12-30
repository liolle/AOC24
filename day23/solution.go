package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
)

func part1(g Graph)int {
  res := g.FindCycles(3)
  cycles := []string{}
  seen := map[string]bool{}
  for key := range res {
    cycle_from_key := res[key] 
    for i := range cycle_from_key {
      if _,has := seen[cycle_from_key[i]] ;!has {
        cycle := NewCycleFromString(cycle_from_key[i])
        cycle.Sort()
        cycle.Permutations()
        pick :=cycle.Per[0]

        if regexp.MustCompile("^(..)*t").MatchString(pick){
          cycles = append(cycles,pick)
        }
        for i := range cycle.Per {
          seen[cycle.Per[i]] = true
        }
      }
    }
  }

  return len(cycles)
}

func part2(g Graph)(out string) {

  R := []string{}
  P := []string{}
  X := []string{}
 
  for key := range g.AdjMatrix{
    P = append(P,key)
  }
  
  g.FindClique(R,P,X)
  largest_clique_idx := 0

  for i := range g.Clique{
    if len(g.Clique[i]) > len(g.Clique[largest_clique_idx]){
      largest_clique_idx = i
    }
  }

  slices.Sort(g.Clique[largest_clique_idx])
  out = ""

  for i := range g.Clique[largest_clique_idx]{
    if out != ""{out += ","}
    out+=g.Clique[largest_clique_idx][i]
  }

  return out  
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

  g1 := ExtractInput(string(buffer))
  g2 := ExtractInput(string(buffer))
  fmt.Println(part1(g1))
  fmt.Println(part2(g2))

}
