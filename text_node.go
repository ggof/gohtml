package gohtml

import (
	"html"
	"io"
)

type TextNode string

func (tn TextNode) Modify(other *TagNode) {
	other.children = append(other.children, tn)
}

func (tn TextNode) Render(w io.Writer) {
	w.Write([]byte(html.EscapeString(string(tn))))
}
