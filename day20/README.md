## Day 20: Race Condition

#### Part1 & Part2
1. Use a BFS to determine the length of the optimal path for every cell that is not a wall.
    - Maintain a list of cells that constitute the path.

2. For each cell in the computed path: 
    - Identify all cells within a Manhattan distance of 2-20 that are not walls
    - Calculate the gain obtained by transitioning from the current cell to the target cell, factoring in the distance.
    - Count how many of these cells yield a gain greater than the required minimum gain (100)
