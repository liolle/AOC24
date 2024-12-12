## Day 4: Ceres Search

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

#### Part1
1. Create a 2D array out of the puzzle input 
2. For each cell make DFS of DFS only allowing neighbors that follow the same direction

#### Part2
1. Create a 2D array out of the puzzle input 
2. Use regular expression to check if `array(NW) + array(SE) + array(NE) + array(SW)` is a valid combination `(SM|MS)(SM|MS)`


