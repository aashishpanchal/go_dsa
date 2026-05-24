package linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingly(t *testing.T) {
	t.Run("PushBack and PushFront", func(t *testing.T) {
		list := NewSingly[int]()
		list.PushBack(1)
		list.PushBack(2)
		list.PushFront(0)
		assert.Equal(t, "[0,1,2]", list.String())
		assert.Equal(t, 3, list.size)
	})

	t.Run("PopFront", func(t *testing.T) {
		list := NewSingly(1, 2)
		list.PopFront()
		assert.Equal(t, "[2]", list.String())
		list.PopFront()
		assert.Equal(t, "[]", list.String())
		list.PopFront() // pop empty
		assert.Equal(t, 0, list.size)
	})

	t.Run("PopBack", func(t *testing.T) {
		list := NewSingly(1, 2, 3)
		list.PopBack()
		assert.Equal(t, "[1,2]", list.String())
		list.PopBack()
		assert.Equal(t, "[1]", list.String())
		list.PopBack()
		assert.Equal(t, "[]", list.String())
		list.PopBack() // pop empty
		assert.Equal(t, 0, list.size)
	})

	t.Run("Reverse", func(t *testing.T) {
		list := NewSingly(1, 2, 3)
		list.Reverse()
		assert.Equal(t, "[3,2,1]", list.String())
		// check if head and tail are updated correctly
		front, _ := list.Front()
		back, _ := list.Back()
		assert.Equal(t, 3, front)
		assert.Equal(t, 1, back)
	})

	t.Run("Insert", func(t *testing.T) {
		list := NewSingly[int]()
		list.Insert(0, 2) // push front
		list.Insert(1, 4) // push back
		list.Insert(1, 3) // insert middle
		assert.Equal(t, "[2,3,4]", list.String())
		err := list.Insert(10, 5)
		assert.Error(t, err)
	})

	t.Run("Remove", func(t *testing.T) {
		list := NewSingly(1, 2, 3)
		err := list.Remove(1) // remove middle
		assert.NoError(t, err)
		assert.Equal(t, "[1,3]", list.String())
		list.Remove(0) // remove head
		assert.Equal(t, "[3]", list.String())
		list.Remove(0) // remove tail
		assert.Equal(t, "[]", list.String())
		err = list.Remove(0)
		assert.Error(t, err)
	})

	t.Run("Get and Set", func(t *testing.T) {
		list := NewSingly(10, 20)
		val, err := list.Get(1)
		assert.NoError(t, err)
		assert.Equal(t, 20, val)
		err = list.Set(1, 25)
		assert.NoError(t, err)
		val, _ = list.Get(1)
		assert.Equal(t, 25, val)
		_, err = list.Get(5)
		assert.Error(t, err)
	})

	t.Run("Clear", func(t *testing.T) {
		list := NewSingly(1)
		list.Clear()
		assert.Equal(t, 0, list.size)
		assert.Equal(t, "[]", list.String())
	})

	t.Run("Empty List Operations", func(t *testing.T) {
		list := NewSingly[int]()
		// Pop operations on empty list should not panic
		list.PopFront()
		list.PopBack()
		assert.Equal(t, 0, list.size)
		// Accessors should return error
		_, err := list.Front()
		assert.Error(t, err)
		_, err = list.Back()
		assert.Error(t, err)
		_, err = list.Get(0)
		assert.Error(t, err)
	})
}
