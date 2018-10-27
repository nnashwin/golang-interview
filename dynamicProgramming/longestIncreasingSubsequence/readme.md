# Longest Increasing Subsequence

## Better Solution
- In the better solution, a table is used to store the actual sequence of the longest subsequence inside of it.
- As we loop through the numbers of the array, we check for a couple conditions:
	1. Is the array empty?
	- If the array is empty, we append the current number to the table slice, and continue looping through the array
	2. Is the current number greater than the number that is currently at the end of the table slice?
	- If the current number is grater than the number that is currently at the end of the table slice, we can also append to the end of the array and increase the length.
	3. Any other case
	- In any other case, we use binary search on the table of subsequence to look for two other conditions
	- 1. we find elements in the current subsequence that are less than the number (which means the number could be in the longestIncreasingSubsequence)
	- 2. We find numbers that are bigger than the number in the table.  In this situation, we set the length of the binary search to this middle number, and keep looking further down for other numbers that are not bigger.
	4. Finally, we add the current number to the end of the longestSubsequence array, MAYBE REPLACING THE OTHER SUBSEQUENCE NUMBER.  Since we only need to know how LONG the subsequence is, replacing a larger number with a smaller number that has the same length sequence is fine.
	
	So, for example, for the current subsequence array [2, 3, 7, 108], if in the for loop we encounter 18, we will
	1. loop through binary search
	2. Figure out that the 18 is less than 108 and
	3. replace the 18 with the 108 at the end by making the longSubLength the index of 108 (3).

	the function then returns 4, the length of the longestSubsequence in O(n log n) time.
