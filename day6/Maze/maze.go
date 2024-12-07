package maze

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var DIRECTIONS = []Position{
  {Row:-1,Col:0, Direction:0},
  {Row:0,Col:1, Direction:1},
  {Row:1,Col:0, Direction:2},
  {Row:0,Col:-1, Direction:3},
}

type Maze struct {
  Map [][]string
  Rows int
  Cols int
  GuardMoves int
  VisitedCells int
  AdditionalObstacles int
  GuardPosition *Position
  GuardDirection int
  NewObstacles []Position
  StartPoint Position
  Path []Position
}

func NewMaze(buffer string) Maze{
  maze := [][]string{}
  guard_position := Position{Row:-1,Col:-1}
  guard_direction := 0
  lines := strings.Split(buffer,"\n")

  for _,line := range(lines[:len(lines)-1]){
    Row := strings.Split(line,"")
    maze = append(maze,Row)
  }

  Rows := len(maze)
  Cols := len(maze[0])

  for i := 0; i < Rows; i++ {
    for j := 0; j < Cols; j++ {
      switch maze[i][j] {
      case ">":
        guard_direction = 1
        guard_position = Position{Row:i,Col:j}
        //maze[i][j] = "2"
        break 
      case "<":
        guard_direction = 3
        guard_position = Position{Row:i,Col:j}
        //maze[i][j] = "8"
        break 
      case "^":
        guard_direction = 0
        guard_position = Position{Row:i,Col:j}
        maze[i][j] = "1"
        break 
      case "v":
        guard_direction = 2
        guard_position = Position{Row:i,Col:j}
        //maze[i][j] = "4"
        break
      case ".":
        maze[i][j] = " "
        break
      default:
        break
      }
    }
  }

  return Maze{
    Map: maze,
    Rows: Rows,
    Cols: Cols,
    GuardMoves :0,
    GuardPosition : &guard_position,
    GuardDirection : guard_direction,
    NewObstacles: []Position{},
    StartPoint: Position{guard_position.Row,guard_position.Col,0},
    Path: []Position{},
    
  }
}

func (maze Maze) InBound(pos Position,direction Position)bool{

  n_Row := pos.Row + direction.Row  
  n_Col := pos.Col + direction.Col  

  return n_Row < maze.Rows  && n_Row >= 0 && n_Col < maze.Cols  && n_Col >= 0  
}

func (maze Maze) IsObstacleNext() bool{
  n_Row := maze.GuardPosition.Row + DIRECTIONS[maze.GuardDirection].Row  
  n_Col := maze.GuardPosition.Col + DIRECTIONS[maze.GuardDirection].Col

  return maze.Map[n_Row][n_Col] == "#"
}

func (maze *Maze) Paint(){
  char := maze.Map[maze.GuardPosition.Row][maze.GuardPosition.Col] 
  cur,err := strconv.Atoi(char)
  if err != nil {cur = 0}
  
  dir_code := int(math.Pow(2,float64(maze.GuardDirection)))
  maze.Map[maze.GuardPosition.Row][maze.GuardPosition.Col] = strconv.Itoa(dir_code | cur)
}

func (maze *Maze) Walk()bool{
  if !maze.InBound(*maze.GuardPosition,DIRECTIONS[maze.GuardDirection]){
    return false
  }

  if maze.IsObstacleNext() {
    maze.GuardDirection = (maze.GuardDirection + 1)%4
    return maze.Walk()
  }

  maze.GuardMoves += 1
  maze.GuardPosition.Add(DIRECTIONS[maze.GuardDirection])

  regex := regexp.MustCompile("[^0-9]+")
  char := maze.Map[maze.GuardPosition.Row][maze.GuardPosition.Col] 

  if regex.MatchString(char)  {maze.VisitedCells += 1} 
  maze.Paint()
  return true 
}

func (maze *Maze) fillX(dir bool){
  r := maze.GuardPosition.Row
  pivot :=  maze.GuardPosition.Col

  for i := pivot; i >= 0 && dir; i-- {
      char := maze.Map[r][i] 
      if char == "#" {break}
      val,_ := strconv.Atoi(char) 
      maze.Map[r][i] = strconv.Itoa(val | 1) 
  }

  for i := pivot; i < maze.Cols && !dir; i++ {
      char := maze.Map[r][i] 
      if char == "#" {break}
      val,_ := strconv.Atoi(char) 
      maze.Map[r][i] = strconv.Itoa(val | 4) 
  }
}

