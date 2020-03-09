# First Duplicate Minimum Index

## Problem Statement
Given an array a that contains only numbers in the range from 1 to a.length, find the first duplicate number for which the second occurrence has the minimal index. 
In other words, if there are more than 1 duplicated numbers, return the number for which the second occurrence has a smaller index than the second occurrence of all other numbers. If there are no elements with duplicates, return -1.

Examples:
For the array [1, 2, 3, 3, 1, 2, 3], the algorithm should output 3.  There are three duplicates, and the index of the second duplicate of the number 3 occurs before the second duplicate of the other numbers.
For the array [1, 2, 3, 1] with only one duplicate, return 1.
For the array [1, 2, 3] with no duplicates, return -1.

## Solution
This solution is similar to many other duplicate-based questions in that a standard solution can be derived by using a hash.  The twist of needing to record the index of the second occurrence of a number turns the hash into a nested hash.
The map that we construct looks like the following.
**Note: The following is not JSON format**
```
{
    int: {
        Occurrences: int,
        SecondIdx: int,
    },
    int: {
        Occurrences: int,
        SecondIdx: int,
    },
}
```

We loop through the array once, recording all occurences of each int in the map.  If an int shows multiple times in the array, we increment the Occurrences and record the second index position.
After we have built the map, we loop through the keys of the map, record the minimum duplicate idx and return either the value with the minimum duplicateIdx or -1 if there are no duplicates.

## Complexity
The Time complexity of the solution is O(2n) in the worst case of having no duplicates.  In this case, we would loop through all the values of the array to build the hash, then loop through all values of the array again when we
loop through the keys of the map. This is reducible to O(n) ultimately.

The space complexity is also O(2n) since for each value in the array in the worst case we store a map with two pointers. This is also reducible to O(n) and ends up with a O(n) space complexity for this solution.

## Optimizations
An optimization we could make for this algorithm would help to reduce the time complexity in the worst-case scenario.  If we add a boolean value in the function of hasDuplicates that we flip when a duplicate is found, we can choose not to
loop through the array if that boolean is false.  This would keep the time complexity of O(n) but marginally increase the amount of time it takes due to not needing to loop through the hashmap a second time.

Another optimization is recorded in the firstDuplicateOptimized function. In the first solution I wrote to the firstDuplicateMinimumIndex problem, I utilized a solution I learned from a previous duplicate problem where I needed to return the first occurrence of a duplicate  
character in a string (or array). In that particular question, you must loop through the array (or string) once first in order to discover which characters are duplicates then loop through again in sequence to find the first occurrence of a duplicate.
In the current question, although the sequence is important to the solution, we are tasked to find the second occurrence of the element instead of the first. We can find this in one loop through the array (or string) with only recording the occurrences of a particular element in the array.
This saves us space complexity and time complexity.
So the map ends up looking like this:
```
{
    int: int,
    int: int,
}
```

In the optimized solution, our algorithm can solve the problem in O(n) time complexity and O(n) space complexity right off the bat.
Save those frames friends!



