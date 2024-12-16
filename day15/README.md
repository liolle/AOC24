## Day 15: Warehouse Woes

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
1. Extract the Warehouse as a 2D array and the moves as a list of vectors.
2. For each move 
    - Recursively push the boxes in front of the robot 
        - make sure all the connected boxes can be pushed 

#### Part2:
##### Same as part 1 but make sure the 2 part of all boxes can be pushed. 
