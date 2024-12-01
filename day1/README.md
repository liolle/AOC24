## Day 1: Historian Hysteria

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
1. Slip each line to get the 2 integers
2. Add those into a 2 slices
3. Iterate over the 2 slices summing the absolute difference?

#### Part2
1. Slip each line to get the 2 integers
2. Add the first integer in a slice and count the second one using a map
3. Iterate over the slice summing `slice[i] * map[slice[i]]`.
