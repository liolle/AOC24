## Day 17: Chronospatial Computer 

#### Part1
1. Execute each instructions in order while the instructions counter in smaller than the number of instructions

#### Part2
1. Initial Observations:
    - The answer fits within a 48-bit space (equivalent to 16 octal digits).
    - The program loops 16 times, printing the 3 rightmost bits of the current value of A in the register during each iteration.
    - In each iteration:
        - A series of operations is performed on the 3 rightmost bits of A. These operations depend on the 10 rightmost bits of A.
        - The value of A is shifted 3 bits to the right. Consequently, the outputted value does not influence subsequent computations.
    - The answer lies in the range [8**15, 8**16).

2. Finding the 3 Most Significant Bits (MSBs):
    - To determine the 3 leftmost bits of the answer `A % 8`, it is necessary to modify the 10 leftmost bits of the current value of A in the register.

3. Step-by-Step Approach:
    - Start with `A = 8**15`.
    - Modify the 10 leftmost bits of A by incrementing:
        - `A += 1 << (38 - 3 * idx)` until the 3 rightmost bits of A match the value corresponding to the current index idx being sought in the answer.
    - Once the match is found, shift A 3 bits to the right `(A >>= 3)` and proceed to the next index 


