package gohtml

import (
	"io"
	"net/url"
	"strings"
)

type NodeModifier interface {
	Modify(*TagNode)
}

type NodeRenderer interface {
	Render(io.Writer)
}

type BodyFunc = func(...NodeRenderer) NodeRenderer

func create(tag string, modifiers ...NodeModifier) BodyFunc {
	node := TagNode{tag: tag}

	for _, modifier := range modifiers {
		modifier.Modify(&node)
	}

	return func(nr ...NodeRenderer) NodeRenderer {
		node.children = nr

		return node
	}
}

func Attribute(name, value string) NodeModifierFunc {
	return func(other *TagNode) {
		other.attributes[name] = value
	}
}

func For[T any](ts []T, transform func(T) NodeRenderer) NodeRendererFunc {
	return func(builder io.Writer) {
		for _, t := range ts {
			transform(t).Render(builder)
		}
	}
}

func If(condition bool, node NodeRenderer) NodeRendererFunc {
	return func(builder io.Writer) {
		if condition {
			node.Render(builder)
		}
	}
}

func ID(id string) NodeModifierFunc            { return Attribute("id", id) }
func Class(classes ...string) NodeModifierFunc { return Attribute("class", strings.Join(classes, " ")) }
func Name(name string) NodeModifierFunc        { return Attribute("name", name) }
func Value(value string) NodeModifierFunc      { return Attribute("value", value) }
func Href(link string) NodeModifierFunc {
	u, err := url.Parse(link)
	if err != nil {
		panic(err)
	}
	return Attribute("href", u.String())
}
func C(value string) TextNode { return TextNode(value) }

