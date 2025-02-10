Problem: LRU Cache Implementation

Implement a Least Recently Used (LRU) Cache using Java. The cache should support the following operations:
	•	get(int key): Retrieve the value of the key if it exists in the cache, otherwise return -1.
	•	put(int key, int value): Insert or update the value of the key. If the cache reaches its capacity, it should remove the least recently used item before inserting a new item.


    A doubly linked list is a good data structure for this problem. A map is there so that gets and puts are O(1).