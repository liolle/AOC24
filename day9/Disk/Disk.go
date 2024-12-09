package disk

import (
	"fmt"
)

type Disk struct {
  Map []Layout
  Compilation []int
}

type Layout struct {
  Files int 
  FreeSpaces int
}

func (disk Disk) Print(){
   fmt.Println(disk.Compilation)
}

func (disk *Disk) Compile(){
  for idx,l := range(disk.Map) {
    for i := 0; i < l.Files; i++ {
      disk.Compilation = append(disk.Compilation,idx)
    }
    
    for i := 0; i < l.FreeSpaces; i++ {
      disk.Compilation = append(disk.Compilation,-1)
    }

  }
}

func (disk *Disk) Fill(){
  for l,r := 0,len(disk.Compilation)-1 ; l<r;  {
    for l<r && disk.Compilation[l] != -1{l++}
    disk.Compilation[l] = disk.Compilation[r]
    if l != r {disk.Compilation[r] = -1} 
    r--
  }

}

func (disk *Disk) FullFill(){

  n := len(disk.Compilation) 

  for r := n-1 ; r>0;  {
    for r>0 && disk.Compilation[r] == -1{r--}

    rl := r
    rr := r
    for rl>=0 && disk.Compilation[rl] == disk.Compilation[rr] {rl--}
    rl++

    for l := 0;l<r; {
      for l<r && disk.Compilation[l] != -1{l++}

      ll := l
      lr := l
      for lr<r && disk.Compilation[lr] == -1{lr++}
      lr--
      if ll<=lr {
        space_left := lr - ll + 1
        file_size := rr - rl + 1

        if space_left >= file_size {
          elem :=  disk.Compilation[rr] 

          for i :=0; i < file_size; i++ {
            disk.Compilation[l+i] = elem
            disk.Compilation[rr-i] = -1
          }

          break
        } 
      }
      l = lr+1

    }
    r = rl-1
  }

}