func A(modifiers ...NodeModifier) BodyFunc             { return create("a", modifiers...) }
func Abbr(modifiers ...NodeModifier) BodyFunc          { return create("abbr", modifiers...) }
func Acronym(modifiers ...NodeModifier) BodyFunc       { return create("acronym", modifiers...) }
func Address(modifiers ...NodeModifier) BodyFunc       { return create("address", modifiers...) }
func Applet(modifiers ...NodeModifier) BodyFunc        { return create("applet", modifiers...) }
func Area(modifiers ...NodeModifier) BodyFunc          { return create("area", modifiers...) }
func Article(modifiers ...NodeModifier) BodyFunc       { return create("article", modifiers...) }
func Aside(modifiers ...NodeModifier) BodyFunc         { return create("aside", modifiers...) }
func Audio(modifiers ...NodeModifier) BodyFunc         { return create("audio", modifiers...) }
func B(modifiers ...NodeModifier) BodyFunc             { return create("b", modifiers...) }
func Base(modifiers ...NodeModifier) BodyFunc          { return create("base", modifiers...) }
func Basefont(modifiers ...NodeModifier) BodyFunc      { return create("basefont", modifiers...) }
func Bdi(modifiers ...NodeModifier) BodyFunc           { return create("bdi", modifiers...) }
func Bdo(modifiers ...NodeModifier) BodyFunc           { return create("bdo", modifiers...) }
func Bgsound(modifiers ...NodeModifier) BodyFunc       { return create("bgsound", modifiers...) }
func Big(modifiers ...NodeModifier) BodyFunc           { return create("big", modifiers...) }
func Blink(modifiers ...NodeModifier) BodyFunc         { return create("blink", modifiers...) }
func Blockquote(modifiers ...NodeModifier) BodyFunc    { return create("blockquote", modifiers...) }
func Body(modifiers ...NodeModifier) BodyFunc          { return create("body", modifiers...) }
func Br(modifiers ...NodeModifier) BodyFunc            { return create("br", modifiers...) }
func Button(modifiers ...NodeModifier) BodyFunc        { return create("button", modifiers...) }
func Canvas(modifiers ...NodeModifier) BodyFunc        { return create("canvas", modifiers...) }
func Caption(modifiers ...NodeModifier) BodyFunc       { return create("caption", modifiers...) }
func Center(modifiers ...NodeModifier) BodyFunc        { return create("center", modifiers...) }
func Circle(modifiers ...NodeModifier) BodyFunc        { return create("circle", modifiers...) }
func Cite(modifiers ...NodeModifier) BodyFunc          { return create("cite", modifiers...) }
func ClipPath(modifiers ...NodeModifier) BodyFunc      { return create("clipPath", modifiers...) }
func Code(modifiers ...NodeModifier) BodyFunc          { return create("code", modifiers...) }
func Col(modifiers ...NodeModifier) BodyFunc           { return create("col", modifiers...) }
func Colgroup(modifiers ...NodeModifier) BodyFunc      { return create("colgroup", modifiers...) }
func Command(modifiers ...NodeModifier) BodyFunc       { return create("command", modifiers...) }
func Content(modifiers ...NodeModifier) BodyFunc       { return create("content", modifiers...) }
func Data(modifiers ...NodeModifier) BodyFunc          { return create("data", modifiers...) }
func Datalist(modifiers ...NodeModifier) BodyFunc      { return create("datalist", modifiers...) }
func Dd(modifiers ...NodeModifier) BodyFunc            { return create("dd", modifiers...) }
func Defs(modifiers ...NodeModifier) BodyFunc          { return create("defs", modifiers...) }
func Del(modifiers ...NodeModifier) BodyFunc           { return create("del", modifiers...) }
func Details(modifiers ...NodeModifier) BodyFunc       { return create("details", modifiers...) }
func Dfn(modifiers ...NodeModifier) BodyFunc           { return create("dfn", modifiers...) }
func Dialog(modifiers ...NodeModifier) BodyFunc        { return create("dialog", modifiers...) }
func Dir(modifiers ...NodeModifier) BodyFunc           { return create("dir", modifiers...) }
func Div(modifiers ...NodeModifier) BodyFunc           { return create("div", modifiers...) }
func Dl(modifiers ...NodeModifier) BodyFunc            { return create("dl", modifiers...) }
func Dt(modifiers ...NodeModifier) BodyFunc            { return create("dt", modifiers...) }
func Element(modifiers ...NodeModifier) BodyFunc       { return create("element", modifiers...) }
func Ellipse(modifiers ...NodeModifier) BodyFunc       { return create("ellipse", modifiers...) }
func Em(modifiers ...NodeModifier) BodyFunc            { return create("em", modifiers...) }
func Embed(modifiers ...NodeModifier) BodyFunc         { return create("embed", modifiers...) }
func Fieldset(modifiers ...NodeModifier) BodyFunc      { return create("fieldset", modifiers...) }
func Figcaption(modifiers ...NodeModifier) BodyFunc    { return create("figcaption", modifiers...) }
func Figure(modifiers ...NodeModifier) BodyFunc        { return create("figure", modifiers...) }
func Font(modifiers ...NodeModifier) BodyFunc          { return create("font", modifiers...) }
func Footer(modifiers ...NodeModifier) BodyFunc        { return create("footer", modifiers...) }
func ForeignObject(modifiers ...NodeModifier) BodyFunc { return create("foreignObject", modifiers...) }
func Form(modifiers ...NodeModifier) BodyFunc          { return create("form", modifiers...) }
func Frame(modifiers ...NodeModifier) BodyFunc         { return create("frame", modifiers...) }
func Frameset(modifiers ...NodeModifier) BodyFunc      { return create("frameset", modifiers...) }
func G(modifiers ...NodeModifier) BodyFunc             { return create("g", modifiers...) }
func H1(modifiers ...NodeModifier) BodyFunc            { return create("h1", modifiers...) }
func H2(modifiers ...NodeModifier) BodyFunc            { return create("h2", modifiers...) }
func H3(modifiers ...NodeModifier) BodyFunc            { return create("h3", modifiers...) }
func H4(modifiers ...NodeModifier) BodyFunc            { return create("h4", modifiers...) }
func H5(modifiers ...NodeModifier) BodyFunc            { return create("h5", modifiers...) }
func H6(modifiers ...NodeModifier) BodyFunc            { return create("h6", modifiers...) }
func Head(modifiers ...NodeModifier) BodyFunc          { return create("head", modifiers...) }
func Header(modifiers ...NodeModifier) BodyFunc        { return create("header", modifiers...) }
func Hgroup(modifiers ...NodeModifier) BodyFunc        { return create("hgroup", modifiers...) }
func Hr(modifiers ...NodeModifier) BodyFunc            { return create("hr", modifiers...) }
func Html(modifiers ...NodeModifier) BodyFunc          { return create("html", modifiers...) }
func I(modifiers ...NodeModifier) BodyFunc             { return create("i", modifiers...) }
func Iframe(modifiers ...NodeModifier) BodyFunc        { return create("iframe", modifiers...) }
func Image(modifiers ...NodeModifier) BodyFunc         { return create("image", modifiers...) }
func Img(modifiers ...NodeModifier) BodyFunc           { return create("img", modifiers...) }
func Input(modifiers ...NodeModifier) BodyFunc         { return create("input", modifiers...) }
func Ins(modifiers ...NodeModifier) BodyFunc           { return create("ins", modifiers...) }
func Isindex(modifiers ...NodeModifier) BodyFunc       { return create("isindex", modifiers...) }
func Kbd(modifiers ...NodeModifier) BodyFunc           { return create("kbd", modifiers...) }
func Keygen(modifiers ...NodeModifier) BodyFunc        { return create("keygen", modifiers...) }
func Label(modifiers ...NodeModifier) BodyFunc         { return create("label", modifiers...) }
func Legend(modifiers ...NodeModifier) BodyFunc        { return create("legend", modifiers...) }
func Li(modifiers ...NodeModifier) BodyFunc            { return create("li", modifiers...) }
func Line(modifiers ...NodeModifier) BodyFunc          { return create("line", modifiers...) }
func LinearGradient(modifiers ...NodeModifier) BodyFunc {
	return create("linearGradient", modifiers...)
}
func Link(modifiers ...NodeModifier) BodyFunc      { return create("link", modifiers...) }
func Listing(modifiers ...NodeModifier) BodyFunc   { return create("listing", modifiers...) }
func Main(modifiers ...NodeModifier) BodyFunc      { return create("main", modifiers...) }
func Map(modifiers ...NodeModifier) BodyFunc       { return create("map", modifiers...) }
func Mark(modifiers ...NodeModifier) BodyFunc      { return create("mark", modifiers...) }
func Marquee(modifiers ...NodeModifier) BodyFunc   { return create("marquee", modifiers...) }
func Mask(modifiers ...NodeModifier) BodyFunc      { return create("mask", modifiers...) }
func Math(modifiers ...NodeModifier) BodyFunc      { return create("math", modifiers...) }
func Menu(modifiers ...NodeModifier) BodyFunc      { return create("menu", modifiers...) }
func Menuitem(modifiers ...NodeModifier) BodyFunc  { return create("menuitem", modifiers...) }
func Meta(modifiers ...NodeModifier) BodyFunc      { return create("meta", modifiers...) }
func Meter(modifiers ...NodeModifier) BodyFunc     { return create("meter", modifiers...) }
func Multicol(modifiers ...NodeModifier) BodyFunc  { return create("multicol", modifiers...) }
func Nav(modifiers ...NodeModifier) BodyFunc       { return create("nav", modifiers...) }
func Nextid(modifiers ...NodeModifier) BodyFunc    { return create("nextid", modifiers...) }
func Nobr(modifiers ...NodeModifier) BodyFunc      { return create("nobr", modifiers...) }
func Noembed(modifiers ...NodeModifier) BodyFunc   { return create("noembed", modifiers...) }
func Noframes(modifiers ...NodeModifier) BodyFunc  { return create("noframes", modifiers...) }
func Noscript(modifiers ...NodeModifier) BodyFunc  { return create("noscript", modifiers...) }
func Object(modifiers ...NodeModifier) BodyFunc    { return create("object", modifiers...) }
func Ol(modifiers ...NodeModifier) BodyFunc        { return create("ol", modifiers...) }
func Optgroup(modifiers ...NodeModifier) BodyFunc  { return create("optgroup", modifiers...) }
func Option(modifiers ...NodeModifier) BodyFunc    { return create("option", modifiers...) }
func Output(modifiers ...NodeModifier) BodyFunc    { return create("output", modifiers...) }
func P(modifiers ...NodeModifier) BodyFunc         { return create("p", modifiers...) }
func Param(modifiers ...NodeModifier) BodyFunc     { return create("param", modifiers...) }
func Path(modifiers ...NodeModifier) BodyFunc      { return create("path", modifiers...) }
func Pattern(modifiers ...NodeModifier) BodyFunc   { return create("pattern", modifiers...) }
func Picture(modifiers ...NodeModifier) BodyFunc   { return create("picture", modifiers...) }
func Plaintext(modifiers ...NodeModifier) BodyFunc { return create("plaintext", modifiers...) }
func Polygon(modifiers ...NodeModifier) BodyFunc   { return create("polygon", modifiers...) }
func Polyline(modifiers ...NodeModifier) BodyFunc  { return create("polyline", modifiers...) }
func Pre(modifiers ...NodeModifier) BodyFunc       { return create("pre", modifiers...) }
func Progress(modifiers ...NodeModifier) BodyFunc  { return create("progress", modifiers...) }
func Q(modifiers ...NodeModifier) BodyFunc         { return create("q", modifiers...) }
func RadialGradient(modifiers ...NodeModifier) BodyFunc {
	return create("radialGradient", modifiers...)
}
func Rb(modifiers ...NodeModifier) BodyFunc       { return create("rb", modifiers...) }
func Rbc(modifiers ...NodeModifier) BodyFunc      { return create("rbc", modifiers...) }
func Rect(modifiers ...NodeModifier) BodyFunc     { return create("rect", modifiers...) }
func Rp(modifiers ...NodeModifier) BodyFunc       { return create("rp", modifiers...) }
func Rt(modifiers ...NodeModifier) BodyFunc       { return create("rt", modifiers...) }
func Rtc(modifiers ...NodeModifier) BodyFunc      { return create("rtc", modifiers...) }
func Ruby(modifiers ...NodeModifier) BodyFunc     { return create("ruby", modifiers...) }
func S(modifiers ...NodeModifier) BodyFunc        { return create("s", modifiers...) }
func Samp(modifiers ...NodeModifier) BodyFunc     { return create("samp", modifiers...) }
func Script(modifiers ...NodeModifier) BodyFunc   { return create("script", modifiers...) }
func Section(modifiers ...NodeModifier) BodyFunc  { return create("section", modifiers...) }
func Select(modifiers ...NodeModifier) BodyFunc   { return create("select", modifiers...) }
func Shadow(modifiers ...NodeModifier) BodyFunc   { return create("shadow", modifiers...) }
func Slot(modifiers ...NodeModifier) BodyFunc     { return create("slot", modifiers...) }
func Small(modifiers ...NodeModifier) BodyFunc    { return create("small", modifiers...) }
func Source(modifiers ...NodeModifier) BodyFunc   { return create("source", modifiers...) }
func Spacer(modifiers ...NodeModifier) BodyFunc   { return create("spacer", modifiers...) }
func Span(modifiers ...NodeModifier) BodyFunc     { return create("span", modifiers...) }
func Stop(modifiers ...NodeModifier) BodyFunc     { return create("stop", modifiers...) }
func Strike(modifiers ...NodeModifier) BodyFunc   { return create("strike", modifiers...) }
func Strong(modifiers ...NodeModifier) BodyFunc   { return create("strong", modifiers...) }
func Style(modifiers ...NodeModifier) BodyFunc    { return create("style", modifiers...) }
func Sub(modifiers ...NodeModifier) BodyFunc      { return create("sub", modifiers...) }
func Summary(modifiers ...NodeModifier) BodyFunc  { return create("summary", modifiers...) }
func Sup(modifiers ...NodeModifier) BodyFunc      { return create("sup", modifiers...) }
func Svg(modifiers ...NodeModifier) BodyFunc      { return create("svg", modifiers...) }
func Table(modifiers ...NodeModifier) BodyFunc    { return create("table", modifiers...) }
func Tbody(modifiers ...NodeModifier) BodyFunc    { return create("tbody", modifiers...) }
func Td(modifiers ...NodeModifier) BodyFunc       { return create("td", modifiers...) }
func Template(modifiers ...NodeModifier) BodyFunc { return create("template", modifiers...) }
func Text(modifiers ...NodeModifier) BodyFunc     { return create("text", modifiers...) }
func Textarea(modifiers ...NodeModifier) BodyFunc { return create("textarea", modifiers...) }
func Tfoot(modifiers ...NodeModifier) BodyFunc    { return create("tfoot", modifiers...) }
func Th(modifiers ...NodeModifier) BodyFunc       { return create("th", modifiers...) }
func Thead(modifiers ...NodeModifier) BodyFunc    { return create("thead", modifiers...) }
func Time(modifiers ...NodeModifier) BodyFunc     { return create("time", modifiers...) }
func Title(modifiers ...NodeModifier) BodyFunc    { return create("title", modifiers...) }
func Tr(modifiers ...NodeModifier) BodyFunc       { return create("tr", modifiers...) }
func Track(modifiers ...NodeModifier) BodyFunc    { return create("track", modifiers...) }
func Tspan(modifiers ...NodeModifier) BodyFunc    { return create("tspan", modifiers...) }
func Tt(modifiers ...NodeModifier) BodyFunc       { return create("tt", modifiers...) }
func U(modifiers ...NodeModifier) BodyFunc        { return create("u", modifiers...) }
func Ul(modifiers ...NodeModifier) BodyFunc       { return create("ul", modifiers...) }
func Var(modifiers ...NodeModifier) BodyFunc      { return create("var", modifiers...) }
func Video(modifiers ...NodeModifier) BodyFunc    { return create("video", modifiers...) }
func Wbr(modifiers ...NodeModifier) BodyFunc      { return create("wbr", modifiers...) }
func Xmp(modifiers ...NodeModifier) BodyFunc      { return create("xmp", modifiers...) }
