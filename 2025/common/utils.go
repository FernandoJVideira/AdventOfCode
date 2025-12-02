package common

// Set represents a set of integers
type Set map[int]struct{}

// NewSet creates a new set
func NewSet(values ...int) Set {
	s := make(Set)
	for _, v := range values {
		s.Add(v)
	}
	return s
}

// Add adds a value to the set
func (s Set) Add(val int) {
	s[val] = struct{}{}
}

// Remove removes a value from the set
func (s Set) Remove(val int) {
	delete(s, val)
}

// Contains checks if a value is in the set
func (s Set) Contains(val int) bool {
	_, exists := s[val]
	return exists
}

// Size returns the number of elements in the set
func (s Set) Size() int {
	return len(s)
}

// ToSlice converts the set to a slice
func (s Set) ToSlice() []int {
	slice := make([]int, 0, len(s))
	for val := range s {
		slice = append(slice, val)
	}
	return slice
}

// StringSet represents a set of strings
type StringSet map[string]struct{}

// NewStringSet creates a new string set
func NewStringSet(values ...string) StringSet {
	s := make(StringSet)
	for _, v := range values {
		s.Add(v)
	}
	return s
}

// Add adds a value to the set
func (s StringSet) Add(val string) {
	s[val] = struct{}{}
}

// Remove removes a value from the set
func (s StringSet) Remove(val string) {
	delete(s, val)
}

// Contains checks if a value is in the set
func (s StringSet) Contains(val string) bool {
	_, exists := s[val]
	return exists
}

// Size returns the number of elements in the set
func (s StringSet) Size() int {
	return len(s)
}

// ToSlice converts the set to a slice
func (s StringSet) ToSlice() []string {
	slice := make([]string, 0, len(s))
	for val := range s {
		slice = append(slice, val)
	}
	return slice
}

// PointSet represents a set of 2D points
type PointSet map[Point]struct{}

// NewPointSet creates a new point set
func NewPointSet(points ...Point) PointSet {
	s := make(PointSet)
	for _, p := range points {
		s.Add(p)
	}
	return s
}

// Add adds a point to the set
func (s PointSet) Add(p Point) {
	s[p] = struct{}{}
}

// Remove removes a point from the set
func (s PointSet) Remove(p Point) {
	delete(s, p)
}

// Contains checks if a point is in the set
func (s PointSet) Contains(p Point) bool {
	_, exists := s[p]
	return exists
}

// Size returns the number of points in the set
func (s PointSet) Size() int {
	return len(s)
}

// Queue represents a FIFO queue
type Queue[T any] struct {
	items []T
}

// NewQueue creates a new queue
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{items: make([]T, 0)}
}

// Enqueue adds an item to the back of the queue
func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the front item
func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

// IsEmpty returns true if the queue is empty
func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns the number of items in the queue
func (q *Queue[T]) Size() int {
	return len(q.items)
}

// Stack represents a LIFO stack
type Stack[T any] struct {
	items []T
}

// NewStack creates a new stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{items: make([]T, 0)}
}

// Push adds an item to the top of the stack
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

// Peek returns the top item without removing it
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

// IsEmpty returns true if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of items in the stack
func (s *Stack[T]) Size() int {
	return len(s.items)
}
