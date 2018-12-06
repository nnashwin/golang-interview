# LruCache

## Problem Definition
 Design and implement a data structure for Least Recently Used (LRU) cache. It should support the following operations: get and put.

get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
put(key, value) - Set or insert the value if the key is not already present. When the cache reached its capacity, it should invalidate the least recently used item before inserting a new item.

Follow up:
Could you do both operations in O(1) time complexity?

Example:

```
LRUCache cache = new LRUCache( 2 /* capacity */ );

cache.put(1, 1);
cache.put(2, 2);
cache.get(1);       // returns 1
cache.put(3, 3);    // evicts key 2
cache.get(2);       // returns -1 (not found)
cache.put(4, 4);    // evicts key 1
cache.get(1);       // returns -1 (not found)
cache.get(3);       // returns 3
cache.get(4);       // returns 4
```

## Questions to Ask
1. Do we just put ints in the cache?  Should it handle other values?
2. Do we need a specific space complexity?

## Complexity
- The Time complexity of all the operations is O(1).  We achieve this by using a Head pointer, tail pointer, a list for the actual cache, and a hashmap of the individual nodes in the cache.
- Since the Nodes are all linked together, we can remove the head node and move the pointer to its next in order to effectively delete the node (while removing it from the hashmap to all the nodes).
- The space complexity is O(2N), where N is the length of the nodes in the cache.  This is because we maintain the map of Nodes and the actual doubly linked list linking all the nodes together.
