package slice

// implement of test-cases for initialization empty slice, Append, Remove, Get, resize, Capacity, Size, Print etc.
import (
	"testing"
)

// new slice initalization implementation
func TestNew(t *testing.T) {
	// create a new empty generic slice of int type
	size := 0
	capacity := 10
	nums := NewSlice[int](size, capacity)

	// check if slice is nil
	if nums == nil {
		t.Fatalf("expected non-nil slice, but got nil slice")
	}

	// check if slice size is 0
	if nums.size != size {
		t.Fatalf("expected size of slice to be 0, got %d", nums.Size())
	}

	//check if initial capacity is 0
	if nums.capacity != capacity {
		t.Fatalf("expected capacity of slice to be 0, got %d", nums.Capacity())
	}

	//verify the array is initalized
	if len(nums.array) != 0 {
		t.Fatalf("Expected array length 0, but got %d", len(nums.array))
	}
}

func TestAppend(t *testing.T) {
	// on new value append, the size of the slice should increase & last element be updated

	nums := NewSlice[int](0, 3)
	initialSize := nums.Size()

	nums.Append(0)
	//nums.size++

	newSize := nums.Size()

	if newSize != initialSize+1 {
		t.Errorf("Expected size after append should be %d, but got %d", len(nums.array), nums.size)
	}

	// Verify that the last element is the value appended
	lastElement, err := nums.Get(newSize - 1)
	if err != nil {
		t.Fatalf("Error retrieving last element: %v", err)
	}

	if lastElement != 0 {
		t.Errorf("Expected last element to be 0, but got %d", lastElement)
	}
}
