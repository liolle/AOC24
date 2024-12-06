package main

import (
  "fmt"
  "os"
  "strconv"
  "strings"
)

func extractInput(buffer string, rules *map[int][]int, pages *[][]int)  {
  parts := strings.Split(buffer,"\n\n")
  r := strings.Split(parts[0],"\n")
  p :=  strings.Split(parts[1],"\n") 
  for _,val := range(r){
    lr := strings.Split(val,"|")
    if len(lr) != 2 {continue}
    left,l_err := strconv.Atoi(lr[0])
    right,r_err := strconv.Atoi(lr[1]) 
    if l_err != nil || r_err != nil {continue}

    // !! arr is a copied 
    if arr,exist := (*rules)[right]; exist{
      (*rules)[right] = append(arr,left) 
    }else {
      (*rules)[right] = []int{left}
    }
  }

  for _,val := range(p[:len(p)-1]){
    str_nums := strings.Split(val,",")
    nums := []int{}
    for _,str_n := range(str_nums){
      n,err := strconv.Atoi(str_n)
      if err != nil {continue}
      nums = append(nums,n)
    }
    (*pages) = append((*pages),nums)
  }
}

func Includes(slice []int, value int) bool {
  for _,v := range(slice) {
    if v == value {return true}
  }
  return false
}

func orderR(page int,rules map[int][]int,ordered_pages *[]int, pages []int,moves *int){
  if Includes(*ordered_pages,page) {return}
  for _,val := range(rules[page]){
    if !Includes(pages,val) || Includes(*ordered_pages,val){continue}
    *moves++
    orderR(val,rules,ordered_pages,pages,moves) 
  }
  (*ordered_pages) = append((*ordered_pages),page)
}

func order(pages []int, rules map[int][]int,moves *int)[]int{
  ordered_pages := []int{}

  for _,page := range(pages) {
    orderR(page,rules,&ordered_pages,pages,moves) 
  }
  return ordered_pages
}

func part1(rules map[int][]int, book *[][]int)int {
  sum := 0
  for _,pages := range(*book){
    moves := 0
    or_pages := order(pages,rules,&moves)
    if moves == 0{
      sum += or_pages[len(or_pages)/2]
    }
  }
  return sum
}

func part2(rules map[int][]int, book *[][]int)int {
  sum := 0
  for _,pages := range(*book){
    moves := 0
    or_pages := order(pages,rules,&moves)
    if moves > 0{
      sum += or_pages[len(or_pages)/2]
    }
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

  rules :=  map[int][]int{}
  book := [][]int{}
  extractInput(string(buffer),&rules,&book)

  fmt.Println(part1(rules,&book))
  fmt.Println(part2(rules,&book))
}
