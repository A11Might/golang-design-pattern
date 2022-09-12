package facade

import (
	"fmt"
	"os"
)

// PageMaker 外观(facade)

func MakeWelcomePage(mailAddr, filename string) {
	mailProp := GetProperties("maildata")
	username := mailProp[mailAddr]
	ofile, _ := os.Create(filename)
	writer := NewHTMLWriter(ofile)
	writer.Title(fmt.Sprintf("Welcome to %s 's page!", username))
	writer.Paragraph(fmt.Sprintf("欢迎来到%s的主页。", username))
	writer.Paragraph("等着你的邮件哦！")
	writer.Mailto(mailAddr, username)
	writer.Close()
	fmt.Printf("%s is created for %s (%s)", filename, mailAddr, username)
}

// Database 复杂子系统(subsystem)

func GetProperties(dbName string) map[string]string {
	return map[string]string{
		"hyuki@hyuki.com":  "Hiroshi Yuki",
		"hanako@hyuki.com": "Hanako Sato",
		"tomura@hyuki.com": "Tomura",
		"mamoru@hyuki.com": "Mamoru Takahashi",
	}
}

// HTMLWriter 复杂子系统(subsystem)
type HTMLWriter struct {
	writer *os.File
}

func NewHTMLWriter(writer *os.File) *HTMLWriter {
	return &HTMLWriter{
		writer: writer,
	}
}

func (h *HTMLWriter) Title(title string) {
	fmt.Fprint(h.writer, "<html>")
	fmt.Fprint(h.writer, "<head>")
	fmt.Fprintf(h.writer, "<title>%s</title>", title)
	fmt.Fprint(h.writer, "</head>")
	fmt.Fprint(h.writer, "<body>\n")
	fmt.Fprintf(h.writer, "<h1>%s</h1>\n", title)
}

func (h *HTMLWriter) Paragraph(msg string) {
	fmt.Fprintf(h.writer, "<p>%s</p>\n", msg)
}

func (h *HTMLWriter) Link(href, caption string) {
	h.Paragraph(fmt.Sprintf("<a href=\"%s\">%s</a>", href, caption))
}

func (h *HTMLWriter) Mailto(mailAddr, username string) {
	h.Link(fmt.Sprintf("mailto:%s", mailAddr), username)
}

func (h *HTMLWriter) Close() {
	fmt.Fprint(h.writer, "</body>")
	fmt.Fprint(h.writer, "</html>\n")
	_ = h.writer.Close()
}
