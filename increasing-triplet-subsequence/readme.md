# increasing-triplet-subsequence

## Definition
Given an unsorted array, return whether an increasing subsequence of length 3 exists in the array or not.

Input: An array
Output: True or False

## Formal Definition of the function
Return true if there exists i, j, k
such that arr[i] < arr[j] < arr[k] given 0 <= i < j < k <= n - 1

Else return false

## Examples
Example 1:
Input: [1, 2, 3, 4, 5]
Output: true


Example 2:
Input: [5, 4, 3, 2, 1]
Output: false

Example 3:
Input: [5, 3, 1, 6, 3, 8]
Output: true

## Gotchas
- Need to set the second_min to the highest number in the language, or else it might conflict with the min pointer.

## Questions to ask
- Should the numbers need to be able to be increased consecutively.  Or if ANY three numbers are in increasing order do they count?

## Things learned
- It is important in this question to specify exactly what the question is asking for.  It seems that it could be extremely simple, deceptively simple, and would throw many
people off track with an unclear definition.
