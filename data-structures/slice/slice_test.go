package slice

// implement of test-cases for initialization empty slice, Append, Remove, Get, resize, Capacity, Size, Print etc.
import (
	"fmt"
	"testing"
)

// new slice initalization implementation
func TestNew(t *testing.T) {
	// create a new empty generic slice of int type
	size := 0
	capacity := 3
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
	//initialSize := nums.Size()

	nums.Append(0)
	nums.Append(1)
	nums.Append(2)
	nums.Append(3)
	//nums.size++

	newSize := nums.Size()

	// if newSize != initialSize+1 {
	// 	t.Errorf("Expected size after append should be %d, but got %d", len(nums.array), nums.size)
	// }

	// Verify that the last element is the value appended
	lastElement, err := nums.Get(newSize - 1)
	if err != nil {
		t.Fatalf("Error retrieving last element: %v", err)
	}

	if lastElement != 3 {
		t.Errorf("Expected last element to be 0, but got %d", lastElement)
	}
	fmt.Println("after append capacity => ", nums.Capacity())
	nums.print()
}

func TestRemove(t *testing.T) {
	nums := NewSlice[int](0, 0)
	nums.Append(10)
	nums.Append(20)
	nums.Append(30)

	err := nums.Remove(1) // removing eleemnt at index 1
	if err != nil {
		t.Fatalf("Unexpected error during removal %v", err)
	}

	if nums.Size() != 2 {
		t.Errorf("Expected size of 2 after removal, but got %d", nums.Size())
	}

	// Verify that the element at index 1 is now 30
	val, err := nums.Get(1)
	if err != nil {
		t.Fatalf("Error retrieving element: %v", err)
	}

	if val != 30 {
		t.Errorf("Expected element at index 1 to be 30, but got %d", val)
	}

	// case 2 removing first elem
	err = nums.Remove(0)
	if err != nil {
		t.Fatalf("Unexpected error during removal: %v", err)
	}

	if nums.Size() != 1 {
		t.Errorf("Expected size 1 after removal, but got %d", nums.Size())
	}

	// Verify that the element at index 0 is now 30
	val, err = nums.Get(0)
	if err != nil {
		t.Fatalf("Error retrieving element: %v", err)
	}

	if val != 30 {
		t.Errorf("Expected element at index 0 to be 30, but got %d", val)
	}

	// remove last elem
	err = nums.Remove(0)
	if err != nil {
		t.Fatalf("Unexpected error during removal: %v", err)
	}

	if nums.Size() != 0 {
		t.Errorf("Expected size 0 after removal, but got %d", nums.Size())
	}

	// Case 4: Attempt to remove an element from an empty slice
	err = nums.Remove(0)
	if err == nil {
		t.Errorf("Expected error when removing from empty slice, but got nil")
	}

	// Case 5: Attempt to remove with an out-of-bounds index
	nums.Append(40)
	err = nums.Remove(5) // Trying to remove element at index 5 which does not exist
	if err == nil {
		t.Errorf("Expected error when removing with an out-of-bounds index, but got nil")
	}
}

// Insert an element at the beginning of the slice.
// Insert an element in the middle of the slice.
// Insert an element at the end of the slice.
// Attempt to insert an element with an out-of-bounds index.
func TestInsert(t *testing.T) {
	size := 0
	capacity := 10

	nums := NewSlice[int](size, capacity)
	nums.Append(10)
	nums.Append(20)
	nums.Append(30)

	//case 1 insert at the begining
	err := nums.Insert(0, 5) //insert 5 at 0 index
	if err != nil {
		t.Fatalf("unexpected error occured during insert %v", err)
	}

	if nums.Size() != 4 {
		t.Errorf("expected size should be 4 after insert at the begining, but gpt %d", nums.Size())
	}

	//verify the number at index 0, it should be 5
	val, err := nums.Get(0)
	if err != nil {
		t.Fatalf("Error retrieving element: %v", err)
	}

	if val != 5 {
		t.Errorf("Expected element at index 0 to be 5, but got %d", val)
	}

	// case 2 insert in the middle
	err = nums.Insert(2, 15) // inserting 15 at index 2
	if err != nil {
		t.Fatalf("Unexpected error during insert in the middle: %v", err)
	}

	if nums.Size() != 5 {
		t.Errorf("Expected size 5 after insert in the middle, but got %d", nums.Size())
	}

	// Verify that the element at index 2 is now 15
	val, err = nums.Get(2)
	if err != nil {
		t.Fatalf("Error retrieving element: %v", err)
	}

	if val != 15 {
		t.Errorf("Expected element at index 2 to be 15, but got %d", val)
	}

	// case 3 insert at the end
	err = nums.Insert(nums.Size(), 35)
	if err != nil {
		t.Fatalf("Unexpected error during insert at the end: %v", err)
	}

	if nums.Size() != 6 {
		t.Errorf("expected size to be 6 after insert at the end, but got %d", nums.Size())
	}

	// Verify that the last element is now 35
	val, err = nums.Get(nums.Size() - 1)
	if err != nil {
		t.Fatalf("Error retrieving element: %v", err)
	}

	if val != 35 {
		t.Errorf("Expected last element to be 35, but got %d", val)
	}

	// Case 4: Attempt to insert with an out-of-bounds index
	err = nums.Insert(10, 40) // Trying to insert at index 10 which is out of bounds
	if err == nil {
		t.Errorf("Expected out-of-bounds error, but got nil")
	}
}

