package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cache = map[string]int{} 

func ExtractInput(buffer string) []int{
  nums := []int{}
  str_nums := strings.Split(buffer," ")
  for _,str_num := range(str_nums){
    num,_ := strconv.Atoi(str_num)
    nums = append(nums,num)
  }
  return nums
}

func TrimNumber(str_num string)int{
  i :=0 
  for _,char := range(str_num) {
    if char != '0' {break}
    i++
  }
  ret,err := strconv.Atoi(str_num[i:])
  if err != nil {ret =0} 
  return ret 
}

func Count(value int,blinks int) int{
  key := strconv.Itoa(value) +":"+ strconv.Itoa(blinks)
  cached_value,exist := cache[key]
  
  if exist {return cached_value}

  if blinks == 0 { 
    return 1 
  }

  str_value := strconv.Itoa(value)
  cnt := 0
  if len(str_value)%2 ==0{
    left,_ := strconv.Atoi(str_value[0:len(str_value)/2])
    right := TrimNumber(str_value[len(str_value)/2:])
    cnt = Count(left,blinks-1) + Count(right,blinks-1) 
  }else if value == 0 {
    cnt = Count(1,blinks-1) 
  }else{
    cnt = Count(value*2024,blinks-1) 
  }
  cache[key] = cnt
  return cnt
}

func Blink(nums []int)int{
  sum := 0
  for _,num := range(nums){
    sum += Count(num,75)
  }
  return sum
}

func part1(nums []int)int {
  sum := 0
  for _,num := range(nums){
    sum += Count(num,25)
  }
  return sum
}

func part2(nums []int)int {
  sum := 0
  for _,num := range(nums){
    sum += Count(num,75)
  }
  return sum
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

  p1 := ExtractInput(strings.Trim(string(buffer),"\n"))
  fmt.Println(part1(p1))
  fmt.Println(part2(p1))
}
