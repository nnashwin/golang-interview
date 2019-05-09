# arrayPermutation

## Definition
A permutation can be specified by an array P, where P[i] represents the location of the element at i in the permutation. For example, [2, 1, 0] represents the permutation where elements at the index 0 and 2 are swapped.

Given an array and a permutation, apply the permutation to the array. For example, given the array ["a", "b", "c"] and the permutation [2, 1, 0], return ["c", "b", "a"].

## Solutions

### Creating a New Array
The first solution creates a new array of the same length as the original array and assigns the values in the original array at the indices specified by the permutation array.  In [Strong-Typed Languages](https://en.wikipedia.org/wiki/Strong_and_weak_typing) like Golang, we need to use a separate array of the type that needs to be returned to arrange the input and the output. 

This implementation runs in O(n) speed since it needs to loop through the permutation array in its entirety in order to set the indices correctly. We also technically incur another O(n) cost to create the new array, but O(n) + O(n) reduces to O(2n), so we still consider it O(n) complexity.

The space complexity of this solution is also O(n), as we create a separate array to hold all of the values.

### Modifying the Permutation Array
In the second solution to the problem, we use the permutation array itself and assign the appropriate values at the permutation index of the first array into the permutation array.

## Questions to ask
1. Will the two input arrays always be the same length?
1. How should we handle two empty arrays?
1. Mention the static / dynamic type limitation and talk about the time complexities due to that. 