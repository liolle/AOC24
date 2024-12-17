package main

type Vect struct {
  Row int
  Col int 
}

func (v *Vect) Add(other Vect)Vect{
  return Vect{v.Row +other.Row, v.Col +other.Col}
}

func (v *Vect) Equals(other Vect)bool{
  return v.Row == other.Row && v.Col  == other.Col
}

func (v *Vect) AddS(other Vect){
  v.Row +=other.Row 
  v.Col +=other.Col
}

func (v *Vect) Set(other Vect){
  v.Row = other.Row
  v.Col = other.Col
}
