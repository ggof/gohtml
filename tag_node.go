package gohtml

import (
	"html"
	"io"
)

type TagNode struct {
	tag        string
	attributes map[string]string
	children   []NodeRenderer
}

func (tn TagNode) Modify(other *TagNode) {
	other.children = append(other.children, tn)
}

func (tn TagNode) Render(builder io.Writer) {
	builder.Write([]byte("<" + tn.tag))

	for k, v := range tn.attributes {
		builder.Write([]byte(" " + k + "="))
		builder.Write([]byte("\"" + html.EscapeString(v) + "\""))
	}

	builder.Write([]byte{'>'})

	for _, c := range tn.children {
		c.Render(builder)
	}

	builder.Write([]byte("</" + tn.tag + ">"))
}
