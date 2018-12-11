# knapsack-0-1

## Problem Definition
A thief is robbing a store and can carry a maximal weight of W into his knapsack. There are n items and weight of ith item is wi and the profit of selecting this item is pi. What items should the thief take?

## Questions to Ask (Not too many, since this is a classic problem):
- Are we sure the values are integers?
- How many copies of the items are allowed (in the classic question, it's usually just 1)?
- Should we try to optimize with DP? (Usually yes)

## Personal Account
The knapsack problem is one I have found particularly confusing in the past.  
The recursive problem seems understandable by following this sequence:
1. The algorithm starts with the last item, sees if taking it, by decreasing the sack weight (W - wt[n - 1]) and adding the value(val[n - 1] + recursive call) end up with a larger value than not choosing it, just making the recursive call n - 1.
2. Return the max of both of those values.

Essentially, we just start from the last value and incrementally compare the result of taking and not taking each item.

This does prove to be slow, as we do MANY comparisons for each situation and add the same value of an item over and over again when we already know how much it will add.

I also always wondered why we can build up a 2d array with this problem.
Now I know.
One of the indices is all of the weights the knapsack could have from 0 to Max weight (W).
The other index is the item index.

So for each item, we see if the item or w is 0.  If so, return 0.
if it's not 0, and the wt of the current item is less than the weight index we are at, we set the current
item and weight value as the max of the value of the current item plus the value of the bag when we take the item and subtract the weight and the 
value we have recorded of when we don't take the item and maintain the current weight.

Even though this previous index was built earlier in the loop, it describes the situation where we have already taken the item.  For example, let's say that we are describing a situation on the second item
where the weight of the item is 20, the value is 30, and the current weight of the knapsack is 23.  In the max calculation at K[2][23], we will compare max(20 + K[1][3], K[1][23]).
The first option was built earlier in the for loop and describes the decisions you can make when you have those weight limitations at that point.  So we are really comparing the max of 20 + best value possible at capacity 3 with previous item and best value possible at capacity 23 with previous item.

Else, if the weight of the item is greater than the current weight index, we set the value of the bag as the same as it was before this item, just like we did in the recursive solution.

