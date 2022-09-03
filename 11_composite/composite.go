package composite

import "fmt"

// Entry 组件(component)
type Entry interface {
	GetName() string
	GetSize() int
	printList(string)
}

// File 树叶(leaf)
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

func (f *File) GetName() string {
	return f.name
}

func (f *File) GetSize() int {
	return f.size
}

func (f *File) printList(prefix string) {
	fmt.Printf("%s/%s (%d)\n", prefix, f.name, f.size)
}

// Directory 复合物(composite)
type Directory struct {
	name      string
	directory []Entry
}

func NewDirectory(name string) *Directory {
	return &Directory{
		name:      name,
		directory: make([]Entry, 0),
	}
}

func (d *Directory) GetName() string {
	return d.name
}

func (d *Directory) GetSize() int {
	size := 0
	for _, file := range d.directory {
		size += file.GetSize()
	}
	return size
}

func (d *Directory) Add(entry Entry) {
	d.directory = append(d.directory, entry)
}

func (d *Directory) PrintList() {
	d.printList("")
}

func (d *Directory) printList(prefix string) {
	fmt.Printf("%s/%s (%d)\n", prefix, d.name, d.GetSize())
	for _, file := range d.directory {
		file.printList(prefix + "/" + d.name)
	}
}
