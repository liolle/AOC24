package Point

type Point struct {
  Row int
  Col int
  Direction string
}

func (p Point) Equals (other Point) bool {
  if p.Row == other.Row && p.Col == other.Col {return true}
  return false 
}

func (p Point) Outbound (Rows int, Cols int) bool {
  return p.Row <0 || p.Row>= Cols ||  p.Col <0 || p.Col>= Rows 
}


func (p Point) GetAdd (other Point) Point {
  return Point{Row: p.Row+other.Row, Col:p.Col +other.Col,Direction:other.Direction} 
}


