package linked_list

import (
	"fmt"
)

type Singly[T comparable] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func NewSingly[T comparable](data ...T) *Singly[T] {
	s := &Singly[T]{size: 0}
	for _, v := range data {
		s.PushBack(v)
	}
	return s
}

func (s *Singly[T]) PushBack(value T) {
	newNode := NewNode(value, nil)
	if s.tail == nil {
		s.head = newNode
		s.tail = newNode
	} else {
		s.tail.next = newNode
		s.tail = newNode
	}
	s.size++
}

func (s *Singly[T]) PushFront(value T) {
	newNode := NewNode(value, nil)
	newNode.next = s.head
	s.head = newNode
	if s.tail == nil {
		s.tail = newNode
	}
	s.size++
}

func (s *Singly[T]) PopFront() {
	if s.head == nil {
		return
	}
	s.head = s.head.next
	if s.head == nil {
		s.tail = nil
	}
	s.size--
}

func (s *Singly[T]) PopBack() {
	if s.head == nil {
		return
	}
	if s.head == s.tail {
		s.head = nil
		s.tail = nil
		s.size = 0
		return
	}
	curr := s.head
	for curr.next != s.tail {
		curr = curr.next
	}
	curr.next = nil
	s.tail = curr
	s.size--
}

func (s *Singly[T]) Reverse() {
	var prev *Node[T]
	curr := s.head
	s.tail = s.head
	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	s.head = prev
}

func (s *Singly[T]) Clear() {
	s.head = nil
	s.tail = nil
	s.size = 0
}

func (s *Singly[T]) Insert(index int, value T) error {
	if index > s.size || index < 0 {
		return fmt.Errorf("index out of range")
	}
	if index == 0 {
		s.PushFront(value)
		return nil
	}
	if index == s.size {
		s.PushBack(value)
		return nil
	}
	curr := s.head
	for i := 0; i < index-1; i++ {
		curr = curr.next
	}
	newNode := NewNode(value, curr.next)
	curr.next = newNode
	s.size++
	return nil
}

func (s *Singly[T]) Remove(index int) error {
	if index >= s.size || index < 0 {
		return fmt.Errorf("index out of range")
	}
	if index == 0 {
		s.PopFront()
		return nil
	}
	curr := s.head
	for i := 0; i < index-1; i++ {
		curr = curr.next
	}
	target := curr.next
	curr.next = target.next
	if target == s.tail {
		s.tail = curr
	}
	s.size--
	return nil
}

func (s *Singly[T]) Front() (T, error) {
	if s.head == nil {
		var zero T
		return zero, fmt.Errorf("list is empty")
	}
	return s.head.data, nil
}

func (s *Singly[T]) Back() (T, error) {
	if s.tail == nil {
		var zero T
		return zero, fmt.Errorf("list is empty")
	}
	return s.tail.data, nil
}

func (s *Singly[T]) Get(index int) (T, error) {
	if index >= s.size || index < 0 {
		var zero T
		return zero, fmt.Errorf("index out of range")
	}
	curr := s.head
	for i := 0; i < index; i++ {
		curr = curr.next
	}
	return curr.data, nil
}

func (s *Singly[T]) Set(index int, value T) error {
	if index >= s.size || index < 0 {
		return fmt.Errorf("index out of range")
	}
	curr := s.head
	for i := 0; i < index; i++ {
		curr = curr.next
	}
	curr.data = value
	return nil
}

func (s *Singly[T]) String() string {
	return s.ToStr(10)
}

func (s *Singly[T]) ToStr(limit int) string {
	if s.head == nil {
		return "[]"
	}
	return s.head.ToStr(limit)
}

func (s *Singly[T]) ForEach(cb func(value T, index int)) {
	curr := s.head
	index := 0
	for curr != nil {
		cb(curr.data, index)
		curr = curr.next
		index++
	}
}
