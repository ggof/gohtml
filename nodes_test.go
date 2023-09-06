package gohtml

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModifierWorksAsExpected(t *testing.T) {

	div := Div(ID("name"))

	out := &strings.Builder{}
	div.Render(out)
	assert.Equal(t, "<div id='name'></div>", out.String())
}
