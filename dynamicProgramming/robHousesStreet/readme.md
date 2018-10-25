# Rob houses street

## Problem Definition
You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security system connected and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given a list of non-negative integers representing the amount of money of each house, determine the maximum amount of money you can rob tonight without alerting the police.

## Gotchas
- With a top-down approach (memoization) to dynamic programming problems, you must start out finding the recursive algorithm. I locked up because I didn't know that 2 + 3 can add up to represent any number (including primes).
- After finding the correct recursive approach, the memoization part was simple, and very similar to a problem called the child jumping stairs problem.

## Time Complexity
- The space complexity of the problem using the table is O(n).
- The time complexity is probably closer to O(n^2).  But using the table The worst case might be around O(n) since we only run through each subproblem once.
