## Day 6: Guard Gallivant

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
1. Use to create the list of cell visited by the guard

#### Part2
1. Collect the list of cell visited by the guard 
2. For each cell in the visited path
    - Add and obstacle 
    - Check for cycle with that new obstacle (same cell with same direction)
    - Clear the added obstacle
3. Sum the number of cell that create a cycle 
