package main

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

func ExtractInput(buffer string)(out Graph){
  lines := strings.Split(strings.Trim(buffer,"\n"),"\n")
  out.AdjMatrix = map[string][]string{}
  out.cache = map[CacheKey][]string{}

  for i :=range lines {
    parts := strings.Split(lines[i],"-")
   out.AdjMatrix[parts[0]] = append(out.AdjMatrix[parts[0]],parts[1]) 
   out.AdjMatrix[parts[1]] = append(out.AdjMatrix[parts[1]],parts[0]) 
  }

  return out
}

type Cycle struct {
  Computers [][2]rune
  Per []string
}

func NewCycleFromString(str string)(out Cycle){
  computers := [][2]rune{}
  for i := 0;i<len(str);i+=2 {
    r := [2]rune{
      rune(str[i]),
      rune(str[i+1]),
    }
    computers = append(computers,r)
    out.Computers = computers
  }
  out.Computers = computers
  out.Per = []string{}
  return out
}

func (c *Cycle) PermutationR(computers [][2]rune, marked []bool){

  if len(computers) == len(c.Computers){
    if c.Per == nil {c.Per = []string{}}
    str := []rune{}

    for i := range computers {
      str = append(str,(computers)[i][0])
      str = append(str,(computers)[i][1])
    }
    c.Per = append(c.Per,string(str))
    return
  }

  for i := 0 ;i<len(c.Computers);i++{
    if marked[i]{continue} 
    marked[i] =true
    computers = append(computers,c.Computers[i])
    c.PermutationR(computers,marked)
    computers = (computers)[:max(0,len(computers)-1)]
    marked[i] =false
  }

}

func (c *Cycle) Permutations(){
  marked := make([]bool,len(c.Computers))
  c.PermutationR([][2]rune{},marked)
}


func (c Cycle) Sort(){
  sort.Slice(c.Computers, func(i, j int) bool {
    if c.Computers[i][0] != c.Computers[j][0] {
      return c.Computers[i][0] < c.Computers[j][0]
    }

    return c.Computers[i][1] < c.Computers[j][1]
  })
}

type CacheKey struct {
  depth int
  current string
  start string
}

type Graph struct {
  AdjMatrix map[string][]string
  cache map[CacheKey][]string
  Clique [][]string
}

func (g *Graph) DFSLimited (depth int, start string,current string,visited *map[string]bool)(out []string){
  if _,has := (*visited)[current]; has {return []string{}}
  (*visited)[current] = true
  key :=CacheKey{depth,current,start} 
  if _,has := g.cache[key];has {
    return g.cache[key]
  }

  if depth == 1  {
    return []string{current}
  } 

  nb := g.AdjMatrix[current]

  for i := range nb {
    res := g.DFSLimited(depth-1,start,nb[i],visited)  
    for i := range res {
      last := res[i][max(0,len(res[i])-2):]  
      if ! slices.Contains(g.AdjMatrix[last],start) {continue} 
      if strings.Contains(res[i],current) {continue}

      out = append(out,current+res[i])
    } 
    delete(*visited,nb[i])
  }
  g.cache[key] = out
  return out
}

func (g *Graph) FindCycle(start string, depth int)(out []string) {
  visited := map[string]bool{}
  arr :=g.DFSLimited(depth,start,start,&visited)
  return arr 
}

func (g *Graph) FindCycles(depth int)(out map[string][]string) {
  out = map[string][]string{}
  for key := range g.AdjMatrix {
    cycles := []string{}
    res :=g.FindCycle(key,depth)  
    for i := range res {
      cycles = append(cycles,res[i])
    }
    out[key] = cycles 
  }
  return out
}

func (g *Graph) FindClique(R []string,P []string,X []string) {

  if len(P) == 0 && len(X) == 0 {
    r := slices.Clone(R)
    g.Clique = append(g.Clique,r)
    return
  }

  for i := range P {

    key := P[i]
    R = append(R,key)
    PP := []string{} 
    XX := []string{} 

    nb:= g.AdjMatrix[key]
  
    for i := range nb {
      if slices.Contains(X,nb[i]) {continue}
      if !slices.Contains(P,nb[i]) {continue}
      PP = append(PP,nb[i])
    } 

    for i := range nb {
      if !slices.Contains(X,nb[i]) {continue}
      XX = append(XX,nb[i])
    } 

    g.FindClique(R,PP,XX)

    R = R[:len(R)-1]
    X = append(X, key)
  }

}

func (g Graph) Print(){
  for key :=range g.AdjMatrix{
    fmt.Println(fmt.Sprintf("%s : %s",key,g.AdjMatrix[key]))
  }
}