// clear Removes all elements
func TestClear(t *testing.T) {
	nums := NewSlice[int](0, 0)
	nums.Append(10)
	nums.Append(20)
	nums.Append(30)
	nums.Append(40)

	nums.Clear()

	if nums.Size() != 0 {
		t.Errorf("expected size to be 0 after clear, got %d", nums.Size())
	}

	if len(nums.array) != 0 {
		t.Errorf("expected internal array length be 0 after clear, got %d", len(nums.array))
	}
}

func TestContains(t *testing.T) {
	nums := NewSlice[int](0, 0)
	nums.Append(10)
	nums.Append(20)
	nums.Append(30)

	if !nums.Contains(20) {
		t.Errorf("expected slice sto contain 20")
	}

	if nums.Contains(100) {
		t.Errorf("Expected slice not to contain 100")
	}
}

func TestFind(t *testing.T) {
	nums := NewSlice[int](0, 10)
	nums.Append(10)
	nums.Append(20)
	nums.Append(30)

	index := nums.Find(20)
	if index != 1 {
		t.Errorf("expected index of 20 to be 1, got %d", index)
	}

	index = nums.Find(100)
	if index != -1 {
		t.Errorf("expected index of 100 to be -1, got %d", index)
	}
}

// Appends elements from another slice.
func TestExtend(t *testing.T) {
	nums := NewSlice[int](0, 0)
	nums.Append(10)
	nums.Append(20)

	anotherNums := NewSlice[int](0, 0)
	anotherNums.Append(30)
	anotherNums.Append(40)

	nums.Extend(anotherNums)

	if nums.Size() != 4 {
		t.Errorf("expected size to be 4 after extend, got %d", nums.Size())
	}

	val, err := nums.Get(2)
	if err != nil || val != 30 {
		t.Errorf("Expected element at index 2 to be 30 but got %d", val)
	}

	val, err = nums.Get(3)
	if err != nil || val != 40 {
		t.Errorf("Expected element at index 3 to be 40, but got %d", val)
	}
}

func TestReverse(t *testing.T) {
	nums := NewSlice[int](0, 0)
	nums.Append(10)
	nums.Append(20)
	nums.Append(30)
	nums.Append(40)

	nums.Reverse()

	// check if elem are reversed
	expected := []int{40, 30, 20, 10}
	for i := 0; i < nums.Size(); i++ {
		val, err := nums.Get(i)
		if err != nil || val != expected[i] {
			t.Errorf("Expected element at index %d to be %d, but got %d", i, expected[i], val)
		}
	}
}

func TestCopy(t *testing.T) {
	nums := NewSlice[int](0, 100)
	nums.Append(10)
	nums.Append(20)
	nums.Append(30)

	copiedNum := nums.Copy()

	if copiedNum.Size() != nums.Size() {
		t.Errorf("expected copied size and num size to be same fater creating new copy, should have been %d but got %d\n", nums.Size(), copiedNum.Size())
	}

	//checking the elements in the copied slice
	for i := 0; i < nums.Size(); i++ {
		originalVal, _ := nums.Get(i)
		copiedVal, _ := copiedNum.Get(i)

		if originalVal != copiedVal {
			t.Errorf("Expected copied element at index %d to be %d, but got %d", i, originalVal, copiedVal)
		}
	}

	// modify the copied slice & check if the original slice is also mutated?
	copiedNum.Append(40)
	if copiedNum.Size() == nums.Size() {
		t.Errorf("Expected copied slice size to differ after modification, but they are the same")

	}
}
