# Is a String a Palindrome Recursive

## Problem Statement
Determine whether or not a string is a palindrome in an iterative fasion.

## Solution
We define the function to take a string and return a boolean.
We loop up to the number of characters / 2 of a string and compare the index and last character - index for each value.

## Gotchas
The "is a string a palindrome" question is a common interview question to see if an engineer has much programming experience.
In the iterative solution of this problem, we use a for loop to iterate over the string instead of recursion.

The only gotcha in this problem is that we need to make sure we set the final index to compare against to len(string) - 1 to begin in order to not hit any slicing memory issues.

## Time Complexity
The time complexity of this algorithm in the worst case is O(log N).  This occurs if a string is a palindrome and the algorithm runs all the way through the string.

## Optimizations
Similar to the recursive solution, we could modify the algorithm to preprocess the string to filter out punctuation and convert the string to lowercase in order to compare more varied inputs.
