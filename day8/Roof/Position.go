package roof

type Position struct {
  X int
  Y int
}

func (p *Position) Add(other Position)  {
  p.X += other.X
  p.Y += other.Y
}

func (p Position) GetAdd(other Position) Position  {
  return Position {
    X:p.X + other.X,
    Y:p.Y + other.Y,
  }
}

func (p Position) Equals (other Position) bool {
  if p.Y == other.Y && p.X == other.X {return true}
  return false 
}

func (p Position) getDouble(other Position)Position{

  vect := Position{
    X: 2*(other.X - p.X),
    Y: 2*(other.Y - p.Y),
  }

  return p.GetAdd(vect) 
}

func (p Position) getVect(other Position)Position{

  vect := Position{
    X: (other.X - p.X),
    Y: (other.Y - p.Y),
  }

  return vect 
}
