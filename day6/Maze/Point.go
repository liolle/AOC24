package maze

type Position struct {
  Row int
  Col int
  Direction int
}

func (p *Position) Add(other Position)  {
  p.Col += other.Col
  p.Row += other.Row
  p.Direction = other.Direction
}

func (p Position) GetAdd(other Position) Position  {
  return Position {
    Col:p.Col + other.Col,
    Row:p.Row + other.Row,
    Direction:other.Direction,
  }
}

func (p Position) Equals (other Position) bool {
  if p.Row == other.Row && p.Col == other.Col && p.Direction == other.Direction {return true}
  return false 
}

func (p Position) EqualsS (other Position) bool {
  if p.Row == other.Row && p.Col == other.Col  {return true}
  return false 
}
