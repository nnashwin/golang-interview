# paint-houses

## Problem Definition
There are a row of n houses, each house can be painted with one of the three colors: red, blue or green. The cost of painting each house with a certain color is different. You have to paint all the houses such that no two adjacent houses have the same color.

The cost of painting each house with a certain color is represented by a n x 3 cost matrix. For example, costs[0][0] is the cost of painting house 0 with color red; costs[1][2] is the cost of painting house 1 with color green, and so on... Find the minimum cost to paint all houses.

## Questions to Ask
- Are the costs of each house positive numbers?
- Can we reuse the array we are given, or should we build a new one?

## Analysis
Since the amount of houses is finite, you just need to store in each row what the minimum cost of painting this house with the two separate colors is at this array index.
So for the first cost index in the array.
For Array Index [1][0], which is the first index of the second array in the 2d array, we add the minimum of the paint cost of the painting of the previous house, house 1, to the current cost of painting this house, house 1. 
So if the previous cost array (for house 1) was [5, 2, 3], and this current cost is [4, 6, 8], then for costs[1][0] we would add 4 to the minimum of 2 and 3, which is 2.  So four would become six and would represent the permutation of painting the first house the 1 color, and the second house the 0 color. 
We continue to do this for each house and obtain the cost of painting all the houses in the final house's paint cost. 
We find the min of those 3 permutations, and return that number

## Array params
The i param, or the nested arrays represent which house is being painted.
The j param, or the indices in each array, represents the cost of the color of the house painted at each index.
