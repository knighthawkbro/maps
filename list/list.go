package list

import (
	"fmt"
)

// node (Private) - Defines the structure for each individual node in a linked list
type node struct {
	key   interface{} // key of node
	value interface{} // Value of Node
	next  *node       // Pointer to the next Node
	list  *List       // Pointer to the list it is attached to
}

func (n node) String() string {
	return fmt.Sprintf("%v=%v", n.key, n.value)
}

// nextNode (Private) - Returns the next node in the list
func (n node) nextNode() *node {
	// returns nil if there is not list AND if the pointer to the next
	// node is the same as the head's next node there for there is next node
	if next := n.next; n.list != nil && next != &n.list.head {
		return next
	}
	return nil
}

// List (Public) - The container for all the linked nodes in a set
type List struct {
	head node // the begining node
	size int  // size of the list
}

// init (Private) - Generates a linked list with Size=0 and head pointing to itself
func (l *List) init() *List {
	l.head.next = &l.head
	l.size = 0
	return l
}

// New (Public) - Returns an initialized list.
func New() *List { return new(List).init() }

// Add (Public) - if key in map, replaces old value with new and returns old value
// otherwise adds key value pair and returns nil
func (l *List) Add(key, value interface{}) interface{} {
	if key == nil || value == nil {
		return nil
	}
	current := l.head.next
	for i := 0; i < l.size; i++ {
		if current.key == key {
			removed := current.value
			current.value = value
			return removed
		}
		current = current.next
	}
	new := &node{key: key, value: value, list: l}
	prev := l.head.next
	new.next = prev
	l.head.next = new
	l.size++
	return nil
}

// Remove (Public) - removes key value pair from map and returns value
// if key not found, returns nil
func (l *List) Remove(key interface{}) interface{} {
	if key == nil {
		return nil
	}
	current := &l.head
	for i := 0; i < l.size; i++ {
		if current.next.key == key {
			removed := current.next.value
			current.next = current.next.next
			l.size--
			return removed
		}
		current = current.next
	}
	return nil
}

// Get (Public) - returns value associated with key
// if key not found, returns null
func (l *List) Get(key interface{}) interface{} {
	if key == nil {
		return nil
	}
	current := l.head.next
	for i := 0; i < l.size; i++ {
		if current.key == key {
			return current.value
		}
		current = current.next
	}
	return nil
}

// Contains (Public) - returns true if key is in map
func (l *List) Contains(key interface{}) bool {
	if key == nil {
		return false
	}
	current := l.head.next
	for i := 0; i < l.size; i++ {
		if current.key == key {
			return true
		}
		current = current.next
	}
	return false
}

// GetKeys (Public) -
func (l *List) GetKeys() string {
	list := "[ "
	current := l.head.next
	for i := 0; i < l.size; i++ {
		list += fmt.Sprintf("%v ", current.key)
		current = current.next
	}
	return list + "]"
}

// Size (Public) - Returns the size of the list
func (l *List) Size() int {
	return l.size
}

// String (Public) - Allows for the fmt.Print* functions to print the list struct as a string.
func (l *List) String() string {
	if l.size == 0 {
		return "{}"
	}
	result := "{ "
	for current := l.head.next; current != nil; current = current.nextNode() {
		result += fmt.Sprintf("%v=%v ", current.key, current.value)
	}
	return result + "}"
}
