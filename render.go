package gohtml

import "strings"

func Render(r NodeRenderer) string {
	w := strings.Builder{}
	r.Render(&w)
	return w.String()
}
