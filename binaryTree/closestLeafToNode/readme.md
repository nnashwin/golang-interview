# closestLeafToNode

## Problem Definition

Closest leaf to a given node in Binary Tree

Given a Binary Tree and a node x in it, find distance of the closest leaf to x in Binary Tree. If given node itself is a leaf, then distance is 0.

Examples:

Input: Root of below tree
       And x = pointer to node 13
          10
       /     \
     12       13
             /
           14
Output 1
Distance 1. Closest leaf is 14.


Input: Root of below tree
       And x = pointer to node 13
          10
       /     \
     12       13
           /     \
         14       15
        /   \     /  \
       21   22   23   24
       /\   /\   /\   /\
      1 2  3 4  5 6  7  8

Output 2
Closest leaf is 12 through 10.

Questions to ask:
1. Is the binary tree a bst?  Does it have order?
