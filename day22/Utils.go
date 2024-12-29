package main

import (
	"strconv"
	"strings"
)

func ExtractIntput(buffer string)(out[]Secret){
  lines := strings.Split(strings.Trim(buffer,"\n"),"\n")

  for i := range lines {
    val,_ := strconv.Atoi(lines[i])
    out = append(out,NewSecret(int64(val)))
  }

  return out
}

type Secret struct {
  Initial int64
  Current int64
}

func NewSecret(initial int64)Secret{
  return Secret {
    Initial: initial,
    Current: initial,
  } 
}

const MOD_VALUE = 16777216
const MULT_VALUE =2048

func (s *Secret) Search(limit int){

  for i:=0;i<limit;i++{
    secret := s.Current

    //MULT

    secret ^= 64 * secret
    secret %= MOD_VALUE 

    //DIV

    secret ^= secret / 32
    secret %= MOD_VALUE 

    //MULT 2048

    secret ^= MULT_VALUE * secret
    secret %= MOD_VALUE 

    s.Current = secret

  }
}

func (s *Secret) Scan(cache map[string]int,limit int){

  sequence := "" 
  seen := map[string]bool{}
  prev_digit := 0
  secret := s.Current
  last_digit := secret %10

  for i:=0;i<limit;i++{
    last_digit = secret %10
    diff := last_digit -int64(prev_digit)

    if i > 0 {
      if i>4 {
        if string(sequence[0]) == "-"{
          sequence = sequence[2:]
        }else {
          sequence = sequence[1:]
        }
      }
      sequence += strconv.Itoa(int(diff))

      if i>4 {
        if _,has := seen[sequence];!has{
          cache[sequence] += int(last_digit)
          seen[sequence] = true
        }      
      }
    }
    //MULT

    secret ^= 64 * secret
    secret %= MOD_VALUE 

    //DIV

    secret ^= secret / 32
    secret %= MOD_VALUE 

    //MULT 2048

    secret ^= MULT_VALUE * secret
    secret %= MOD_VALUE 

    s.Current = secret
    prev_digit = int(last_digit)
  }

}

