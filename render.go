package gohtml

import "strings"

func Render(r NodeRenderer) string {
	w := strings.Builder{}
	r.Render(&w)
	return w.String()
}

func Home() NodeRenderer {
	return Html(
		Div(ID("hello"), Class("some", "class"),
			C("Hello world!"),
			Ul(
				For([]string{"some", "name"}, func(s string) NodeRenderer { return Li(C(s)) }),
			),
		),
	)
}
