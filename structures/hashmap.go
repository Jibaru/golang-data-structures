package structures

import (
	"errors"
	"fmt"
	"hash/fnv"
)

const initialBucketsSize = 100

var (
	ErrHashMapKeyNotFound = errors.New("hash map key not found")
)

// hashMap struct
type hashMap[K comparable, V comparable] struct {
	buckets [][]entry[K, V]
	size    int
}

// entry struct
type entry[K comparable, V comparable] struct {
	key   K
	value V
}

// NewHashMap creates a new emoty HashMap
func NewHashMap[K comparable, V comparable]() HashMapper[K, V] {
	return newHashMap[K, V](initialBucketsSize)
}

// newHashMap creates a new HashMap with an initial bucket size
func newHashMap[K comparable, V comparable](initialBuckets int) *hashMap[K, V] {
	buckets := make([][]entry[K, V], initialBuckets)
	return &hashMap[K, V]{
		buckets: buckets,
		size:    0,
	}
}

// hash function using FNV-1a algorithm
func hash(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32())
}

// bucketIndex calculates the bucket index
func (hm *hashMap[K, V]) bucketIndex(key K) int {
	keyStr := fmt.Sprintf("%v", key)
	return hash(keyStr) % len(hm.buckets)
}

// resize increases the number of buckets and rehashes the entries
func (hm *hashMap[K, V]) resize(newBuckets int) {
	newHashMap := newHashMap[K, V](newBuckets)
	for _, bucket := range hm.buckets {
		for _, e := range bucket {
			newHashMap.PushAt(e.key, e.value)
		}
	}
	hm.buckets = newHashMap.buckets
}

// PushAt adds a key-value pair to the HashMap
func (hm *hashMap[K, V]) PushAt(key K, value V) {
	if hm.size > len(hm.buckets)*2 { // load factor > 2, every bucket has more than 2 entries
		hm.resize(len(hm.buckets) * 2)
	}
	index := hm.bucketIndex(key)
	for i, e := range hm.buckets[index] {
		if e.key == key {
			hm.buckets[index][i].value = value
			return
		}
	}
	hm.buckets[index] = append(hm.buckets[index], entry[K, V]{key, value})
	hm.size++
}

// PopAt removes a key-value pair from the HashMap and returns the value
func (hm *hashMap[K, V]) PopAt(key K) (V, error) {
	index := hm.bucketIndex(key)
	for i, e := range hm.buckets[index] {
		if e.key == key {
			val := e.value
			hm.buckets[index] = append(hm.buckets[index][:i], hm.buckets[index][i+1:]...)
			hm.size--
			if hm.size < len(hm.buckets)/4 && len(hm.buckets) > 8 { // load factor < 0.25, every bucket has less than 0.25 entries
				hm.resize(len(hm.buckets) / 2)
			}
			return val, nil
		}
	}
	var zero V
	return zero, ErrHashMapKeyNotFound
}

// Has checks if a key exists in the HashMap
func (hm *hashMap[K, V]) Has(key K) bool {
	index := hm.bucketIndex(key)
	for _, e := range hm.buckets[index] {
		if e.key == key {
			return true
		}
	}
	return false
}

// GetAt retrieves the value associated with a key without removing it
func (hm *hashMap[K, V]) GetAt(key K) (V, error) {
	index := hm.bucketIndex(key)
	for _, e := range hm.buckets[index] {
		if e.key == key {
			return e.value, nil
		}
	}
	var zero V
	return zero, ErrHashMapKeyNotFound
}

// Find searches for a value in the HashMap
func (hm *hashMap[K, V]) Find(value V) bool {
	for _, bucket := range hm.buckets {
		for _, e := range bucket {
			if e.value == value {
				return true
			}
		}
	}
	return false
}

// Size returns the number of key-value pairs in the HashMap
func (hm *hashMap[K, V]) Size() int64 {
	return int64(hm.size)
}
