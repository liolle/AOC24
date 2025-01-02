package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func ExtractInput(buffer string)(out Bag){
  schemas := strings.Split(strings.Trim(buffer,"\n"),"\n\n")
  
  locks := []Schema{}
  keys := []Schema{}

  for i := range schemas {
    schema := NewSchema((schemas[i]))
    if schema.Type == LOCK {
      locks = append(locks,schema)
    }else {
      keys = append(keys,schema)
    }
  } 

  out = Bag{
    Locks:locks,
    Keys:keys,
  }
  return out
}

type SCHEMA_TYPE int 

const (
  LOCK SCHEMA_TYPE = iota
  KEY
)

type Schema struct {
  Type SCHEMA_TYPE 
  KEY string
  PINS []int
}

type Bag struct {
  Locks []Schema
  Keys []Schema
}

func (b Bag) CountValidKey()(out int, keys []string){

  for i := range b.Keys {
    for j := range b.Locks {
      valid := true
      for k := range b.Keys[i].PINS {
        sum := b.Keys[i].PINS[k] + b.Locks[j].PINS[k]
        if sum >5{
          valid = false
          break
        }
      }
      if valid {
        lock := b.Locks[j].KEY
        key := b.Keys[i].KEY
        keys = append(keys,fmt.Sprintf("%s-%s",lock,key))
        out++
      }
    }
  }
  return out, keys
}

func NewSchema(str string)Schema{
  lines := strings.Split(str,"\n")
  var stype SCHEMA_TYPE 
  key := ""
  pins := make([]int,len(lines[0]))

  if regexp.MustCompile("\\.+").MatchString(lines[0]){
    stype = KEY
  }else{
    stype = LOCK
  }

  for i := 0;i<len(lines);i++{
    for j := 0;j<len(lines[i]);j++{
      if string(lines[i][j]) == "#"{
        pins[j]++ 
      }
    }
  }

  for i := range pins {
    pins[i]--
    key += strconv.Itoa(pins[i])
  }

  return Schema{
    Type: stype,
    KEY: key,
    PINS: pins,
  } 
}
