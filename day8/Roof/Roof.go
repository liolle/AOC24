package roof

import "fmt"


type Roof struct {
  Frequencys map[string][]Position
  Antennas []Position
  Rows int
  Cols int
  Antinodes []Position
}

func (roof Roof) InBoundPoint(p Position) bool{
  return p.X >=0 && p.Y >= 0 && p.X<roof.Cols && p.Y<roof.Rows
}

func (roof Roof) InBound(x int, y int) bool{
  return x >=0 && y >= 0 || x<roof.Cols && y<roof.Rows
}

func IncludesPosition(slice []Position, value Position) bool {
  for _,v := range(slice) {
    if v.Equals(value) {
      return true
    }
  }
  return false
}

func Includes(slice []Position, x int, y int) bool {
  for _,v := range(slice) {
    if v.X == x && v.Y == y {
      return true
    }
  }
  return false
}

func (roof *Roof) ComputeAntinodes(){
  for _,frequency := range(roof.Frequencys){
    for i := 0; i < len(frequency); i++ {
      for j := 0; j < len(frequency); j++ {

        if i == j {continue}

        a1 := frequency[i].getDouble(frequency[j]) 

        if roof.InBoundPoint(a1) && !IncludesPosition(roof.Antinodes,a1)  {
          roof.Antinodes = append(roof.Antinodes,a1)
        }

      }
    }
  }
} 

func (roof *Roof) ComputeAntinodesHamonic(){
  for _,frequency := range(roof.Frequencys){
    for i := 0; i < len(frequency); i++ {
      for j := 0; j < len(frequency); j++ {

        if i == j {continue}

        v1 := frequency[i].getVect(frequency[j]) 
        a1 := Position{frequency[i].X,frequency[i].Y} 

        a1.Add(v1)

        for roof.InBoundPoint(a1) {
          if !IncludesPosition(roof.Antinodes,a1) && !IncludesPosition(roof.Antennas,a1){
            roof.Antinodes = append(roof.Antinodes,a1)
          }
          a1.Add(v1)
        }

      }
    }
  }
} 


func (roof Roof) Print()  {

  for i := 0; i < roof.Rows; i++ {
    for j := 0; j < roof.Cols; j++ {
      if Includes(roof.Antinodes,j,i){
        fmt.Print("#")
      }else {
        fmt.Print(".")
      }
    }
    fmt.Print("\n")
  }
}

