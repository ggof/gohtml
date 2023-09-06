package gohtml

import (
	"html"
	"io"
)

type TextNode string

func (tn TextNode) Modify(other *TagNode) {
	other.Children = append(other.Children, tn)
}

func (tn TextNode) Render(w io.Writer) {
	w.Write([]byte(html.EscapeString(string(tn))))
}
