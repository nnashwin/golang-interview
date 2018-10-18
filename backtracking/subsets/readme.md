# subsets

## Problem Definition
Given a set of distinct integers, nums, return all possible subsets (the power set).

Note: The solution set must not contain duplicate subsets.

Example:

Input: nums = [1,2,3]
Output:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]

## Gotchas
- Golang sometimes pollutes slice memory when append a slice to a slice of slices
-- I.E. I found that although my subset slice was appending and removing fine in the backtracking algorithm, without creating a new copy of the slice,
I would overwrite parts of previously written slices in the powerset.
- You MUST pass in a pointer of the variable whose memory will be altered

## Questions to Ask
- Do we add the blank "[]" array as a subset
- How do we handle empty nums arrays

## Main Idea
In backtracking, you use a depth-first search-like approach to find out all of the different subsets.
As the algorithm is called in a decreasing sized for-loop, elements are added and removed from the subset before and after visit.
So for the first index of the array, 1 in the test, we will first add the empty array as a subset, add the value at the first index to the array and recursively
call the recursive backtrack call which starts the function again.
The next function will add the [1] subset to the results and start the for loop from the second index of the nums array.
This will continue until we add up all the subsets from the first for loop calls to the master array, then the functions on the call stack will begin to return.
After we have added [1, 2, 3], the call stack will return to the previous function, remove the 3 from the array, and return.
This will continue until the call stack returns to the first function call, when the for loop will then start calling functions from the second index of the array, 2, and start the process over again.
