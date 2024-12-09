## Day 9: Disk Fragmenter

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
1. Group number 2 by 2
    - the first number being the File size and the second being The space after this file.
2. Create an array combining the index of each group the sizes and spaces
3. For each group
    - Append the index of the group file_size times in the array 
    - Append the amount of space just after

4. Use 2 pointer to swap empty space in the from of the array with value from the end of the array.

#### Part2
##### 1 - 3 (Same a part1)

4. For each group of identical number stating from the end of the array 
    - List all available space until a big enough space if found to copy the entire group

5. If a space if found copy the group and at the from of the array and delete if from the back.
