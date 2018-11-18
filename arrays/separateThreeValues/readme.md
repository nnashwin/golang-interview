# separate-three-values

## Problem Definition
Given an array of strictly the characters 'R', 'G', and 'B', segregate the values of the array so that all the Rs come first, the Gs come second, and the Bs come last. You can only swap elements of the array.

Do this in linear time and in-place.

For example, given the array ['G', 'B', 'R', 'R', 'B', 'R', 'G'], it should become ['R', 'R', 'R', 'G', 'G', 'B', 'B'].

## Questions to ask
- When you say In-place, do you mean space complexity of O(1)?
- Will there be any variable values?  Are we sure they are 'R', 'G', and 'B'?

## First Solution
In the first solution, we loop over the array twice, once forwards and once backwards, while swapping either the R or B to keep the G in the middle of the array.
Although this maintains O(N) complexity, it is technically O(2N).  So it does reduce down, but we can do better.

## Dutch national flag solution
After consulting the oracle, I discovered that Dijkstra had solved this WAAAYYYY before my time in something called the [Dutch national flag problem](https://en.wikipedia.org/wiki/Dutch_national_flag_problem).
Basically, we can swap out the > or < mid conditionals in the pseudocode with B and R respectively.
This solves the problem in O(N) regardless of the input.
