package slice

import (
	"errors"
	"fmt"
	"reflect"
)

// we will try to implement custom slice & its functionality based on the top of array.
// implement initialization of empty slice, Append, Remove, Get, resize, Capacity, Size, Print etc with test-cases.

type Slice[T any] struct {
	array    []T
	size     int
	capacity int
}

// new slice initialization of generic type and returns a generic type of slice
func NewSlice[T any](size, capacity int) *Slice[T] {
	return &Slice[T]{
		array:    make([]T, size, capacity),
		size:     size,
		capacity: capacity,
	}
}

// get the element at the index, return value of type T and error
func (s *Slice[T]) Get(index int) (T, error) {
	if index < 0 || index > s.size {
		var zeroValue T
		return zeroValue, errors.New("index out of bound")
	}
	return s.array[index], nil
}

// append data to slice to add element at the end
func (s *Slice[T]) Append(value T) {
	if s.size == s.capacity {
		fmt.Println("size & capacity is same, full capacity is uzed, append failed")
		s.resize()
	}
	s.array = append(s.array, value)
	s.size++
	s.print()
}

// resize function, to increase the capacity of the slice to double when capacity == size
func (s *Slice[T]) resize() {
	newCapacity := s.capacity * 2
	if newCapacity == 0 {
		newCapacity = 1
	}

	newArr := make([]T, s.size, newCapacity)
	copy(newArr, s.array)
	s.array = newArr
	s.capacity = newCapacity
}

// print to display the slice contents
func (s *Slice[T]) print() {
	fmt.Printf("slice elements %v\n", s.array)
}

// size & capacity func
func (s *Slice[T]) Size() int {
	return s.size
}

func (s *Slice[T]) Capacity() int {
	return s.capacity
}

// Remove removes an element at a specific index and shifts elements left
func (s *Slice[T]) Remove(index int) error {
	if index < 0 || index >= s.size {
		return errors.New("index out of bound")
	}

	s.array = append(s.array[:index], s.array[index+1:]...)
	s.size--

	return nil
}

// Insert: Adds an element at a specific index.
func (s *Slice[T]) Insert(index int, Value T) error {

	if index < 0 || index > s.size {
		return errors.New("index out of bounds")
	}

	if s.size == s.capacity {
		s.resize()
	}

	s.array = append(s.array[:index], append([]T{Value}, s.array[index:]...)...)
	s.size++
	return nil
}

// Clear: Removes all elements without deallocating the underlying array.
func (s *Slice[T]) Clear() {

	s.array = s.array[:0]
	s.size = 0
}

// Contains: Checks if a value exists in the slice.
func (s *Slice[T]) Contains(value T) bool {

	for _, v := range s.array {
		// we are using reflect package since our value & slice is of generic Type. to avoid using reflect, we can add the constraint to slice struct called "comparable", but it will restrict slice funtionality to other ADT like maps, etc
		if reflect.DeepEqual(v, value) {
			return true
		}
	}

	return false
}

// Find: Finds the index of a specific element.
func (s *Slice[T]) Find(value T) int {
	for i, v := range s.array {
		if reflect.DeepEqual(v, value) {
			return i
		}
	}
	return -1
}

// Extend: Appends elements from another slice.
func (s *Slice[T]) Extend(otherArr *Slice[T]) {
	for _, v := range otherArr.array {
		s.Append(v)
	}
}

// Reverse: Reverses the order of elements.
func (s *Slice[T]) Reverse() {

	for i, j := 0, len(s.array)-1; i < j; i, j = i+1, j-1 {
		s.array[i], s.array[j] = s.array[j], s.array[i]
	}
	s.print()
}

// Copy: Creates a deep copy of the slice.
func (s *Slice[T]) Copy() *Slice[T] {
	newSlice := NewSlice[T](s.size, s.capacity)
	copy(newSlice.array, s.array)
	return newSlice
}