func (maze *Maze) fillY(dir bool){
  c := maze.GuardPosition.Col
  pivot :=  maze.GuardPosition.Row

  for i := pivot; i >= 0 && dir; i-- {
    char := maze.Map[i][c] 
      if char == "#" {break}
    val,_ := strconv.Atoi(char) 
    maze.Map[i][c] = strconv.Itoa(val | 2) 
  }

  for i := pivot; i < maze.Rows && ! dir; i++ {
    char := maze.Map[i][c] 
      if char == "#" {break}
    val,_ := strconv.Atoi(char) 
    maze.Map[i][c] = strconv.Itoa(val | 8) 
  }
}

func Includes(slice []Position, value Position) bool {
  for _,v := range(slice) {
    if v.Equals(value) {
      return true
    }
  }
  return false
}

func IncludesS(slice []Position, value Position) bool {
  for _,v := range(slice) {
    if v.EqualsS(value) {
      return true
    }
  }
  return false
}

func (maze *Maze) chekValid(row int,col int) bool{
  n_mask := int(math.Pow(2,float64((maze.GuardDirection +1)%4)))  

  char := maze.Map[row][col] 
  cur,err := strconv.Atoi(char)
  if err != nil {cur = 0}
  if (n_mask & cur) == n_mask {
    maze.NewObstacles = append(maze.NewObstacles,maze.GuardPosition.GetAdd(DIRECTIONS[maze.GuardDirection]))
    return true
  }
  return false
}

func (maze *Maze) Scan(){

  mask := int(math.Pow(2,float64(maze.GuardDirection))) 
  row := maze.GuardPosition.Row 
  col := maze.GuardPosition.Col 

  char := maze.Map[row][col] 
  cur,err := strconv.Atoi(char)
  if err != nil {cur = 0}
  maze.Map[row][col] = strconv.Itoa(cur | mask) 

  switch mask {
  case 1:
    for i := col; i < maze.Cols; i++ {
     if maze.chekValid(row,i) {break}
    }
    break

  case 2:
    for i := row; i < maze.Rows; i++ {
      if maze.chekValid(i,col) {break}
    }
    break

  case 4:
    for i := col; i >= 0; i-- {
      if maze.chekValid(row,i) {break}
    }

    break
  case 8:
    for i := row; i >= 0; i-- {
      if maze.chekValid(i,col) {break}
    }
    break
  }
}

func (maze *Maze) Travel(path *[]Position, cycle *bool)bool{

  //maze.Scan()

  if !maze.InBound(*maze.GuardPosition,DIRECTIONS[maze.GuardDirection]){
      *path = append(*path,*maze.GuardPosition)
    return false
  }

  if maze.IsObstacleNext() {
    maze.GuardDirection = (maze.GuardDirection + 1)%4
    return maze.Travel(path,cycle)
  }

  
  if Includes(*path,*maze.GuardPosition) {*cycle = true}
  *path = append(*path,*maze.GuardPosition)
  maze.GuardMoves += 1
  maze.GuardPosition.Add(DIRECTIONS[maze.GuardDirection])

  return true 
} 

func (maze *Maze) Print(){
  pattern := `[^0-9]+`
  regex := regexp.MustCompile(pattern)
  for _,Row := range(maze.Map){
    for _,char := range(Row){
      if char == "#"{
        fmt.Print(" \033[31m"+char)
      }else if regex.MatchString(char)  {
        fmt.Print(" \033[32m"+char)
      }else {
        fmt.Print(" \033[0m"+char)
      }
    }
    fmt.Println()
  }
}

func (maze *Maze) PrintP2(){
  pattern := `[^#]`
  regex := regexp.MustCompile(pattern)
  for i := 0; i < maze.Rows; i++ {
    for j := 0; j < maze.Rows; j++ {
      if i == maze.StartPoint.Row && j == maze.StartPoint.Col {
        maze.Map[i][j] = "1"
        continue
      }
      if IncludesS(maze.NewObstacles, Position{i,j,0}) {
        maze.Map[i][j] = "O"
        continue
      } 
      char := maze.Map[i][j] 

      if regex.MatchString(char) {maze.Map[i][j] = " "}
    }
  }
  maze.Print()
}

func (maze *Maze) Clear(){
  pattern := `[^#]`
  regex := regexp.MustCompile(pattern)
  for i := 0; i < maze.Rows; i++ {
    for j := 0; j < maze.Rows; j++ {
      if i == maze.StartPoint.Row && j == maze.StartPoint.Col {
        maze.Map[i][j] = "1"
        continue
      }
       
      char := maze.Map[i][j] 
      if regex.MatchString(char) {maze.Map[i][j] = " "}
    }

  }
  maze.GuardPosition.Row = maze.StartPoint.Row 
  maze.GuardPosition.Col = maze.StartPoint.Col
  maze.GuardPosition.Direction = 0
  maze.GuardDirection = 0
}

func (maze *Maze) AddObstacle(row int,col int){
  maze.Map[row][col] = "#"
}

func (maze *Maze) RemoveObstacle(row int,col int){
  maze.Map[row][col] = " "
}

