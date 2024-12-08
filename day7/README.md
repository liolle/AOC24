## Day 7: Bridge Repair

#### Setup 
Extract input form file 
```go
  filePath := "input.txt"
  buffer, err := os.ReadFile(filePath)
  if err != nil {
    fmt.Println("Error opening file:", err)
    return
  }
```

#### Part1 && Part 2
1. For each line brute force by using a BFS over all possible operation combination 
2. For each line sum the number of valid combination (where computation amount to Target)
3. Sum the Target of lines that have at leas 1 valid combination
