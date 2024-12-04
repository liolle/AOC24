package mult

type Mult struct  {
  Left int
  Right int 
}

func (m *Mult) Mult() int {
  return m.Left * m.Right
}
