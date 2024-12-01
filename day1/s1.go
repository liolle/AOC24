package main 

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
    "sort"
    "math"
)
// get input into two list
// sort the input 
// iterate on the two list to get the sum of diff idx 
func main()  {

  file, err := os.Open("input1.txt")
  if err != nil {
    return
  }

  gp1 := []int{}
  gp2 := []int{}

  defer file.Close()
  
  buffer := bufio.NewReader(file)

  for {
    line, _, err := buffer.ReadLine()
    if len(line) > 0 {
      parts := strings.Split(string(line),"   ")
      num1, err1 := strconv.Atoi(parts[0])
      num2, err2 := strconv.Atoi(parts[1])

      if err1 != nil || err2 != nil {
        continue
      }

      gp1 = append(gp1,num1)
      gp2 = append(gp2,num2)
    }
    if err != nil {
      break
    }
  }
  sort.Ints(gp1)
  sort.Ints(gp2)


  sum_diff := 0
  for i := 0; i < len(gp1); i++ {
   sum_diff += int(math.Abs(float64(gp1[i])-float64(gp2[i])))
  }
  
  fmt.Println(sum_diff)

}




