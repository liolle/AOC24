## Day 11: Plutonian Pebbles

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
1. For each value if the input.
    - recursively compute the number of generated stone

2. Sum up all the computed values 

#### Part2
1. For each value if the input.
    - recursively compute the number of generated stone
    - store the computed values in cache

2. Sum up all the computed values 
