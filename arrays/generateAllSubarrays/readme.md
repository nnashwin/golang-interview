# Print all Subarrays

## Problem Statement
Given an array, print out all subarrays in the statement.

Input of the function will be the an array of ints
The function will print out all subarrays.

## Unoptimized Solution
In the unoptimized solution, we nest 3 for loops to print out the correct values.
The first for loop index (i) is the start of each sub-array index.  It runs from 0 to the length of the array.
The second for loop (j) is the end of each sub-array index. It runs from i to the length of the array.
The third for loop (k) runs through each index between i and k and focuses on one element to print at a time.
You can think of these two indices as a bookends on a book shelf that partition off a selection of books you can choose from.
The third index would be like the focus of your eyes or your hand running through each book in the book-ended selection.

If we use the array [1, 3, 5, 7], the i index will start from 0. For each iteration of the first nested loop j, the range of values that are included in the sub-array increase.
When i is 0 and j is 0, the subarray can only be in the range 0, 0.  K only runs once and prints out the value of that subarray ([1]).
When i is 0 and j is 1, the subarray can be in the range 0, 1.  K will run twice and print out the value of that subarray ([1, 3]).
This continues all the way to the end of the array.
Then, when we get to the end of the array, i increments to 1 and j resets itself.
So the first iteration of the loop will be i at 1 and j at i (1).  So in the first iteration, the subarray can only be in the range 1, 1.  K will run only once and print out the value of that subarray([3]).
In the second iteration, i will be at 1 and j at 2.  The subarray is in the range 1, 2, and k will run twice and print out the value of the subarray ([3, 5])

This continues until the end of the array.

Since their are 3 nested for loops, the time complexity of this algorithm is O(n^3).  The space complexity is O(1). We can make it faster than this.

## Optimized solution
In the optimized solution, we use a pattern found from a [different problem](https://web.stanford.edu/class/cs9/sample_probs/SubarraySums.pdf) to limit the amount of loops we use.
Instead of using one loop to set the beginning index (i), one loop to set the end index (j), and the third to actually focus on each element (k), we take an approach that mixes the j and k loops together.
In the optimized solution, we use two data structures to help us keep track of all values we touched and all subarrays we have seen.  The first data structure is a result array that stores arrays of ints.  This represents an array that will store ALL subarrays. This result array of arrays is instantiated before any looping occurs.
The second data structure we use is an array of ints called a subArrContainer.  This data structure holds an increasing amount of values that represents the a specific subarray at each tick of the j loop.

Whereas in the unoptimized solution, at i = 0 and j = 1 the the i and j pointer only acts as bookends to reign in the range of values the k pointer can run through.  In the optimized solution, when i = 0 and j = 1, before the loop even begins, the subArrContainer already holds the value that comprises the subarray at i = 0 and and j = 0 ([1]).
After the loop logic executes at i = 0 j = 1, the subArrContainer contains the values of subarray at i = 0 j = 0 ([1]) AND the values that are in the subarray at i = 0, j = 1 ([1, 3]).
Since the subArrContainer contains a valid subarray in each iteration of the j loop, we append the container each time we add a new value to the subArrContainer.

The time complexity of this solution is O(n^2 + n). The n is the final loop through all of the result subarrays.  It can be simplified to O(n^2).
The space complexity should be O(max size of subarray + max size of subarray - 1 + max size of subarray - 2...) for each starting index of a subarray. This is because we will store an array of arrays that contain all subarrays from the longest length down to 1, then from the longest length - 1 down to 1, then from the longest length - 2 down to 1.
I am almost certain there are better math terms to use than what I just used.  If you know them, please either make a pr or add an issue.

## Best solution
After looking at my "optimized" solution, I realized there was a better one.  You can edit out the logic complexity of using the results array and looping through that array by just printing after each subArray container append.
This works because the subArrContainer contains the subArray from index i to the current iteration of index j at this point.  You can just print out that subArray each time you append the value to the container and call it a day.
This results in time complexity O(n^2) and space complexity of O(length of subarray).
