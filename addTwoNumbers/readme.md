# Add Two Numbers

## Problem Statement
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.

## Things Learned
- Feeding a traversal algorithm a dummy node and manipulating the .Next works well with linked lists
- Not accidentally appending a final Node to the List with a default value is difficult.
- There is a lot to learn from not knowing the answer to a problem, trying the problem by yourself for half an hour, then reading the answer / approach
- If a number is under 10, dividing by 10 will equal 0
- If a number is under 10, modulo by 10 will equal the number

## Questions to Ask from Problem Statement
- How do you want me to handle linked lists of different lengths?
- How should we handle empty linked lists?
