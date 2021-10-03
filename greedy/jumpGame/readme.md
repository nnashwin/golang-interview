# Jump Game

## Problem Statement
You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.

Return true if you can reach the last index, or false otherwise.

 
```
Example 1:

Input: nums = [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
Example 2:

Input: nums = [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.
 

Constraints:

1 <= nums.length <= 104
0 <= nums[i] <= 105
```

## Solution
When I first encountered this problem I wasn't sure how to solve this.

I thought it might be a dp problem, but it's structure also doesn't seem to require doing continuous calculation while returning the result.

In further iterations of this problem where we must find the minimum jumps required, using some type of recursion (or dp) is necessary to store all possible combinations we can use.

It turns out that the main mechanism we use to determine whether a jump can happen is seeing if we can drag a lastJumpSpot index all the way from the left of the array to the beginning (if there is a path).

We loop backwards through the array, and if the max jump size (each value at the idx) of the idx + the idx is greater than the lastJumpSpot, we modify the lastJumpSpot to equal the current index.

After we run the algorithm, we return a conditional that the lastJumpSpot == 0 (the beginning of the array).  This signifies the frog being able to jump to each spot in the array.

## Time / Space Complexity
Since we loop backwards through the array, the time complexity of the solution is O(n), and the space complexity is O(1) as we just keep a pointer to the lastJumpIdx that we modify.

If we did choose to use a dp array, the space complexity would also be O(n) for each position in the dp array.