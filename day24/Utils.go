package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func ExtractInput(buffer string) (out Board){
  parts := strings.Split(strings.Trim(buffer,"\n"),"\n\n")
  out.Wires = map[string]*Gate{}
  p1_lines := strings.Split(parts[0],"\n") 
  p2_lines := strings.Split(parts[1],"\n") 

  for _,line := range p1_lines  {
    gt := strings.Split(line,": ")
    value,_ := strconv.Atoi(string(gt[1])) 
    name := string(gt[0])
    out.Wires[name] = &Gate{value,nil,nil,AND,true}
  }

  gates_plan := map[string]PGate{}
  for _,line := range p2_lines  {
    in1 := regexp.MustCompile("^...").FindString(line)
    in2 := regexp.MustCompile("(AND|OR|XOR) (...)").FindStringSubmatch(line)[2]
    out := regexp.MustCompile("...$").FindString(line)
    gtype := regexp.MustCompile("(AND|OR|XOR)").FindString(line) 
    gates_plan[out] = PGate{out,in1,in2,GTypeFromString[gtype]}
  }
  out.Plan = gates_plan

  return out
}

type GATE_TYPE int
type GATE_FUNC func(int,int) int

const (
  AND GATE_TYPE = iota 
  OR
  XOR
  NIL
)

var GTypeToString = map[GATE_TYPE]string{
  AND : "AND",
  OR : "OR",
  XOR : "XOR",
}

var GTypeFromString = map[string]GATE_TYPE {
  "xor" : XOR,
  "XOR" : XOR,
  "or" : OR,
  "OR" : OR,
  "and" : AND,
  "AND" : AND,
}

type GATE_OUTPUT struct {
  Target string
  Value int
}

type Gate struct {
  Out int
  In1 *Gate
  In2 *Gate
  Type GATE_TYPE 
  Active bool
}

type PGate struct {
  Out string
  In1 string
  In2 string
  Type GATE_TYPE
}

func (g *Gate) Activate(){
  if g.Active {return}
  g.In1.Activate()
  g.In2.Activate()

  switch g.Type {

  case AND:
    g.Out = g.In1.Out & g.In2.Out 
    break
  case OR:
    g.Out = g.In1.Out | g.In2.Out 
    break
  case XOR:
    g.Out = g.In1.Out ^ g.In2.Out 
    break
  }
  g.Active = true
}

type Board struct {
  Wires map[string]*Gate
  Plan map[string]PGate 
}

func (b *Board) Build(gates map[string]PGate){
  if len(gates) == 0 {return}

  for key := range gates {
    in1,has_in1 := b.Wires[gates[key].In1]
    in2,has_in2 := b.Wires[gates[key].In2]

    if !has_in1 || !has_in2 {continue}

    b.Wires[gates[key].Out] = &Gate{
      Out :0,
      In1: in1,
      In2: in2,
      Active:false,
      Type: gates[key].Type,
    } 
    delete(gates,key)
  }
  b.Build(gates)
}

func (b *Board) TurnOn(){
  for key := range b.Wires {
    if !b.Wires[key].Active  {
      (*b).Wires[key].Activate()
    }
  }
}

func (b *Board) Count() (out int){
  for key := range b.Wires {
    gate := b.Wires[key]
    if gate.Out != 1 {continue}

    if regexp.MustCompile("z[0-9]{2,}").MatchString(key){
      val_str := regexp.MustCompile("[0-9]{2,}").FindString(key)
      val,_ := strconv.Atoi(val_str) 
      out +=int(math.Pow(2,float64(val))) 
    }
  }
  return out
}


