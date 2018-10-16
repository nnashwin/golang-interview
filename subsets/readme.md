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
