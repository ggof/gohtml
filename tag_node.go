package gohtml

import (
	"html"
	"strings"
)

type TagNode struct {
	tag        string
	attributes map[string]string
	children   []NodeRenderer
}

func (tn TagNode) Modify(other *TagNode) {
	other.children = append(other.children, tn)
}

func (tn TagNode) Render(builder *strings.Builder) {
	builder.WriteRune('<')
	builder.WriteString(tn.tag)

	for k, v := range tn.attributes {
		builder.WriteRune(' ')
		builder.WriteString(k)
		builder.WriteRune('=')
		builder.WriteString("\"" + html.EscapeString(v) + "\"")
	}

	builder.WriteRune('>')

	for _, c := range tn.children {
		c.Render(builder)
	}

	builder.WriteString("</" + tn.tag + ">")
}