func (b Board) RValide(gate string, prev *PGate)(out bool,blame string,reason string){
  reason = ""

  p_gate,has := b.Plan[gate]
  left := p_gate.In1
  right := p_gate.In1

  if prev == nil {
    lf,blamel,l_reason :=b.RValide(left,&p_gate)
    rg,blamer,r_reason :=b.RValide(right,&p_gate) 
    reason +=l_reason 
    reason += " "
    reason += r_reason
    blame += blamel + blamer
    return  lf&&rg,blame, reason
  }

  gtype := p_gate.Type

  //base →
  if !has {
    if prev.Type != AND && prev.Type != XOR {
      reason = fmt.Sprintf("Base input can only be called by AND/XOR")
      reason = fmt.Sprintf(prev.Out)  
      blame += prev.Out
      return false, blame, reason
    }
    
    if regexp.MustCompile("z[0-9]{2}").MatchString(prev.Out) && !regexp.MustCompile("z00").MatchString(prev.Out){
      reason = fmt.Sprintf("Z gates can't directly call base input")
      blame = fmt.Sprintf(prev.Out)  
      return false, blame, reason
    }
    return true,blame,reason
  }

  // AND →
  if (gtype == AND){
    if prev.Type != OR {
      reason = fmt.Sprintf("AND can only be called by a OR gate")
      blame = fmt.Sprintf(gate)  
      return false,blame, reason
    }
  }

  // AND ← 
  if prev.Type == AND {
    if gtype == AND {
      reason = fmt.Sprintf("AND can't call a AND") 
      blame = fmt.Sprintf(gate)  
      return false,blame, reason
    }
  }

  // XOR → 
  if (gtype == XOR){
    if prev.Type != XOR && prev.Type != AND{
      blame = fmt.Sprintf(gate)  
      reason = fmt.Sprintf("XOR can only be called by a AND/XOR") 
      return false,blame,reason
    }
  }

  // XOR ← 
  if prev.Type == XOR {
    if gtype == AND {
      reason = fmt.Sprintf("XOR can't call a AND") 
      blame = fmt.Sprintf(prev.Out)  
      return false,blame,reason
    }

    is_output := regexp.MustCompile("z[0-9]{2}").MatchString(prev.Out)

    lf := b.Plan[prev.In1]
    rt := b.Plan[prev.In2]
    if lf.Type == OR && rt.Type == XOR && !is_output {
      reason = fmt.Sprintf("(%s XOR %s) XOR (%s OR %s) → can only have a Z as output",rt.In1,rt.In2,lf.In1,lf.In2)  
      blame = fmt.Sprintf(prev.Out)  
      return false,blame,reason
    }

    if lf.Type == XOR && rt.Type == OR && !is_output {
      reason = fmt.Sprintf("(%s XOR %s) XOR (%s OR %s) → can only have a Z as output",lf.In1,lf.In2,rt.In1,rt.In2)  
      blame = fmt.Sprintf(prev.Out)  
      return false,blame,reason
    }
  }

  // OR →
  if gtype == OR {
    if prev.Type != AND && prev.Type != XOR {
      reason = fmt.Sprintf("OR gate can only be called by a AND/XOR")  
      blame = fmt.Sprintf(prev.Out)  
      return false,blame,reason
    }
  }

  // OR ← 
  if prev.Type == OR {
    if gtype != AND {
      reason = fmt.Sprintf("OR gate can only call AND")  
      blame = fmt.Sprintf(gate)  
      return false,blame,reason
    }
  }

  if prev.Type != XOR && !regexp.MustCompile("[a-z]{3}").MatchString(prev.Out) && !regexp.MustCompile("z45").MatchString(prev.Out){
    reason = fmt.Sprintf("AND/OR gate should alway output an intermatiade value exept for z45")  
    blame = fmt.Sprintf(prev.Out)  
    return false,blame,reason 
  } 
  return true,blame,reason
}

func (b Board) Valide(gate string)(out bool, blame string,reason string){
  valid,blame,reason := b.RValide(gate,nil) 
  if !valid {
    blame = strings.Trim(string(blame[0:len(blame)/2])," ") 
    reason = strings.Trim(string(reason[0:len(reason)/2])," ") 
    reason = fmt.Sprintf("%s\n", reason) 
    gt := b.Plan[gate]
    fmt.Printf("%s %s %s → %s\n%s\n", gt.In1, GTypeToString[gt.Type], gt.In2, gt.Out, reason)
  } 
  return valid,blame,reason 
}
