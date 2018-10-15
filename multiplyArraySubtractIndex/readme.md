# multiplyArraySubtractIndex

## Definition
Given an array of integers, return a new array such that each element at index i of the new array is the product of all the numbers in the original array except the one at i.

For example, if our input was [1, 2, 3, 4, 5], the expected output would be [120, 60, 40, 30, 24]. 

If our input was [3, 2, 1], the expected output would be [2, 3, 6].

## Follow up
Follow-up: what if you can't use division?

Potential Solution to the follow up:
It seems that we might be able to multiply by 0.1 instead of dividing for integers under 10.
If a value at a particular index is over 10, we chould perhaps define a function that calculates how many decimal places
it has and which "multiplication factor" to multiply by.

I started writing the solution to this problem and realized that handling float multiplication to a specific precision is not trivial in golang.
So I won't be working out that solution for now.


## Complexity
The time complexity of the following algorithm  should be O(2 * length of input array) because we loop through the input twice, once to get the total, and another time to calculate the values at each index.
The space complexity is O(1) because we keep the same input array.

## Gotchas
- What happens when we multiply by 0?

## Questions to Ask
- How do you want us to handle 0?
- How do you want us to handle negative numbers?
- How do you want us to handle negative decimals or fractions?

