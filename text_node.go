package gohtml

import (
	"html"
	"strings"
)

type TextNode string

func (tn TextNode) Modify(other *TagNode) {
	other.children = append(other.children, tn)
}

func (tn TextNode) Render(builder *strings.Builder) {
	builder.WriteString(html.EscapeString(string(tn)))
}
