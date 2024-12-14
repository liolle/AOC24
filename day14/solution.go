package main

import (
  "fmt"
  "os"
  // "image"
  //"image/color"
  //"image/png"
  "regexp"
  "strconv"
  "strings"
)

func ExtractInput(buffer string)[]Robot{

  out := []Robot{}
  lines := strings.Split(buffer,"\n")

  for _,line := range lines {
    px,_ := strconv.Atoi(regexp.MustCompile("p=(-{0,1}[0-9]+)").FindStringSubmatch(line)[1])
    py,_ := strconv.Atoi(regexp.MustCompile("p=-{0,1}[0-9]+,(-{0,1}[0-9]+)").FindStringSubmatch(line)[1])

    vx,_ := strconv.Atoi(regexp.MustCompile("v=(-{0,1}[0-9]+)").FindStringSubmatch(line)[1])
    vy,_ := strconv.Atoi(regexp.MustCompile("v=-{0,1}[0-9]+,(-{0,1}[0-9]+)").FindStringSubmatch(line)[1])

    rb :=Robot{
      Start:Vect{py,px},
      End: Vect{py,px},
      Velocity:Vect{vy,vx},
      Quadrant: -1,
      Rows: 103,
      Cols: 101,
    } 

    out = append(out,rb )
  }

  return out
}

type Vect struct {
  Row int
  Col int 
}

type Robot struct{
  Start Vect
  End Vect
  Velocity Vect
  Quadrant int 
  Rows int
  Cols int
}

type RobotNet []Robot

type Display struct {
  Image[][]int
  Product int
}

func (r *Robot) Normalise(){
  if (*r).End.Row <0 {(*r).End.Row = r.Rows + (*r).End.Row } 
  if (*r).End.Col <0 {(*r).End.Col = r.Cols + (*r).End.Col } 
}

func (r *Robot) AssignQuadrant(){
  row,col := r.End.Row, r.End.Col 
  rm := r.Rows /2
  cm := r.Cols /2
  if row>=0 && row< rm && col>=0 && col< cm {
    r.Quadrant = 0
  }else if row>=0 && row< rm && col>cm && col< r.Cols{
    r.Quadrant = 1
  }else if row>rm && row< r.Rows && col>=0 && col< cm{
    r.Quadrant = 2
  }else if row>rm && row< r.Rows && col>cm && col< r.Cols{
    r.Quadrant = 3
  }else {
    r.Quadrant = -1
  } 
}

func (r *Robot) Translate(scal int){
  r.End.Row = (r.End.Row + r.Velocity.Row *scal)% r.Rows
  r.End.Col = (r.End.Col +r.Velocity.Col *scal)%r.Cols
  r.Normalise()
  r.AssignQuadrant()
}

func part1(robots RobotNet)int {

  q := [4]int{0,0,0,0}

  robots.Translate(100)
  for _,r := range robots {
    if r.Quadrant >=0 {
      q[r.Quadrant]++
    }
  }

  prod := 1
  fmt.Println(q)
  for _,cnt := range q {
    prod *= cnt
  }

  return prod
}

func (r *RobotNet) Translate( n int){
  // second arg <elem> of the for i,elm := range(arr) is alway a copy
  for i := range *r {
    (*r)[i].Translate(n)
  }
}

func (r *RobotNet) ComputeDisplay()Display{
  rows,cols := (*r)[0].Rows, (*r)[0].Cols
  q := [4]int{0,0,0,0}

  image := make([][]int,rows) 

  for i := range image {
    image[i] = make([]int, cols)
  }

  for i := 0; i < rows; i++ {
    for j := 0; j < cols; j++ {
      image[i][j] = 0 
    }
  }

  for _,rb := range *r {
    if rb.Quadrant >=0 {
      q[rb.Quadrant]++
    }
    image[rb.End.Row][rb.End.Col]++
  }

  prod := 1
  for _,cnt := range q {
    prod *= cnt
  }


  return Display{
    Image: image,
    Product:prod, 
  }
}

func (d *Display) Print(){

  for _,row := range d.Image{
    for _,num := range row {
      if num >0 {
        fmt.Print("#")
      }else{
        fmt.Print(" ")
      } 
    }
    fmt.Print("\n")
  }
  fmt.Println(d.Product)

}

func part2(robots RobotNet)int {
  /*
displays := []Display{}

for i := 0; i < 10000; i++ {
robots.Translate(1)
displays = append(displays,robots.ComputeDisplay())  
}

for i := range displays {
rows,cols := len(displays[i].Image), len(displays[i].Image[0]) 
img := image.NewGray(image.Rect(0, 0, rows, cols))

fileName := "result/output_" + strconv.Itoa(i) + ".png"

file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
if err != nil {
fmt.Println("Error creating or opening file:", err)
}
defer file.Close()

for y := 0; y < rows; y++ {
for x := 0; x < cols; x++ {
if displays[i].Image[y][x] > 0 {
img.SetGray(x, y, color.Gray{Y: 0}) // Black pixel for > 0
} else {
img.SetGray(x, y, color.Gray{Y: 255}) // White pixel for 0
}
}
}

defer file.Close()

err = png.Encode(file, img)
if err != nil {
panic(err)
}

}
*/
  robots.Translate(7037)
  d := robots.ComputeDisplay()
  d.Print()

  return 0
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

  r1 := ExtractInput(strings.Trim(string(buffer),"\n"))
  r2 := ExtractInput(strings.Trim(string(buffer),"\n"))

  fmt.Println(part1(r1))
  fmt.Println(part2(r2))
}
