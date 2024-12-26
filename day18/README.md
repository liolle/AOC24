## Day 18: Lavaduct Lagoon

#### Part1
1. Use a BFS to find the shortest path.

#### Part2
1. Use BFS to calculate the shortest path from the starting point to the destination.

2. Maintain a list of cells that constitute the current shortest path.

3. Use an obstacle index to identify the next obstacle to be placed.
    - Increment the obstacle index and continue adding obstacles until the newly added obstacle intersects a cell on the current shortest path.

4. Whenever an obstacle blocks a cell on the current path:
    - Recompute the shortest path using BFS.
    - Update the tracked path to reflect the new shortest path.
    - If no viable path exists, return the obstacle associated with the last obstacle index.
