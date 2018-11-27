# lastNonDuplicate

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

## Difference from firstNonDuplicate
1. This implementation does run two for loops, but one starts from the front, and another from the back of the string array.
2. I chose not to loop through the string with range in the first loop because it uses the rune, and when looping backwards you don't have the rune so easily accessible.
3. If the rune method is used, you need to return string(rune(str[i])) to return a string from the function.  This seems like too many conversions for me!
