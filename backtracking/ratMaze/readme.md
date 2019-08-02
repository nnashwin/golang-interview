# rat in a maze

## Problem Definition

A Maze is given as Nx*N binary matrix of blocks where source block is the upper left most block i.e., maze[0][0] and destination block is lower rightmost block i.e., maze[N-1][N-1]. A rat starts from source and has to reach the destination. The rat can move only in two directions: forward and down.
In the maze matrix, 0 means the block is a dead end and 1 means the block can be used in the path from source to destination. 

Start from point [0 0] of the maze and find your way through to [n - 1, n - 1]

The rat can move in two directions: horizontally and vertically in any position that is not blocked.

```
Input: 
[[1 0 0 0], [1 1 1 1], [0 0 1 0], [0 1 1 1]]

Output:
[[1 0 0 0], [1 1 1 0], [0 0 1 0], [0 0 1 1]]
```
