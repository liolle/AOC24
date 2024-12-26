package main

import (
	"slices"
	"strings"
)

func ExtractInput(buffer string)Layout{ 
  parts := strings.Split(buffer,"\n\n")
  pat := strings.Split(parts[0],", ")
  patterns := []Pattern{}
  cache := map[string]int{}

  for i := range pat {
    patterns = append(patterns, Pattern(pat[i]))
  }

  slices.SortFunc(patterns,ComparePatterns)

  return Layout{
    Patterns: patterns,
    Desighs: strings.Split(strings.Trim(parts[1],"\n"),"\n"),
    Cache:cache,
  }
}

type Pattern string 

type Layout struct {
  Patterns []Pattern
  Desighs []string
  Cache map[string]int
}

func ComparePatterns(a Pattern, b Pattern)int{
  return len(a)-len(b)
}

func (p Pattern) StartWith(pattern string)bool{
  if (len(p)>len(pattern)){return false}
  
  for i:=0;i<len(p);i++{
    if p[i] != pattern[i] {return false}
  }
  return true
}

func (l *Layout) MatchStart(design string)[]Pattern{
  out := []Pattern{}
  for i := range l.Patterns{
    if l.Patterns[i].StartWith(design){
      out = append(out,l.Patterns[i])
    } 
  }
  return out
}

func (l *Layout) Solve(node Node)int {
  if node.Desigh == "" {return 1} 
  if val,has := l.Cache[node.Desigh];has{return val}

  found := 0 
  matching_patterns := l.MatchStart(node.Desigh)

  for i := range matching_patterns {
    nd :=node.move(matching_patterns[i]) 
    fd :=  l.Solve(nd)
    found += fd
  }

  l.Cache[node.Desigh] = found

  return found 
}

type Node struct {
  Desigh string
  Score int
  Prev *Node
}

func (n *Node) move(pattern Pattern) Node{
  return Node{
    Desigh:n.Desigh[len(pattern):],
    Score: n.Score+1,
    Prev: n,
  }
}
