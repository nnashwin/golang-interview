# Group Anagrams 

## Problem Definition
Given an array of strings as input, group all of the anagrams together into their own array.  Return an array of all of the anagram arrays
grouped this way.

Input: myStrings = ["bat", "ant", "tan", "eat", "nat", "ate", "tea"]
Output: [["bat"], ["ant", "tan", "nat"], ["eat", "ate", "tea"]]

Note: A word is an anagram of another if it uses the exact same characters as the other word.

So "ant" and "tan", are anagrams, but "bat" and "eat" are not because of the "b" and "e".

Also, "at" is not an anagram of "ate" because they don't have the exact same characters.

## Explanation
My first intuition was to use a map of string to array of strings (or slice of strings in golang map[string][]string) to store the values, but this isn't the actual brute force way to solve the problem.

The brute force way can loop through the characters in each string and compare the amount of characters it has in the string to each other string we have seen before.

So if we run it on the example above it could look like the following:

```
Input: myStrings = ["bat", "ant", "tan", "eat", "nat", "ate", "tea"]
Output: [["bat"], ["ant", "tan", "nat"], ["eat", "ate", "tea"]]

The answer starts out as an empty array []

As we loop through "bat", we see that there are no groups in the answer, so we can append "bat".

The answer will be [["bat"]] after we see "bat".


For "ant", we start looping through the strings in all of our groups to see if there is a match.

We will first encounter "bat", and we can loop through the characters of "ant" or "bat" to see if they match one another.

They do not, so we will append "ant" to it's own array in the result.

We continue doing this until the problem is complete.

The time complexity of the following will be O(n^2) since we could compare all of the characters in all of the strings for each string in the worst case scenario.

The space complexity should be O(n * l) where n is the length of the input array and l is the length of each string. 
```

There are methods we can use to make the character comparisons in the two strings quicker, but it is much more useful to just use a different data structure to groups of strings instead.

## Optimizations

In a more optimal solution, we use a hash table to store the sorted string characters, sort each string as we loop through and either add it to the hash table when it matches a previous key or create a new entry when it is a new key.

The approach using this method is as follows:

```
Input: myStrings = ["bat", "ant", "tan", "eat", "nat", "ate", "tea"]
Output: [["bat"], ["ant", "tan", "nat"], ["eat", "ate", "tea"]] 

1. Loop through each string in the input.

2. For each string, sort the string {lexographically} and save the sorted result.

3. See if the sorted string is in the hash table.
-- 1. If so, append the unsorted string to this group.
-- 2. If not, create a new group with the unsorted string as its first entry.

4.  Finally, loop through the hash table and append each group to the result array.
```

The time complexity of the solution will depend on which sort you use to implement the character sorting in the strings.

Using a comparison based sort (like the one I used in the solution) will have a O(n log n) time complexity and O(1) space complexity if you use an in-place sort like Quicksort.

The benefits of this type of sort is that it will not have huge performance hits when you change the set of characters used in the input.

If the character set that could be in the set of strings are lower case letters or some other bounded character sit, you can sort a string using [Counting Sort](https://en.wikipedia.org/wiki/Counting_sort) in O(n).

Our algorithm will group anagrams in O(N * L) in this circumstance where the N is the number of strings and the L is the length of each string.

## Complexities
Brute Force comparing arrays to arrays: O(N^2) time and O(N * L) space.

Optimization using comparison based sort: O(N * l log l) time and O(N * L) space.

Optimization using counting sort: O(N * L) time and O(N * length of character set) space.