# Move Zeroes

## Problem Definition
Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.

For example, given nums = [0, 1, 0, 3, 12], after calling your function, nums should be [1, 3, 12, 0, 0].

Note: You must do this in-place without making a copy of the array. Minimize the total number of operations.

## Explanation
I have previously done a similar problem where we move all 0s to the end of the array in any order by having one pointer start at the beginning and the other start from the end (len(arr) - 1) of the array and switch them when there are 0s, so I thought this problem might be similar.
A better way though is to loop through the array in order, keeping one pointer to progress through the loop and the other to hold a spot for when you encounter non-0 values to swap.

In the test array we use [1, 2, 0, 3, 0, 5], when we enounter the first element one, we will be swapping the value of and 0, which will still be 1, incrementing the index of the place-holder pointer by one, and continuing the loop.

We continue to increment both pointers in this way until we hit the 3 index (arr[2]) of the array.  Since this is our first 0, we will not be swapping the j (at index 2 at this point) with i (at index 1). This iteration of the loop will finish, then the j pointer will increment again.

This next index is the first time the swap will actually happen successfully.  Since the value at index j (at index 3 now) is not a 0, we will be swapping it with pointer i at the previous index (2).

The result of this operation will be [1, 2, 3, 0, 0, 5].

The loop will continue until the j pointer reaches the final index of the array, 5.  At this point, we will again swap the j and i pointers leaving the array at [1, 2, 3, 5, 0, 5]  (the extra five being there due to the final swap between the i pointer and the j pointer). 

At this point in time, we have swapped all of the non-zero integers into their correct places, but we haven't finished moving the zeroes to the end of the array.  The j index is already at the length of the array, so we run another for loop (a while loop in other languages) to run until the i index
reaches the length of the array.

In each iteration of this loop, we set the i index in the array to 0.

We can move zeroes to the beginning by looping through the array in reverse and running the same logic.

## Complexity
The space complexity of the following operation is O(1) and the time complexity is O(2 * length of array) because we could potentially need to loop over the entire array twice if the array only has 0 values in the worst case.
O(2n) reduces to O(n).
