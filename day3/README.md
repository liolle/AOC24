## Day 3: Mull It Over

#### Setup 
Extract input form file 
```go
filePath := "input.txt"
buffer, err := os.ReadFile(filePath)
if err != nil {
  fmt.Println("Error opening file:", err)
  return
}```

#### Part1
1. Use regular expression to extract all valid multiplication group
2. Compute each multiplication

#### Part2
1. Discard invalid group by splitting the input over "don't()" and "do()" 
2. Compute the remaining multiplication
