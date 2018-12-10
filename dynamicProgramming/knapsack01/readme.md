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

