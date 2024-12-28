package main

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)
type KeyPadInput int
type NumPadInput int

const (
    UP KeyPadInput = iota
    DOWN
    LEFT
    RIGHT
    KEY_PAD_A
    START
)

const (
    ZERO NumPadInput = iota
    ONE
    TWO
    THREE
    FOUR
    FIVE
    SIX
    SEVEN
    EIGHT
    NINE
    NUM_PAD_A
)

type KeyPadVect struct{
  From KeyPadInput
  To KeyPadInput
}

var KeyPadLen = map[KeyPadVect]int{
  {UP,UP}: 1,
  {UP,DOWN}: 2,
  {UP,LEFT}: 3,
  {UP,RIGHT}: 3,
  {UP,KEY_PAD_A}: 2,
  {KEY_PAD_A,UP}: 2,
  {KEY_PAD_A,DOWN}: 3,
  {KEY_PAD_A,LEFT}: 4,
  {KEY_PAD_A,RIGHT}: 2,
  {KEY_PAD_A,KEY_PAD_A}: 1,
  {LEFT,UP}: 3,
  {LEFT,DOWN}: 2,
  {LEFT,LEFT}: 1,
  {LEFT,RIGHT}: 3,
  {LEFT,KEY_PAD_A}: 4,
  {DOWN,UP}: 2,
  {DOWN,DOWN}: 2,
  {DOWN,LEFT}: 2,
  {DOWN,RIGHT}: 2,
  {DOWN,KEY_PAD_A}: 3,
  {RIGHT,UP}: 3,
  {RIGHT,DOWN}: 2,
  {RIGHT,LEFT}: 3,
  {RIGHT,RIGHT}: 1,
  {RIGHT,KEY_PAD_A}: 2,
} 

var KeyPadPaths = map[KeyPadVect][]string{
  {UP,UP}: {},
  {UP,DOWN}: {"v"},
  {UP,LEFT}: {"v<"},
  {UP,RIGHT}: {"v>"},
  {UP,KEY_PAD_A}: {">"},
  {KEY_PAD_A,UP}: {"<"},
  {KEY_PAD_A,DOWN}: {"<v","v<"},
  {KEY_PAD_A,LEFT}: {"<v<","v<<"},  
  {KEY_PAD_A,RIGHT}: {"v"},
  {KEY_PAD_A,KEY_PAD_A}: {},
  {LEFT,UP}: {">^"},
  {LEFT,DOWN}: {">"},
  {LEFT,LEFT}: {},
  {LEFT,RIGHT}: {">>"},
  {LEFT,KEY_PAD_A}: {">>^",">^>"},
  {DOWN,UP}: {"^"},
  {DOWN,DOWN}: {},
  {DOWN,LEFT}: {"<"},
  {DOWN,RIGHT}: {">"},
  {DOWN,KEY_PAD_A}: {">^","^>"},
  {RIGHT,UP}: {"^<","<^"},
  {RIGHT,DOWN}: {"<"},
  {RIGHT,LEFT}: {"<<"},
  {RIGHT,RIGHT}: {},
  {RIGHT,KEY_PAD_A}: {"^"},
}

var keyPadInputName = map[KeyPadInput]string{
    UP:        "^",
    DOWN:      "v",
    LEFT:      "<",
    RIGHT:     ">",
    KEY_PAD_A: "A",
    START:     "X",
}

var numPadInputName = map[NumPadInput]string{
    ZERO:      "0",
    ONE:       "1",
    TWO:       "2",
    THREE:     "3",
    FOUR:      "4",
    FIVE:      "5",
    SIX:       "6",
    SEVEN:     "7",
    EIGHT:     "8",
    NINE:      "9",
    NUM_PAD_A: "A",
}

var numPadInputValue = map[string]NumPadInput{
    "0": ZERO,
    "1": ONE,
    "2": TWO,
    "3": THREE,
    "4": FOUR,
    "5": FIVE,
    "6": SIX,
    "7": SEVEN,
    "8": EIGHT,
    "9": NINE,
    "A": NUM_PAD_A,
}

