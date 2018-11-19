# gold-mine-problem

## Problem Statement

Gold Mine Problem

Given a gold mine of n*m dimensions. Each field in this mine contains a positive integer which is the amount of gold in tons. Initially the miner is at first column but can be at any row. 
He can move only (right->,right up /,right down\) that is from a given cell, the miner can move to the cell diagonally up towards the right or right or diagonally down towards the right. Find out maximum amount of gold he can collect.

Examples:

Input : mat[][] = {{1, 3, 3},
                   {2, 1, 4},
                  {0, 6, 4}};
Output : 12
{(1,0)->(2,1)->(2,2)}

Input: mat[][] = { {1, 3, 1, 5},
                   {2, 2, 4, 1},
                   {5, 0, 2, 3},
                   {0, 6, 1, 2}};
Output : 16
(2,0) -> (1,1) -> (1,2) -> (0,3) OR
(2,0) -> (3,1) -> (2,2) -> (2,3)

Input : mat[][] = {{10, 33, 13, 15},
                  {22, 21, 04, 1},
                  {5, 0, 2, 3},
                  {0, 6, 14, 2}};
Output : 83


## Questions to Ask
- Verify there can not be negative numbers in the mine
- Verify the miner will start at the first column, but potentially be in any row.

## Things I struggled with / gotchas
- When you are modifying the table at an index, you must do the index of the MINE + the right, top right, and bottom right of the DP TABLE.
-- Although I knew this in other problems, I forgot it here.  So during my implementation, I originally had this bug and realized it was because I was effectively picking up gold and dropping it.
- You really have to understand that the way tabulation (bottom-up) works is by calculating the subproblem over and over again while updating the dp table.
-- The reason it can actually work is because the input of the subproblem will increase with each row or column iteration that is stored in the subtable.

So if I am grabbing the max value, I will update the dp table starting from the right:

Step one will look like this.
dp[][] = {{0, 0, 0, 0},
		  {0, 0, 0, 0},
		  {0, 0, 0, 0},
		  {0, 0, 0, 0}};

Step two will have nothing to compare to, so you only get 15:
dp[][] = {{0, 0, 0, 15},
		  {0, 0, 0, 0},
		  {0, 0, 0, 0},
		  {0, 0, 0, 0}};

Step three will have compared 1 and 15 and chosen 15 to add to the 13 that is at the number 2 index of the first nested array:
dp[][] = {{0, 0, 28, 15},
		  {0, 0, 0, 0},
		  {0, 0, 0, 0},
		  {0, 0, 0, 0}};

The algorithm is greedish, so it takes the max of what it can see, but in the next nested array (mat[1]) it will see the results in the dp table from the first pass through and can choose those as well.
