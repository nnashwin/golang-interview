# wordSearch

## Problem Definition
Given a 2D board and a word, find if the word exists in the grid.

The word can be constructed from letters of sequentially adjacent cell, where "adjacent" cells are those horizontally or vertically neighboring. The same letter cell may not be used more than once.

Example:

board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

Given word = "ABCCED", return true.
Given word = "SEE", return true.
Given word = "ABCB", return false.

## Questions to Ask
- I'd like to modify the array we will be using, do you think it's fine if I use a "-"
- Will the inner arrays of the 2d array all be the same length?
- Which direction can you find the string in, any?