var keyPadInputValue = map[string]KeyPadInput{
  "^" : UP,
  "v" : DOWN,
  "<": LEFT,
  ">": RIGHT,
  "A": KEY_PAD_A,
  "X" :START,
}

type NumPadMovement struct {
    Direction KeyPadInput
    Target    NumPadInput
}

var numPadNeighbors = map[NumPadInput][]NumPadMovement{
    ZERO: {
        {UP, TWO},
        {RIGHT, NUM_PAD_A},
    },
    ONE: {
        {UP, FOUR},
        {RIGHT, TWO},
    },
    TWO: {
        {LEFT, ONE},
        {UP, FIVE},
        {RIGHT, THREE},
        {DOWN, ZERO},
    },
    THREE: {
        {LEFT, TWO},
        {UP, SIX},
        {DOWN, NUM_PAD_A},
    },
    FOUR: {
        {UP, SEVEN},
        {RIGHT, FIVE},
        {DOWN, ONE},
    },
    FIVE: {
        {LEFT, FOUR},
        {UP, EIGHT},
        {RIGHT, SIX},
        {DOWN, TWO},
    },
    SIX: {
        {LEFT, FIVE},
        {UP, NINE},
        {DOWN, THREE},
    },
    SEVEN: {
        {RIGHT, EIGHT},
        {DOWN, FOUR},
    },
    EIGHT: {
        {LEFT, SEVEN},
        {RIGHT, NINE},
        {DOWN, FIVE},
    },
    NINE: {
        {LEFT, EIGHT},
        {DOWN, SIX},
    },
    NUM_PAD_A: {
        {LEFT, ZERO},
        {UP, THREE},
    },
}

func ExtractInput(buffer string)[]NumTarget{
  targets := []NumTarget{}
  lines := strings.Split(strings.Trim(buffer,"\n"),"\n")

  for i := range lines{
    content := lines[i]
    num_str := regexp.MustCompile("[0-9]{3}").FindString(content)
    coef,_ := strconv.Atoi(num_str)

    code := []NumPadInput{}
    inputs :=strings.Split(content,"")

    for i := range inputs {
      code = append(code,numPadInputValue[inputs[i]])
    }


    targets = append(targets,NumTarget{
      Code : code, 
      Coef : coef,
    })
  }
  return targets 
}

type NumTarget struct{
  Code []NumPadInput
  Coef int
}

type NumNode struct{
  Pos NumPadInput
  Score int
  Path string
} 

func (t NumTarget) FindMinPath(start NumNode, end NumPadInput)[]NumNode{
  nodes := []NumNode{start}
  min_score := math.MaxInt32
  paths := []NumNode{}

  for len(nodes) >0 {
    node := nodes[0]
    nodes = nodes[1:]

    if node.Score >min_score {continue}
    if node.Pos == end{
      if node.Score > min_score {break}
      min_score = node.Score
      paths = append(paths,node)
      continue
    }

    nb,_ := numPadNeighbors[node.Pos] 

    for i := range nb{
      mv,key := nb[i].Direction,nb[i].Target
      nodes = append(nodes,NumNode{
        Pos:key,
        Score:node.Score+1,
        Path: ""+ node.Path+keyPadInputName[mv],
      })
    }
  }
  return paths
}

func (t NumTarget) ComputeBestPath(paths *[]string, last *NumNode, idx int, code []NumPadInput){
  if idx >= len(code){
    *paths = append(*paths,last.Path)
    return
  }

  p_node := t.FindMinPath(*last,code[idx])

  for i := range p_node{
    node := NumNode{
      Pos: p_node[i].Pos,
      Score: p_node[i].Score,
      Path: ""+ p_node[i].Path + keyPadInputName[KEY_PAD_A],
    }
    t.ComputeBestPath(paths,&node,idx+1,code)
  }
}

