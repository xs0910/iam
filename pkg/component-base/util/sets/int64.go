package sets

import (
	"reflect"
	"sort"
)

// Int64 is a set of int64s, implemented via map[int64]struct{} for minimal memory consumption.
type Int64 map[int64]Empty

// NewInt64 creates an Int64 from a list of values.
func NewInt64(items ...int64) Int64 {
	ss := Int64{}
	ss.Insert(items...)
	return ss
}

// Int64KeySet creates an Int64 from a keys of a map[int64](? extends interface{}).
// If the value passed in is not actually a map, this will panic.
func Int64KeySet(theMap interface{}) Int64 {
	v := reflect.ValueOf(theMap)
	ret := Int64{}

	for _, keyValue := range v.MapKeys() {
		ret.Insert(keyValue.Interface().(int64))
	}
	return ret
}

// Insert adds items to the set.
func (s1 Int64) Insert(items ...int64) Int64 {
	for _, item := range items {
		s1[item] = Empty{}
	}
	return s1
}

// Delete removes all items from the set.
func (s1 Int64) Delete(items ...int64) Int64 {
	for _, item := range items {
		delete(s1, item)
	}
	return s1
}

// Has returns true if and only if item is contained in the set.
func (s1 Int64) Has(item int64) bool {
	_, contained := s1[item]
	return contained
}

// HasAll returns true if and only if all items are contained in the set.
func (s1 Int64) HasAll(items ...int64) bool {
	for _, item := range items {
		if !s1.Has(item) {
			return false
		}
	}
	return true
}

// HasAny returns true if any items are contained in the set.
func (s1 Int64) HasAny(items ...int64) bool {
	for _, item := range items {
		if s1.Has(item) {
			return true
		}
	}
	return false
}

// Difference returns a set of objects that are not in s2
// For example:
// s1 = {a1, a2, a3}
// s2 = {a1, a2, a4, a5}
// s1.Difference(s2) = {a3}
// s2.Difference(s1) = {a4, a5}
func (s1 Int64) Difference(s2 Int64) Int64 {
	result := NewInt64()
	for key := range s1 {
		if !s2.Has(key) {
			result.Insert(key)
		}
	}
	return result
}

// Union returns a new set which includes items in either s1 or s2.
// For example:
// s1 = {a1, a2}
// s2 = {a3, a4}
// s1.Union(s2) = {a1, a2, a3, a4}
// s2.Union(s1) = {a1, a2, a3, a4}
func (s1 Int64) Union(s2 Int64) Int64 {
	result := NewInt64()
	for key := range s1 {
		result.Insert(key)
	}
	for key := range s2 {
		result.Insert(key)
	}
	return result
}

// Intersection returns a new set which includes the item in BOTH s1 and s2
// For example:
// s1 = {a1, a2}
// s2 = {a2, a3}
// s1.Intersection(s2) = {a2}
func (s1 Int64) Intersection(s2 Int64) Int64 {
	var walk, other Int64
	result := NewInt64()
	if s1.Len() < s2.Len() {
		walk = s1
		other = s2
	} else {
		walk = s2
		other = s1
	}
	for key := range walk {
		if other.Has(key) {
			result.Insert(key)
		}
	}
	return result
}

// IsSuperset returns true if and only if s1 is a superset of s2.
func (s1 Int64) IsSuperset(s2 Int64) bool {
	for item := range s2 {
		if !s1.Has(item) {
			return false
		}
	}
	return true
}

// Equal returns true if and only if s1 is equal (as a set) to s2.
// Two sets are equal if their membership is identical.
// (In practice, this means same elements, order doesn't matter)
func (s1 Int64) Equal(s2 Int64) bool {
	return len(s1) == len(s2) && s1.IsSuperset(s2)
}

type sortableSliceOfInt64 []int64

func (s sortableSliceOfInt64) Len() int           { return len(s) }
func (s sortableSliceOfInt64) Less(i, j int) bool { return lessInt64(s[i], s[j]) }
func (s sortableSliceOfInt64) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// List returns the contents as a sorted int64 slice.
func (s1 Int64) List() []int64 {
	res := make(sortableSliceOfInt64, 0, len(s1))
	for key := range s1 {
		res = append(res, key)
	}
	sort.Sort(res)
	return res
}

// UnsortedList returns the slice with contents in random order.
func (s1 Int64) UnsortedList() []int64 {
	res := make([]int64, 0, len(s1))
	for key := range s1 {
		res = append(res, key)
	}
	return res
}

// PopAny Returns a single element from the set.
func (s1 Int64) PopAny() (int64, bool) {
	for key := range s1 {
		s1.Delete(key)
		return key, true
	}
	var zeroValue int64
	return zeroValue, false
}

// Len returns the size of the set.
func (s1 Int64) Len() int {
	return len(s1)
}

func lessInt64(lhs, rhs int64) bool {
	return lhs < rhs
}
