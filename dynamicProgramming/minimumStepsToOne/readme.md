# minimum steps to one

## Problem Definition
On a positive integer, you can perform any one of the following 3 steps:
1. Subtract one from it
2. If it's divisible by 2, divide by 2
3. If it's divisible by 3, divide by 3

Given a positive integer n, find the minimum number of steps that takes n to 1.

## Thoughts on Bottom-Up solution
- In general, we build a table and for each index run each operation (if they apply): 1 + dp[index], 1 + dp[index % 2], and 1 + dp[index % 3]
- Because we always attach the min number of steps with each operation to that index in the table, we are effectively trying out each operation that could happen at that index and only keeping the least one.
- The value at each index in the table is the minimum amount of steps that is required to get to that index in the table
- We do not have the actual steps that are required to get to that index, however.

## Questions to Ask
- Are there any space or time constraints on the problem?
- Do we need to find out exactly which operations get to that index?
- Do you want us to use top-down, bottom-up, or any specific solution?

