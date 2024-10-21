package slice

import (
	"errors"
	"fmt"
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
	}
	s.array = append(s.array, value)
	s.size++
	s.print()
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
