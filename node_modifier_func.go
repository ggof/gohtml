package gohtml

type NodeModifierFunc func(*TagNode)

func (nmf NodeModifierFunc) Modify(node *TagNode) { nmf(node) }
