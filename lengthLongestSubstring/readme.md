# Length of Longest Substring

## Problem Statement
Given a string, find the length of the longest substring without repeating characters.

Example 1:

Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Example 2:

Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Example 3:

Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

## Things Learned
- In golang, the first part of a map accessor returned is the value stored under the key.
-- If there is no value stored under the key, that value is the empty value for that type i.e. for a map[string]int,  val, _ := map["cookies"], if there are no cookies, this value is 0.
- If I just loop through and assign the characters to the hash, I only get the tail end comparison.  If we imagine the examined substring as a "Sliding Window", my window has gaps because I'm blindly resetting on each duplicate character.  I'm not comparing the substrings in between.
- The Sliding Window can be optimized: instead of moving the pointer one at a time when a duplicate is found, we can store the index of the duplicate each time it is found and use it to jump the Head pointer of the sliding window to the location of the duplicated letter instead of incrementing it by one.
- In this way, instead of potentially having a time complexity of O(2n), we only have to loop through the elements once and can have a time complexity of O(n) since we won't have to wait for the Head pointer AND Tail pointer to move to the end of the string.

## Questions to Ask from Problem Statement
- Are there any memory constraints that are imposed on the algorithm?
- Can we assume all strings are lowercase letters, or will there be capital letters too?
- If a character in the substring is found to repeat, do we count that character as the beginning of the new substring?

## Complexity
- The run time complexity of the algorithm should be O(length of string), which is O(n)
- The space complexity is in the worst-case O(length of string) for when the string has all unique characters and the hash stores each character in the string, which is also O(n)
