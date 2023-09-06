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

func Attribute(name, value string) NodeModifierFunc {
	return func(node *TagNode) {
		node.Attributes[name] = value
	}
}

func C(value string) TextNode { return TextNode(value) }

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

func create(tag string, modifiers ...NodeModifier) TagNode {
	node := TagNode{tag: tag, Attributes: make(map[string]string), Children: make([]NodeRenderer, 0)}

	for _, modifier := range modifiers {
		modifier.Modify(&node)
	}

	return node
}

func A(modifiers ...NodeModifier) TagNode          { return create("a", modifiers...) }
func Abbr(modifiers ...NodeModifier) TagNode       { return create("abbr", modifiers...) }
func Acronym(modifiers ...NodeModifier) TagNode    { return create("acronym", modifiers...) }
func Address(modifiers ...NodeModifier) TagNode    { return create("address", modifiers...) }
func Applet(modifiers ...NodeModifier) TagNode     { return create("applet", modifiers...) }
func Area(modifiers ...NodeModifier) TagNode       { return create("area", modifiers...) }
func Article(modifiers ...NodeModifier) TagNode    { return create("article", modifiers...) }
func Aside(modifiers ...NodeModifier) TagNode      { return create("aside", modifiers...) }
func Audio(modifiers ...NodeModifier) TagNode      { return create("audio", modifiers...) }
func B(modifiers ...NodeModifier) TagNode          { return create("b", modifiers...) }
func Base(modifiers ...NodeModifier) TagNode       { return create("base", modifiers...) }
func Basefont(modifiers ...NodeModifier) TagNode   { return create("basefont", modifiers...) }
func Bdi(modifiers ...NodeModifier) TagNode        { return create("bdi", modifiers...) }
func Bdo(modifiers ...NodeModifier) TagNode        { return create("bdo", modifiers...) }
func Bgsound(modifiers ...NodeModifier) TagNode    { return create("bgsound", modifiers...) }
func Big(modifiers ...NodeModifier) TagNode        { return create("big", modifiers...) }
func Blink(modifiers ...NodeModifier) TagNode      { return create("blink", modifiers...) }
func Blockquote(modifiers ...NodeModifier) TagNode { return create("blockquote", modifiers...) }
func Body(modifiers ...NodeModifier) TagNode       { return create("body", modifiers...) }
func Br(modifiers ...NodeModifier) TagNode         { return create("br", modifiers...) }
func Button(modifiers ...NodeModifier) TagNode     { return create("button", modifiers...) }
func Canvas(modifiers ...NodeModifier) TagNode     { return create("canvas", modifiers...) }
func Caption(modifiers ...NodeModifier) TagNode    { return create("caption", modifiers...) }
func Center(modifiers ...NodeModifier) TagNode     { return create("center", modifiers...) }
func Circle(modifiers ...NodeModifier) TagNode     { return create("circle", modifiers...) }
func Cite(modifiers ...NodeModifier) TagNode       { return create("cite", modifiers...) }
func ClipPath(modifiers ...NodeModifier) TagNode   { return create("clipPath", modifiers...) }
func Code(modifiers ...NodeModifier) TagNode       { return create("code", modifiers...) }
func Col(modifiers ...NodeModifier) TagNode        { return create("col", modifiers...) }
func Colgroup(modifiers ...NodeModifier) TagNode   { return create("colgroup", modifiers...) }
func Command(modifiers ...NodeModifier) TagNode    { return create("command", modifiers...) }
func Content(modifiers ...NodeModifier) TagNode    { return create("content", modifiers...) }
func Data(modifiers ...NodeModifier) TagNode       { return create("data", modifiers...) }
func Datalist(modifiers ...NodeModifier) TagNode   { return create("datalist", modifiers...) }
func Dd(modifiers ...NodeModifier) TagNode         { return create("dd", modifiers...) }
func Defs(modifiers ...NodeModifier) TagNode       { return create("defs", modifiers...) }
func Del(modifiers ...NodeModifier) TagNode        { return create("del", modifiers...) }
func Details(modifiers ...NodeModifier) TagNode    { return create("details", modifiers...) }
func Dfn(modifiers ...NodeModifier) TagNode        { return create("dfn", modifiers...) }
func Dialog(modifiers ...NodeModifier) TagNode     { return create("dialog", modifiers...) }
func Dir(modifiers ...NodeModifier) TagNode        { return create("dir", modifiers...) }
func Div(modifiers ...NodeModifier) TagNode        { return create("div", modifiers...) }
func Dl(modifiers ...NodeModifier) TagNode         { return create("dl", modifiers...) }
func Dt(modifiers ...NodeModifier) TagNode         { return create("dt", modifiers...) }
func Element(modifiers ...NodeModifier) TagNode    { return create("element", modifiers...) }
func Ellipse(modifiers ...NodeModifier) TagNode    { return create("ellipse", modifiers...) }
func Em(modifiers ...NodeModifier) TagNode         { return create("em", modifiers...) }
func Embed(modifiers ...NodeModifier) TagNode      { return create("embed", modifiers...) }
func Fieldset(modifiers ...NodeModifier) TagNode   { return create("fieldset", modifiers...) }
func Figcaption(modifiers ...NodeModifier) TagNode { return create("figcaption", modifiers...) }
func Figure(modifiers ...NodeModifier) TagNode     { return create("figure", modifiers...) }
func Font(modifiers ...NodeModifier) TagNode       { return create("font", modifiers...) }
func Footer(modifiers ...NodeModifier) TagNode     { return create("footer", modifiers...) }
func ForeignObject(modifiers ...NodeModifier) TagNode {
	return create("foreignObject", modifiers...)
}
func Form(modifiers ...NodeModifier) TagNode     { return create("form", modifiers...) }
func Frame(modifiers ...NodeModifier) TagNode    { return create("frame", modifiers...) }
func Frameset(modifiers ...NodeModifier) TagNode { return create("frameset", modifiers...) }
func G(modifiers ...NodeModifier) TagNode        { return create("g", modifiers...) }
func H1(modifiers ...NodeModifier) TagNode       { return create("h1", modifiers...) }
func H2(modifiers ...NodeModifier) TagNode       { return create("h2", modifiers...) }
func H3(modifiers ...NodeModifier) TagNode       { return create("h3", modifiers...) }
func H4(modifiers ...NodeModifier) TagNode       { return create("h4", modifiers...) }
func H5(modifiers ...NodeModifier) TagNode       { return create("h5", modifiers...) }
func H6(modifiers ...NodeModifier) TagNode       { return create("h6", modifiers...) }
func Head(modifiers ...NodeModifier) TagNode     { return create("head", modifiers...) }
func Header(modifiers ...NodeModifier) TagNode   { return create("header", modifiers...) }
func Hgroup(modifiers ...NodeModifier) TagNode   { return create("hgroup", modifiers...) }
func Hr(modifiers ...NodeModifier) TagNode       { return create("hr", modifiers...) }
func Html(modifiers ...NodeModifier) TagNode     { return create("html", modifiers...) }
func I(modifiers ...NodeModifier) TagNode        { return create("i", modifiers...) }
func Iframe(modifiers ...NodeModifier) TagNode   { return create("iframe", modifiers...) }
func Image(modifiers ...NodeModifier) TagNode    { return create("image", modifiers...) }
func Img(modifiers ...NodeModifier) TagNode      { return create("img", modifiers...) }
func Input(modifiers ...NodeModifier) TagNode    { return create("input", modifiers...) }
func Ins(modifiers ...NodeModifier) TagNode      { return create("ins", modifiers...) }
func Isindex(modifiers ...NodeModifier) TagNode  { return create("isindex", modifiers...) }
func Kbd(modifiers ...NodeModifier) TagNode      { return create("kbd", modifiers...) }
func Keygen(modifiers ...NodeModifier) TagNode   { return create("keygen", modifiers...) }
func Label(modifiers ...NodeModifier) TagNode    { return create("label", modifiers...) }
func Legend(modifiers ...NodeModifier) TagNode   { return create("legend", modifiers...) }
func Li(modifiers ...NodeModifier) TagNode       { return create("li", modifiers...) }
func Line(modifiers ...NodeModifier) TagNode     { return create("line", modifiers...) }
func LinearGradient(modifiers ...NodeModifier) TagNode {
	return create("linearGradient", modifiers...)
}
func Link(modifiers ...NodeModifier) TagNode      { return create("link", modifiers...) }
func Listing(modifiers ...NodeModifier) TagNode   { return create("listing", modifiers...) }
func Main(modifiers ...NodeModifier) TagNode      { return create("main", modifiers...) }
func Map(modifiers ...NodeModifier) TagNode       { return create("map", modifiers...) }
func Mark(modifiers ...NodeModifier) TagNode      { return create("mark", modifiers...) }
func Marquee(modifiers ...NodeModifier) TagNode   { return create("marquee", modifiers...) }
func Mask(modifiers ...NodeModifier) TagNode      { return create("mask", modifiers...) }
func Math(modifiers ...NodeModifier) TagNode      { return create("math", modifiers...) }
func Menu(modifiers ...NodeModifier) TagNode      { return create("menu", modifiers...) }
func Menuitem(modifiers ...NodeModifier) TagNode  { return create("menuitem", modifiers...) }
func Meta(modifiers ...NodeModifier) TagNode      { return create("meta", modifiers...) }
func Meter(modifiers ...NodeModifier) TagNode     { return create("meter", modifiers...) }
func Multicol(modifiers ...NodeModifier) TagNode  { return create("multicol", modifiers...) }
func Nav(modifiers ...NodeModifier) TagNode       { return create("nav", modifiers...) }
func Nextid(modifiers ...NodeModifier) TagNode    { return create("nextid", modifiers...) }
func Nobr(modifiers ...NodeModifier) TagNode      { return create("nobr", modifiers...) }
func Noembed(modifiers ...NodeModifier) TagNode   { return create("noembed", modifiers...) }
func Noframes(modifiers ...NodeModifier) TagNode  { return create("noframes", modifiers...) }
func Noscript(modifiers ...NodeModifier) TagNode  { return create("noscript", modifiers...) }
func Object(modifiers ...NodeModifier) TagNode    { return create("object", modifiers...) }
func Ol(modifiers ...NodeModifier) TagNode        { return create("ol", modifiers...) }
func Optgroup(modifiers ...NodeModifier) TagNode  { return create("optgroup", modifiers...) }
func Option(modifiers ...NodeModifier) TagNode    { return create("option", modifiers...) }
func Output(modifiers ...NodeModifier) TagNode    { return create("output", modifiers...) }
func P(modifiers ...NodeModifier) TagNode         { return create("p", modifiers...) }
func Param(modifiers ...NodeModifier) TagNode     { return create("param", modifiers...) }
func Path(modifiers ...NodeModifier) TagNode      { return create("path", modifiers...) }
func Pattern(modifiers ...NodeModifier) TagNode   { return create("pattern", modifiers...) }
func Picture(modifiers ...NodeModifier) TagNode   { return create("picture", modifiers...) }
func Plaintext(modifiers ...NodeModifier) TagNode { return create("plaintext", modifiers...) }
func Polygon(modifiers ...NodeModifier) TagNode   { return create("polygon", modifiers...) }
func Polyline(modifiers ...NodeModifier) TagNode  { return create("polyline", modifiers...) }
func Pre(modifiers ...NodeModifier) TagNode       { return create("pre", modifiers...) }
func Progress(modifiers ...NodeModifier) TagNode  { return create("progress", modifiers...) }
func Q(modifiers ...NodeModifier) TagNode         { return create("q", modifiers...) }
func RadialGradient(modifiers ...NodeModifier) TagNode {
	return create("radialGradient", modifiers...)
}
func Rb(modifiers ...NodeModifier) TagNode       { return create("rb", modifiers...) }
func Rbc(modifiers ...NodeModifier) TagNode      { return create("rbc", modifiers...) }
func Rect(modifiers ...NodeModifier) TagNode     { return create("rect", modifiers...) }
func Rp(modifiers ...NodeModifier) TagNode       { return create("rp", modifiers...) }
func Rt(modifiers ...NodeModifier) TagNode       { return create("rt", modifiers...) }
func Rtc(modifiers ...NodeModifier) TagNode      { return create("rtc", modifiers...) }
func Ruby(modifiers ...NodeModifier) TagNode     { return create("ruby", modifiers...) }
func S(modifiers ...NodeModifier) TagNode        { return create("s", modifiers...) }
func Samp(modifiers ...NodeModifier) TagNode     { return create("samp", modifiers...) }
func Script(modifiers ...NodeModifier) TagNode   { return create("script", modifiers...) }
func Section(modifiers ...NodeModifier) TagNode  { return create("section", modifiers...) }
func Select(modifiers ...NodeModifier) TagNode   { return create("select", modifiers...) }
func Shadow(modifiers ...NodeModifier) TagNode   { return create("shadow", modifiers...) }
func Slot(modifiers ...NodeModifier) TagNode     { return create("slot", modifiers...) }
func Small(modifiers ...NodeModifier) TagNode    { return create("small", modifiers...) }
func Source(modifiers ...NodeModifier) TagNode   { return create("source", modifiers...) }
func Spacer(modifiers ...NodeModifier) TagNode   { return create("spacer", modifiers...) }
func Span(modifiers ...NodeModifier) TagNode     { return create("span", modifiers...) }
func Stop(modifiers ...NodeModifier) TagNode     { return create("stop", modifiers...) }
func Strike(modifiers ...NodeModifier) TagNode   { return create("strike", modifiers...) }
func Strong(modifiers ...NodeModifier) TagNode   { return create("strong", modifiers...) }
func Style(modifiers ...NodeModifier) TagNode    { return create("style", modifiers...) }
func Sub(modifiers ...NodeModifier) TagNode      { return create("sub", modifiers...) }
func Summary(modifiers ...NodeModifier) TagNode  { return create("summary", modifiers...) }
func Sup(modifiers ...NodeModifier) TagNode      { return create("sup", modifiers...) }
func Svg(modifiers ...NodeModifier) TagNode      { return create("svg", modifiers...) }
func Table(modifiers ...NodeModifier) TagNode    { return create("table", modifiers...) }
func Tbody(modifiers ...NodeModifier) TagNode    { return create("tbody", modifiers...) }
func Td(modifiers ...NodeModifier) TagNode       { return create("td", modifiers...) }
func Template(modifiers ...NodeModifier) TagNode { return create("template", modifiers...) }
func Text(modifiers ...NodeModifier) TagNode     { return create("text", modifiers...) }
func Textarea(modifiers ...NodeModifier) TagNode { return create("textarea", modifiers...) }
func Tfoot(modifiers ...NodeModifier) TagNode    { return create("tfoot", modifiers...) }
func Th(modifiers ...NodeModifier) TagNode       { return create("th", modifiers...) }
func Thead(modifiers ...NodeModifier) TagNode    { return create("thead", modifiers...) }
func Time(modifiers ...NodeModifier) TagNode     { return create("time", modifiers...) }
func Title(modifiers ...NodeModifier) TagNode    { return create("title", modifiers...) }
func Tr(modifiers ...NodeModifier) TagNode       { return create("tr", modifiers...) }
func Track(modifiers ...NodeModifier) TagNode    { return create("track", modifiers...) }
func Tspan(modifiers ...NodeModifier) TagNode    { return create("tspan", modifiers...) }
func Tt(modifiers ...NodeModifier) TagNode       { return create("tt", modifiers...) }
func U(modifiers ...NodeModifier) TagNode        { return create("u", modifiers...) }
func Ul(modifiers ...NodeModifier) TagNode       { return create("ul", modifiers...) }
func Var(modifiers ...NodeModifier) TagNode      { return create("var", modifiers...) }
func Video(modifiers ...NodeModifier) TagNode    { return create("video", modifiers...) }
func Wbr(modifiers ...NodeModifier) TagNode      { return create("wbr", modifiers...) }
func Xmp(modifiers ...NodeModifier) TagNode      { return create("xmp", modifiers...) }
