package array

import (
	"fmt"
)

// pair (private)
type pair struct {
	key   interface{}
	value interface{}
}

// String (Public) - formats the array when fmt.Print is called.
func (p *pair) String() string {
	return fmt.Sprintf("%v=%v", p.key, p.value)
}

// Array (Public) - Structure that defines
type Array struct {
	size       int
	collection []pair
}

// Init (Public) - initializes the array with whatever size is provided, This is what can be overrided by the user.
func (a *Array) Init(capacity int) *Array {
	if capacity < 0 {
		return nil
	}
	a.collection = make([]pair, capacity)
	a.size = 0
	return a
}

// New (Public) - Returns an initialized array with default size of 10.
func New() *Array { return new(Array).Init(10) }

// Add (Public) - if key in map, replaces old value with new and returns old value
// otherwise adds key value pair and returns nil
func (a *Array) Add(key, value interface{}) interface{} {
	err := checkForNil(key, value)
	if err != nil {
		return err
	}
	for i := 0; i < a.size; i++ {
		if a.collection[i].key == key {
			removed := a.collection[i].value
			a.collection[i].value = value
			return removed
		}
	}
	a.ensureSpace()
	a.collection[a.size] = pair{key, value}
	a.size++
	return nil
}

// Remove (Public) - removes key value pair from map and returns value
// if key not found, returns nil
func (a *Array) Remove(key interface{}) interface{} {
	if key == nil {
		return nil
	}
	for i := 0; i < a.size; i++ {
		if a.collection[i].key == key {
			removed := a.collection[i].value
			a.shiftLeft(i)
			a.size--
			return removed
		}
	}
	return nil
}

// Get (public) - returns value stored for key
// if key not found, returns nil
func (a *Array) Get(key interface{}) interface{} {
	if key == nil {
		return nil
	}
	for i := 0; i < a.size; i++ {
		if a.collection[i].key == key {
			return a.collection[i].value
		}
	}
	return nil
}

// Contains (Public) - returns true if key is in map
func (a *Array) Contains(key interface{}) bool {
	for i := 0; i < a.size; i++ {
		if a.collection[i].key == key {
			return true
		}
	}
	return false
}

// GetKeys (Public) -
func (a *Array) GetKeys() string {
	if a.size == 0 {
		return "[ ]"
	}
	s := "[ "
	for i := 0; i < a.size; i++ {
		s += fmt.Sprintf("%v ", a.collection[i].key)
	}
	return s + "]"
}

// Size (Public) - Returns the size of the array
func (a *Array) Size() int {
	return a.size
}

// String (Public) - formats the map as a string
func (a *Array) String() string {
	if a.size == 0 {
		return "{}"
	}
	s := "{ "
	for i := 0; i < a.size; i++ {
		s += fmt.Sprintf("%v ", a.collection[i].String())
	}
	return s + "}"
}

// checkForNil (private) - check if the key or the value is nil
func checkForNil(key, value interface{}) error {
	if key == nil || value == nil {
		return fmt.Errorf("Key/Value can't be nil")
	}
	return nil
}

// ensureSpace (Private) - Sees if the size and capacity of the array are the same. If so,
// It creates a new array with double the capacity and overwrites the old array with a new
// array, then clears the new array for the GC.
func (a *Array) ensureSpace() {
	if a.size == cap(a.collection) {
		new := new(Array).Init(cap(a.collection) * 2)
		new.size = a.size
		for i := 0; i < a.size; i++ {
			new.collection[i] = a.collection[i]
		}
		*a = *new
		new = nil
	}
}

// shiftLeft (Private) - Moves all the items left after index (Destructive)
func (a *Array) shiftLeft(index int) {
	for i := index; i < a.size-1; i++ {
		a.collection[i] = a.collection[i+1]
	}
}
