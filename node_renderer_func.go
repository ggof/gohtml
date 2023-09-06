package gohtml

import (
	"io"
)

type NodeRendererFunc func(io.Writer)

func (nrf NodeRendererFunc) Render(w io.Writer) { nrf(w) }
