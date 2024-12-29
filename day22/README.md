## Day 22: Monkey Market

#### Part1
1. For each buyer 
    - Perform the secret computation 2,000 times. 
    - Return the results obtained from the computations.

#### Part2
1. For each buyer 
    - Perform the computation 2,000 times.  
    - At each iteration
        - Extract the last digit of the current secret.
        - Calculate the difference between the current last digit and the previous one.
        - Generate a key based on the last 4 differences.

2. Maintain a global cache to track keys:
    - For each buyer 
        - increment the global value of each key by the number of bananas sold at its first occurrence.

3. Pick the key with the highest score in the global cache.
