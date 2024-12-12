## Day 12: Garden Groups

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
1. Extract Garden map from input 
2. Use BFS to Find the different sections 
3. For each section 
    - For each cell in the section 
        - check if the cell is a border
    - Sum up the valid bordering cells

#### Part2
1. Extract Garden map from input 
2. Use BFS to Find the different sections 
3. For each section 
    - For each cell in the section 
        - Count the number of corners
    - Sum all corners found 
