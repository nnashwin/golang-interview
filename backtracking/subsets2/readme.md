# subsets2

## Problem Definition
Given a collection of integers that might contain duplicates, nums, return all possible subsets (the power set).

Note: The solution set must not contain duplicate subsets.

Example:

Input: [1,2,2]
Output:
[
  [2],
  [1],
  [1,2,2],
  [2,2],
  [1,2],
  []
]

## Gotchas
- How will you store if you have seen a subset before? (potentially different for different languages)

## Time Complexity
- In my first hacky complexity, since I need to add all of the numbers in each temp array up, the time complexity should be at least (n^2) power 
with a space complexity of O(N - 1) since potentially you could have a duplicate stored for each subset (probably except for the empty array)
