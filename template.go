package gohtml

func BaseLayout(name string, body NodeRenderer) NodeRenderer {
	return Html()(
		Body()(body),
	)
}

func toItem(name string) NodeRenderer {
	return Li()(
		C(name),
	)
}

func IndexPage(names []string) NodeRenderer {
	return BaseLayout("Index",
		Div()(
			Ul(ID("hello"), Class(""))(
				For(names, toItem),
			),
		),
	)
}
