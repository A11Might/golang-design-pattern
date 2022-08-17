package iterator

// Aggregate 集合
type Aggregate interface {
	Iterator() Iterator
}

// Iterator 迭代器
type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type Book struct {
	name string
}

func NewBook(name string) *Book {
	return &Book{name: name}
}

func (b *Book) GetName() string {
	return b.name
}

// BookShelf 具体集合(ConcreteAggregate)
type BookShelf struct {
	books []*Book
	last  int
}

func NewBookShelf(maxSize int) *BookShelf {
	return &BookShelf{
		books: make([]*Book, 0, maxSize),
		last:  0,
	}
}

func (b *BookShelf) GetBookAt(index int) *Book {
	return b.books[index]
}

func (b *BookShelf) AppendBook(book *Book) {
	b.books = append(b.books, book)
	b.last++
}

func (b *BookShelf) GetLength() int {
	return b.last
}

func (b *BookShelf) Iterator() Iterator {
	return NewBookShelfIterator(b)
}

// BookShelfIterator 具体迭代器(concreteIterator)
type BookShelfIterator struct {
	bookShelf *BookShelf
	index     int
}

func NewBookShelfIterator(bookShelf *BookShelf) *BookShelfIterator {
	return &BookShelfIterator{
		bookShelf: bookShelf,
		index:     0,
	}
}

func (b *BookShelfIterator) HasNext() bool {
	if b.index < b.bookShelf.GetLength() {
		return true
	} else {
		return false
	}
}

func (b *BookShelfIterator) Next() interface{} {
	book := b.bookShelf.GetBookAt(b.index)
	b.index++
	return book
}
