# maxSumOfThreeSubarrays

## Problem Definition
In a given array nums of positive integers, find three non-overlapping subarrays with maximum sum. Each subarray will be of size k, and we want to maximize the sum of all 3*k entries. Return the result as a list of indices representing the starting position of each interval (0-indexed). If there are multiple answers, return the lexicographically smallest one. Example:

Input: [1,2,1,2,6,7,5,1], 2
Output: [0, 3, 5]
Explanation: Subarrays [1, 2], [2, 6], [7, 5] correspond to the starting indices [0, 3, 5].
We could have also taken [2, 1], but an answer of [1, 3, 5] would be lexicographically larger.

## Questions to ask
- Is the maximum sum talking about the maximum sum within the array, or the maximum sum of relative indices in the array?
-- I.E. so if I had 3 arrays [1, 2], [3, 4], [5 6] would I be adding the arrays together, or adding the relative indices together?
