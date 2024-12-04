package main

import (
	"aoc24/stack"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func part1(input [][]int) int{
  count := 0
  for i:=0;i<len(input);i++{
    add := 1
    slope := "asc"
    for j := 1; j < len(input[i]); j++ {
      diff := math.Abs(float64(input[i][j]) - float64(input[i][j-1]))
      if(diff <1 || diff >3){
        add = 0
        break
      }

      if(j == 1){
        if(input[i][j]<input[i][j-1]){
          slope = "desc"
        }
        continue
      }

      if( 
      (slope == "asc" && input[i][j]<input[i][j-1] ) ||
      (slope == "desc" && input[i][j]>input[i][j-1] ) ){
        add = 0
        break
      }
    } 
    count += add
  }
  return count
}

func valid(left int, right int, arr []int,diffs stack.Stack[int])bool{
  if(left >= len(arr) || right > len(arr)){return false}
  diff := arr[left] - arr[right]
  last_diff,err := diffs.Peek() 
  if(err != nil){
    return math.Abs(float64(diff))>=1 && math.Abs(float64(diff))<=3 
  }
  if(diff<0 && last_diff >0 || diff>0 && last_diff <0){return false}
  if(math.Abs(float64(diff))<1 || math.Abs(float64(diff))>3){return false}
  return true 
}

func validate(left int, right int, arr []int, diffs stack.Stack[int],limit int,route stack.Stack[[]int]) bool{
  if(limit <0){return false}
  if(right>=len(arr) || left >= len(arr)-1){
    return true
  }
  n_diffs := diffs.Copy()
  r := route.Copy()
  pt := []int{left,right}
  r.Push(pt)
  v := valid(left,right,arr,diffs) 
  diff := arr[left] - arr[right]
  if(v){
    n_diffs.Push(diff) 
    return validate(right,right+1,arr,n_diffs,limit,r)
  }else {
    // remove left 
   nn_diffs := diffs.Copy()
    nn_diffs.Pop()
    next := false
          next = validate(left,right,arr,nn_diffs,limit-1,r)

   
  
    if(left>0){
      rd := n_diffs.Copy()
      n_diffs.Pop()
      next = next || validate(left-1,right,arr,n_diffs,limit-1,r) || validate(left,right+1,arr,rd,limit-1,r) 
    }else {
      next = next || validate(right,right+1,arr,n_diffs,limit-1,r) || validate(left,right+1,arr,n_diffs,limit-1,r)

    }
    return next
  }
}

func part2(input [][]int,limit int) int{
  count := 0
  st := stack.Stack[int]{} 
  sr := stack.Stack[[]int]{} 
  for i:=0;i<len(input);i++{
    v := validate(0,1,input[i],st,limit,sr) 
    if(v){count++}
  }
  return count
}

func main()  {
  file, err := os.Open("input1.txt")
  //file, err := os.Open("sample.txt")
  //file, err := os.Open("edge1.txt")
  if err != nil {
    return
  }
  defer file.Close()
  buffer := bufio.NewReader(file)
  input := [][]int{} 

  for {
    line,_,err := buffer.ReadLine()

    if(err != nil){
      break
    }

    p := strings.Split(string(line)," ")

    res := []int{}

    for _,num := range p{
      val,err := strconv.Atoi(num)
      if(err != nil){
        continue
      }
      res = append(res,val)
    }
    input = append(input,res)
  }
  fmt.Println(part1(input))
  fmt.Println(part2(input,1))
}
