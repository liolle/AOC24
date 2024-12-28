package main

import (
	"fmt"
	"math"
	"os"
)

type CacheKey struct {
  code string
  depth int
}

func FindPath(start string,end string)[]string{
  return KeyPadPaths[KeyPadVect{keyPadInputValue[start],keyPadInputValue[end]}] 
}

func Expand(cache map[CacheKey]int64,code string,depth int) (ret int64) {
  key :=CacheKey{code,depth} 
  if val,has := cache[key];has {
    ret = val
    return ret 
  }

  if depth == 1 {
    ret = int64(len(code)) 
    cache[key] = ret
    return ret 
  }

  last := "A"

  for _,v := range code {
    key := KeyPadVect{keyPadInputValue[last],keyPadInputValue[string(v)]}
    paths := KeyPadPaths[key]
    var p_min int64
    p_min = math.MaxInt

    if len(paths) == 0 {
      ret+=1
      continue
    }

    for _,p := range paths {
      res :=Expand(cache,p+"A",depth-1) 
      p_min = int64(math.Min(float64(p_min),float64(res)))    
    }

    ret += p_min
    last =string(v) 
  }

  cache[key] = ret
  return ret
}


func search (target NumTarget)[]string{
  h_paths :=[]string{} 

  target.ComputeBestPath(&h_paths,&NumNode{
    Pos: NUM_PAD_A,
    Score: 0,
    Path: "",
  },0,target.Code)

  return h_paths
}

func part1(targets []NumTarget,depth int)(sum int64) {
  cache := map[CacheKey]int64{}
  for i := range targets {
    codes := search(targets[i])
    var c_min int64
    c_min = math.MaxInt
    for j:= range codes{
      c_min = min(c_min,Expand(cache,codes[j],depth))
    }
    str :=fmt.Sprintf("%d * %d",c_min,targets[i].Coef)
    fmt.Println(str)
    sum += c_min*int64(targets[i].Coef)
  }

  return sum
}

func main()  {
  // Open the file
  filePath := "input.txt"
  //filePath := "sample.txt"
  buffer, err := os.ReadFile(filePath)
  if err != nil {
    fmt.Println("Error opening file:", err)
    return
  }

  ts := ExtractInput(string(buffer))
  fmt.Println(part1(ts,3))
  fmt.Println(part1(ts,26))

}
