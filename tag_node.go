package gohtml

import (
	"html"
	"io"
)

type TagNode struct {
	tag        string
	Attributes map[string]string
	Children   []NodeRenderer
}

func (tn TagNode) Modify(other *TagNode) {
	other.Children = append(other.Children, tn)
}

func (tn TagNode) Render(w io.Writer) {
	w.Write([]byte("<" + tn.tag))

	for k, v := range tn.Attributes {
		w.Write([]byte(" " + k + "="))
		w.Write([]byte("'" + html.EscapeString(v) + "'"))
	}

	w.Write([]byte{'>'})

	for _, c := range tn.Children {
		c.Render(w)
	}

	w.Write([]byte("</" + tn.tag + ">"))
}
