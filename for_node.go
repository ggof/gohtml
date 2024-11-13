package gohtml

import "io"

type ForNode[T any] struct {
	elements []T
	mapper   func(T) NodeRenderer
}

func (fn ForNode[T]) Render(w io.Writer) {
	for _, e := range fn.elements {
		fn.mapper(e).Render(w)
	}
}

func (fn ForNode[T]) Modify(node *TagNode) {
	for _, e := range fn.elements {
		node.Children = append(node.Children, fn.mapper(e))
	}
}

func For[T any](ts []T, transform func(T) NodeRenderer) ForNode[T] {
	return ForNode[T]{ts, transform}
}
