## Day 13: Claw Contraption

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
1. Use BFS with a priority queue to move toward the best result OR use Cramer's Rule 

#### Part2
1. Use Cramer's Rule 
