## Day 14: Restroom Redoubt

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
1. Create an Array of Robot from the input 
2. For each robot 
    - Compute the new position after 100 second by multiplying the velocity by 100 the add it of the initial position
    - Normalize the position by using the modulo on the new coordinates   
    - Base on the normalized coordinates assign a quadrant to the robot
3. For each quadrant count the number of robots 

#### Part2
1. Create an Array of Robot from the input 
2. For each robot 
    - Compute the new position after 1 second 
    - Normalize the position by using the modulo on the new coordinates 
3. Convert the robot placement into an image 
    - Look for an image of a tree 
5. Repeat 3 & 4 until the tree is found.
6. Once found retrieve amount of second past to get there.
