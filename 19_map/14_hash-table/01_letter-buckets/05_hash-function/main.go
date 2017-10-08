/*
https://stackoverflow.com/questions/9073903/what-does-bucket-entries-mean-in-the-context-of-a-hashtable

Hash table or hash map is a data structure that uses a hash function to map identifying values,
known as keys (e.g., a person's name), to their associated values (e.g., their telephone number).
Thus, a hash table implements an associative array.
The hash function is used to transform the key into the index (the hash) of an array element (the slot or bucket)
where the corresponding value is to be sought.
A bucket is simply a fast-access location (like an array index) that is the the result of the hash function.

TL;DR
It is like an array index and it is the result of the hash function.
*/
package main

import "fmt"

func main() {
	n := hashBucket("Go", 12)
	fmt.Println(n) // the hash for the bucket is 11
}

func hashBucket(word string, buckets int) int {
	// we could use letter := rune(word[0])
	// in that case we will need to change the types as
	// buckets int32 and return type int32
	letter := int(word[0])
	bucket := letter % buckets
	return bucket
}
