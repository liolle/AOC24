package main

import (
	"regexp"
	"slices"
	"strconv"
	"strings"
)


func ExtractInput(buffer string)Computer{
  n_str :=regexp.MustCompile("Program: ((,*[0-9])+)").FindStringSubmatch(buffer)[1] 
  num_str := strings.Split(n_str,",")
  instructions := []Instruction{}
  nums := []int{}

  for i :=range num_str {
    num,_ := strconv.Atoi(string(num_str[i]))
    nums = append(nums,num)
  }
  
  for i :=0;i<len(num_str)-1;i+=2{
    opcode,_ := strconv.Atoi(num_str[i])
    operand,_ := strconv.Atoi(num_str[i+1])
    instructions = append(instructions,Instruction{Opcode:opcode,Operand:operand})
  }

  a,_:= strconv.Atoi(regexp.MustCompile("Register A: ([0-9]+)").FindStringSubmatch(buffer)[1])
  b,_:= strconv.Atoi(regexp.MustCompile("Register B: ([0-9]+)").FindStringSubmatch(buffer)[1])
  c,_:= strconv.Atoi(regexp.MustCompile("Register C: ([0-9]+)").FindStringSubmatch(buffer)[1])

  return Computer{
    A:a,
    B:b,
    C:c,
    InstructionPointer:0,
    Instructions:instructions,
    OutputBuffer:[]int{},
    Target:nums,
  }
}


type Computer struct {
  A int 
  B int 
  C int
  InstructionPointer int  
  Instructions []Instruction
  OutputBuffer []int
  Target []int
}

type Instruction struct {
  Opcode int
  Operand int 
}

func (c *Computer) reset(B int, C int){
  c.InstructionPointer = 0
  c.A = 0
  c.B = B
  c.C = C
  c.OutputBuffer =[]int{} 
}

func padLeft(s string, padChar string, totalLength int) string {
	if len(s) >= totalLength {return s }
	padding := strings.Repeat(padChar, totalLength-len(s))
	return padding + s
}

func Chunk(s string)[]string{
  chunks := []string{}

  for i:=len(s)-1;i>=0;i-=3 {
    chunk := ""
    for j:=i;j>i-3 && j>=0;j--{
      chunk += string(s[j])
    }
    chunks = append(chunks,chunk)
  }
   slices.Reverse(chunks)
  return chunks 
}

func toBits(val int)string{
  return padLeft(strconv.FormatInt(int64(val),2),"0",32 )
}

func (c Computer) Check(idx int)bool{
  if c.Target[idx] != c.OutputBuffer[0] {return false} 
  return true
}

func (c *Computer) Compute() {
  for c.InstructionPointer < len(c.Instructions){
    c.ComputeInstruction(c.Instructions[c.InstructionPointer])
  }
}

func convertBitArray(arr []int)int{
  num := 0
  j := 0
  for i:= len(arr)-1;i>=0;i--{
    num = num + arr[i]<<(3*j)
    j++
  }
  return num
}

func (c *Computer) ComputeInstruction(instruction Instruction){
  switch instruction.Opcode {
  case 0:
    c.A = c.A>>c.getOperand(instruction.Operand)   
    break
  case 1:
    c.B = c.B^instruction.Operand
    break
  case 2:
    c.B = c.getOperand(instruction.Operand)%8
    break
  case 3:
    if(c.A != 0){
      c.InstructionPointer = instruction.Operand
      return
    }
    break
  case 4:
    c.B = c.B^c.C
    break
  case 5:
    c.PushToBuffer(c.getOperand(instruction.Operand)%8)
    break
  case 6:
    c.B = c.A>>c.getOperand(instruction.Operand)   
    break
  case 7:
    c.C = c.A>>c.getOperand(instruction.Operand)   
    break
  default:
    break
  }
  c.InstructionPointer++
}

func (c *Computer) getOperand(operand int)int{
  switch operand {
  case 0:
    return operand
  case 1:
    return operand
  case 2:
    return operand
  case 3:
    return operand
  case 4:
    return c.A
  case 5:
    return c.B 
  case 6:
    return c.C 
  default:
    break
  }
  panic("Invalid operand: "+strconv.Itoa(operand))
}

func (c *Computer) PushToBuffer(val int){
  c.OutputBuffer = append(c.OutputBuffer,val)
}
