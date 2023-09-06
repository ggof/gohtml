package gohtml

import "io"

type ConditionNode struct {
	condition bool
	inner     NodeRenderer
}

func (cn ConditionNode) Render(w io.Writer) {
	if cn.condition {
		cn.inner.Render(w)
	}
}

func If(condition bool, node NodeRenderer) ConditionNode {
	return ConditionNode{condition, node}
}
