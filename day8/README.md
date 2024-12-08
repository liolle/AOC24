## Day 8: Resonant Collinearity

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
1. Create a frequency map containing all the antennas of a given frequency
2. For all pair of antennas in the same frequency 
    - Get the vector difference between the two `Vec(antenna2 - antenna1)`
    - Add an antinode at `2*Vec(antenna2 - antenna1)` 

#### Part2
1. Same a part 1 but Add an antinode at `i*Vec(antenna2 - antenna1)` 
    - i start at 2 and increase by 1 as long as the new point is in the board.
2. The answer is the number of antinode + the number of antennas
