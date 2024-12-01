package main 

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
)
// count number in the second list 
// iterate over number in the first list adding the number occurence in the sum 
func main()  {

  file, err := os.Open("input1.txt")
  if err != nil {
    return
  }

  gp1 := []int{}
  gp2 := make(map[int]int)
  

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

      _,ok := gp2[num2]
      if ok {
        gp2[num2] += 1
      }else {
        gp2[num2] = 1
      }
    }
    if err != nil {
      break
    }
  }

  sum := 0
  for i := 0; i < len(gp1); i++ {
      cnt,ok := gp2[gp1[i]]
      if ok {
        sum += cnt * gp1[i]
      }
  }
  
  fmt.Println(sum)
}




