# Is a String a Palindrome Recursive

## Problem Statement
Determine whether or not a string is a palindrome recursively

## Solution
We define the function to take a string and return a boolean.
The two base cases:
1. If the characters at the front and end of the string are not the same, return false.
2. If the length of the string is 1 or less, it is a palindrome and we should return true.

The recurse case is hit if none of the previous two base cases are triggered.

## Gotchas
The "is a string a palindrome" question is a common interview question to see if an engineer has much programming experience.
The little twist of adding recursion on top of the algorithm also lets the interview dive a bit into the cs background of a candidate.

I had originally seen this algorithm in Javascript and was tripped up by appropriate slicing of the text in the recursive solution.

Given a string "racecar", you compare the character at the last index in the array (str[len(str) - 1]) with the first character in the array (str[i]).  In this string they are both r respectively.
The gotcha can occur when you call the isPalindrome recursively and is a Go-specific idiosyncracy.

Slices of a slice include the first index and not the second.
When recursing, you should take taking s[1:len(str) - 1] instead of s[1:len(str) - 2].
So for the string "racecar", the first recursion will amount to s[1:6] ("aceca") instead of s[1:5] ("acec").

## Time Complexity
The time complexity of this algorithm in the worst case is O(log N).  This occurs if a string is a palindrome and the algorithm runs all the way through the string.

## Optimizations
We could modify the algorithm to filter out punctuation and convert the string to lowercase in order to compare more varied inputs.
