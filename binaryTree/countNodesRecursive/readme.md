# Count Nodes in a binary tree recursively

## Problem Statement
The recursive structure of a binary tree makes it easy to count nodes recursively. 

There are 3 things we can count:

1. The total number of nodes 
2. The number of leaf nodes 
3. The number of internal nodes

### Counting all nodes
The number of nodes in a binary tree is the number of nodes in the root’s left subtree, plus the number of nodes in its right subtree, plus one (for the root itself).

## Solution
The solution for counting the nodes in the bst revolves around just walking the tree.  Using the recursive algorithm uses the call stack behind the scenes.
We first set the count to equal one for each call on the stack, then we add this number to all of the calls that will be added to the stack with each recursive call.
The result is the total number of nodes that are contained within the tree.

## Time Complexity
The time complexity of the algorithm is O(n) where n is the number of each node contained within the binary search tree.

