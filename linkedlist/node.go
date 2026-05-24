package linked_list

import (
	"fmt"
	"strings"
)

type Node[T any] struct {
	data T
	next *Node[T]
}

func NewNode[T comparable](value T, next *Node[T]) *Node[T] {
	return &Node[T]{data: value, next: next}
}

func (n *Node[T]) ToStr(limit int) string {
	curr := n
	var sb strings.Builder
	sb.WriteString("[")
	count := 0
	for curr != nil && count < limit {
		if count > 0 {
			sb.WriteString(",")
		}
		fmt.Fprintf(&sb, "%v", curr.data)
		curr = curr.next
		count++
	}
	if curr != nil {
		sb.WriteString("...")
	}
	sb.WriteString("]")
	return sb.String()
}
