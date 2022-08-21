package abstractfactory

import (
	"fmt"
	"os"
	"strings"
)

type ILink interface {
	MakeHTML() string
}

// Link 抽象产品(abstract product)
type Link struct {
	ILink
}

type ITray interface {
	getTray() []ILink
	setTray([]ILink)
	MakeHTML() string
}

// Tray 抽象产品(abstract product)
type Tray struct {
	ITray
}

func (t *Tray) Add(item ILink) {
	t.setTray(append(t.getTray(), item))
}

type IPage interface {
	getTitle() string
	getContent() []ITray
	setContent([]ITray)
	MakeHTML() string
}

// Page 抽象产品(abstract product)
type Page struct {
	IPage
}

func (p *Page) Add(item ITray) {
	p.setContent(append(p.getContent(), item))
}

func (p *Page) Output() {
	filename := p.getTitle() + ".html"
	file, _ := os.Create(filename)
	file.WriteString(p.MakeHTML())
	_ = file.Close()
	fmt.Println(filename, "编写完成。")
}

// Factory 抽象工厂(abstract factory)
type Factory interface {
	CreateLink(caption, url string) *Link
	CreateTray(caption string) *Tray
	CreatePage(title, author string) *Page
}

func GetFactory(className string) Factory {
	switch className {
	case "ListFactory":
		return NewListFactory()
	default:
		return nil
	}
}

// ListLink 具体产品(concrete product)
type ListLink struct {
	caption string
	url     string
}

func NewListLink(caption, url string) ILink {
	return &ListLink{
		caption: caption,
		url:     url,
	}
}

func (l *ListLink) MakeHTML() string {
	return fmt.Sprintf("    <li><a href=\"%s\">%s</a></li>\n", l.url, l.caption)
}

// ListTray 具体产品(concrete product)
type ListTray struct {
	caption string
	tray    []ILink
}

func NewListTray(caption string) ITray {
	return &ListTray{
		caption: caption,
		tray:    make([]ILink, 0),
	}
}

func (l *ListTray) getTray() []ILink {
	return l.tray
}

func (l *ListTray) setTray(tray []ILink) {
	l.tray = tray
}

func (l *ListTray) MakeHTML() string {
	buffer := new(strings.Builder)
	buffer.WriteString("<li>\n")
	buffer.WriteString(l.caption + "\n")
	buffer.WriteString("<ul>\n")
	for i := range l.tray {
		buffer.WriteString(l.tray[i].MakeHTML())
	}
	buffer.WriteString("</ul>\n")
	buffer.WriteString("</li>\n")
	return buffer.String()
}

// ListPage 具体产品(concrete product)
type ListPage struct {
	title   string
	author  string
	content []ITray
}

func NewListPage(title, author string) IPage {
	return &ListPage{
		title:   title,
		author:  author,
		content: make([]ITray, 0),
	}
}

func (l *ListPage) getTitle() string {
	return l.title
}

func (l *ListPage) getContent() []ITray {
	return l.content
}

func (l *ListPage) setContent(content []ITray) {
	l.content = content
}

func (l *ListPage) MakeHTML() string {
	buffer := new(strings.Builder)
	buffer.WriteString("<html><head><title>" + l.title + "</title></head>\n")
	buffer.WriteString("<body>\n")
	buffer.WriteString("<h1>" + l.title + "</h1>\n")
	buffer.WriteString("<ul>\n")
	for i := range l.content {
		buffer.WriteString(l.content[i].MakeHTML())
	}
	buffer.WriteString("</ul>\n")
	buffer.WriteString("<hr><address>" + l.author + "</address>")
	buffer.WriteString("</body></html>\n")
	return buffer.String()
}

// ListFactory 具体工厂(concrete factory)
type ListFactory struct {
}

func NewListFactory() Factory {
	return &ListFactory{}
}

func (l *ListFactory) CreateLink(caption, url string) *Link {
	return &Link{
		ILink: NewListLink(caption, url),
	}
}

func (l *ListFactory) CreateTray(caption string) *Tray {
	return &Tray{
		ITray: NewListTray(caption),
	}
}

func (l *ListFactory) CreatePage(title, author string) *Page {
	return &Page{
		IPage: NewListPage(title, author),
	}
}
