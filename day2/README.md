## Day 2: Red-Nosed Reports

#### Setup 
Extract input form file 
```go
// create File struct 
file, err := os.Open("input1.txt")
if err != nil {return}
defer file.Close()
// create a buffer with the file content
buffer := bufio.NewReader(file)
// read line by line 
line, _, err := buffer.ReadLine()
```

#### Part1
1. Fro each line iterate over the member of the array making sure that for all i in the array `array[i] - array[i-1]` is either between -3 and -1 for all i of between 1 and 3  

#### Part2
1. Use backtracking to test the validity of the array by removing one value, (limit the number of removal to 1).

