## Day 21: Keypad Conundrum  

#### Part1 && Part2
1. For each input code 
    - Extract the multiplication coefficient associated with the input code.
    - Generate the initial list of directional Keypad codes.

2. Recursively compute the minimum length of the best path for each directional code list: 
    - Recursively compute the length of the best.
        - Start each computation with a depth of `3 - 26` 
        - For every pair of directional Keypad inputs within the current directional code:.
            - Recursively determine the minimum length by decrementing the depth `depth -1`
            - Sum the best results for each pair to compute the overall result.
        - When the depth reaches 1, return the length of the current directional code.
        - To go past 3 robot use a cache with combination of `start_key`, `end_key`, and `depth` as the cache key. 

3. Compute the result
    - Multiply the coefficient of each input code by the length of the best path found for its corresponding directional code.
    - Sum these products to compute the final result.
