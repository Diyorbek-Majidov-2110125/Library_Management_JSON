package jsons

import (
	"Library_CRM/storage/repo"
	"fmt"
)

type BookLibrary repo.BookLibrary

func (l *BookLibrary) GetBooks()([]repo.Book){
	return l.Books
}

func (l *BookLibrary) AddBook(book repo.Book) {
	l.Books = append(l.Books, book)
}

func (l *BookLibrary) RemoveBook(book repo.Book) {
	for i, val := range l.Books{
		if book.ID == val.ID {
			l.Books = append(l.Books[:i], l.Books[i+1:]...)
		}
	}
}

func (l *BookLibrary) UpdateBook(book repo.Book) {
	for i, val := range l.Books {
		if book.ID == val.ID {
			l.Books[i] = book
		}
	}
}

func (l *BookLibrary) GetBookByID(book repo.Book) {
	for _, val := range l.Books {
		if book.ID == val.ID {
			fmt.Println(val)
		}
	}
}



