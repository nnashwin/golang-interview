# firstNonDuplicate

## Problem Definition
Given a string, find its last non-repeating character

Given a string, find the  non-repeating character in it. For example, if the input string is “GeeksforGeeks”, then output should be ‘r’ and if input string is “GeeksQuiz”, then output should be ‘z’.

## Questions
- Is there a specific time / space complexity you want it in?
- How do you want me to handle an empty input?

## Complexity
- The time complexity of this problem is O(2n) since in the worst case we could loop through the string twice.
This reduces to O(n).
- The space complexity of this problem is O(n), since if every character in the string is different, the hash would contain each character.
