# Knight's Tour
> Classic backtracking problem.  It's even number one on [geeks for geeks](https://www.geeksforgeeks.org/the-knights-tour-problem-backtracking-1/)

## Problem Statement 
From Wikipedia:
A knight's tour is a sequence of moves of a knight on a chessboard such that the knight visits every square exactly once. If the knight ends on a square that is one knight's move from the beginning square (so that it could tour the board again immediately, following the same path), the tour is closed; otherwise, it is open.[1][2]

The knight's tour problem is the mathematical problem of finding a knight's tour. Creating a program to find a knight's tour is a common problem given to computer science students.[3] Variations of the knight's tour problem involve chessboards of different sizes than the usual 8 × 8, as well as irregular (non-rectangular) boards.

## Implementation
Backtracking algorithms have stumped me since I first encountered them a couple of years ago.  Yes, the basic pattern I can understand.

We can see the basic backtracking part of the function in the recurFunc on line 25 of the main.go file in this folder.
At each step of the algorithm, we calculate the next move the knight could possibly do. 

The recursive function in backtracking algorithms tend to follow this pattern:
1. if the condition for the solution is found, return true 
2. Check if a potential move is safe
3. if so, make the move.
4. call the recursive function with the same move.
5. if the result of the recursive function is true, return true
6. if not, reset the index to be unvisited
7. return false.  For the move chosen higher up in the recursive stack, all of its branches in the tree did not result in a solution.  So we must rewind the stack and try a different move from a higher level.

Step 6 is important here. This step allows us to calculate solutions that **fail further down the call stack** . Using this technique, we don't need to throw out the first couple steps of the algorithm that we have proved to be correct.  Resetting the index to not used also enables the algorithm to unstick itself at specific locations.
There could be situations where choosing a sequence of moves backs the knight into a corner where it can not make any safe moves to get out.  In this situation, the backtracking algorithm will "let the knight leave" by unwinding the stack and marking the index as -1.  Ultimately, the index that is unmarked will need to be visited 
but it will be in a different sequence.

## Gotchas
I feel that the gotchas with backtracking algorithms are problem-specific and algorithm-specific. If you haven't solved a backtracking algorithm before, you will need to have the basic structure memorized to a point where you can reproduce it for this problem. You will also need to piece together what constraints make up the isSafe function.
In this situation it is comprised by the following: 
1. The move is on the board (not less than 0 and not greater than the length of the board)
2. The space on the board hasn't been filled yet. (== -1)

## Time Complexity
I searched around for time complexity of the algorithm for a long time and it should be O(n!).
The reason for this is that for each time we recurse, one of the remaining vertices is selected, so the total number of vertices that could be selected is decreased by one. However, at each recursion of the function, the knight might need to try out ALL of the spaces in function calls further down the recursive stack in order to find the correct next step. He might also get stuck at this step and need to unwind the call stack up again.
This means that at each step N of the algorithm the knight might need to try all of the spaces N.  
So Time Complexity of (N) = N * the next recursed function call, which is Time complexity of (N - 1).
So N* * N-1 * N-2... will result in O(N!).
