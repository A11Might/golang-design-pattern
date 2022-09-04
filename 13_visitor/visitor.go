package visitor

import "fmt"

// Visitor 访问者(visitor)
type Visitor interface {
	VisitFile(*File)
	VisitDirectory(*Directory)
}

// ListVisitor 具体的访问者(concrete visitor)
type ListVisitor struct {
	currentDir string
}

func NewListVisitor() *ListVisitor {
	return &ListVisitor{
		currentDir: "",
	}
}

func (l *ListVisitor) VisitFile(file *File) {
	fmt.Println(l.currentDir + "/" + file.ToString())
}

func (l *ListVisitor) VisitDirectory(directory *Directory) {
	fmt.Println(l.currentDir + "/" + directory.ToString())
	saveDir := l.currentDir
	l.currentDir = l.currentDir + "/" + directory.GetName()
	for _, entry := range directory.Iterator() {
		entry.Accept(l)
	}
	l.currentDir = saveDir
}

// Element 元素(element)
type Element interface {
	Accept(Visitor)
	GetName() string
	GetSize() int
}

// File 具体的元素(concrete element)
type File struct {
	name string
	size int
}

func NewFile(name string, size int) *File {
	return &File{
		name: name,
		size: size,
	}
}

func (f *File) Accept(v Visitor) {
	v.VisitFile(f)
}

func (f *File) GetName() string {
	return f.name
}

func (f *File) GetSize() int {
	return f.size
}

func (f *File) ToString() string {
	return fmt.Sprintf("%s (%d)", f.GetName(), f.GetSize())
}

// Directory 具体的元素(concrete element)
type Directory struct {
	name string
	dir  []Element
}

func NewDirectory(name string) *Directory {
	return &Directory{
		name: name,
	}
}

func (d *Directory) Accept(v Visitor) {
	v.VisitDirectory(d)
}

func (d *Directory) GetName() string {
	return d.name
}

func (d *Directory) GetSize() int {
	size := 0
	for _, file := range d.dir {
		size += file.GetSize()
	}
	return size
}

func (d *Directory) Add(element Element) {
	d.dir = append(d.dir, element)
}

func (d *Directory) Iterator() []Element {
	return d.dir
}

func (f *Directory) ToString() string {
	return fmt.Sprintf("%s (%d)", f.GetName(), f.GetSize())
}
