package gohtml

type NodeModifierFunc func(*TagNode)

func (nmf NodeModifierFunc) Modify(other *TagNode) { nmf(other) }
