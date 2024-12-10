## Day 10: Hoof It

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
1. Create a 2D array from the input
2. For each trailheads start position (cell = 0)
    - Use a BFS to find the paths toward all the cells where the value = 9 
    - Make sure to only allow increasing value path.
3. Count the paths of length 9 
    - If we have multiple paths of src to des only count it once.

#### Part2
##### Same as Part1 but count all the paths
