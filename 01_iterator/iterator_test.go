package iterator

import "fmt"

func ExampleIterator() {
	bookShelf := NewBookShelf(4)
	bookShelf.AppendBook(NewBook("Around the World in 80 Days"))
	bookShelf.AppendBook(NewBook("Bible"))
	bookShelf.AppendBook(NewBook("Cinderella"))
	bookShelf.AppendBook(NewBook("Daddy-Long-Legs"))
	var it Iterator
	it = bookShelf.Iterator()
	for it.HasNext() {
		book := it.Next().(*Book)
		fmt.Println(book.GetName())
	}

	// Output:
	// Around the World in 80 Days
	// Bible
	// Cinderella
	// Daddy-Long-Legs
}
