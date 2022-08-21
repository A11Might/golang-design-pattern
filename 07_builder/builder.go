package builder

import (
	"fmt"
	"os"
	"strings"
)

// Builder 建造者
type Builder interface {
	MakeTitle(title string)
	MakeString(str string)
	MakeItems(items []string)
	Close()
}

// Director 监工
type Director struct {
	Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{
		Builder: builder,
	}
}

func (d *Director) Construct() {
	d.MakeTitle("Greeting")
	d.MakeString("从早上至下午")
	d.MakeItems([]string{
		"早上好。",
		"下午好。",
	})
	d.MakeString("晚上")
	d.MakeItems([]string{
		"晚上好。",
		"晚安。",
		"再见。",
	})
	d.Close()
}

// TextBuilder 具体的建造者(concrete builder)
type TextBuilder struct {
	buffer strings.Builder
}

func NewTextBuilder() *TextBuilder {
	return new(TextBuilder)
}

func (t *TextBuilder) MakeTitle(title string) {
	t.buffer.WriteString("====================\n")
	t.buffer.WriteString(fmt.Sprintf("『%s』\n\n", title))
}

func (t *TextBuilder) MakeString(str string) {
	t.buffer.WriteString(fmt.Sprintf("■%s\n\n", str))
}

func (t *TextBuilder) MakeItems(items []string) {
	for i := range items {
		t.buffer.WriteString(fmt.Sprintf("    .%s\n", items[i]))
	}
	t.buffer.WriteByte('\n')
}

func (t *TextBuilder) Close() {
	t.buffer.WriteString("====================\n")
}

func (t *TextBuilder) GetResult() string {
	return t.buffer.String()
}

// HTMLBuilder 具体的建造者(concrete builder)
type HTMLBuilder struct {
	filename string
	writer   *os.File
}

func NewHTMLBuilder() *HTMLBuilder {
	return new(HTMLBuilder)
}

func (h *HTMLBuilder) MakeTitle(title string) {
	h.filename = title + ".html"
	h.writer, _ = os.Create(h.filename)
	fmt.Fprintf(h.writer, "<html><head><title>%s</html></head></title>", title)
	fmt.Fprintf(h.writer, "<h1>%s</h1>", title)
}

func (h *HTMLBuilder) MakeString(str string) {
	fmt.Fprintf(h.writer, "<p>%s</p>", str)
}

func (h *HTMLBuilder) MakeItems(items []string) {
	fmt.Fprintln(h.writer, "<ul>")
	for i := range items {
		fmt.Fprintf(h.writer, "<li>%s</li>\n", items[i])
	}
	fmt.Fprintln(h.writer, "</ul>")

}

func (h *HTMLBuilder) Close() {
	fmt.Fprintln(h.writer, "</body></html>")
	_ = h.writer.Close()
}

func (h *HTMLBuilder) GetResult() string {
	return h.filename
}
