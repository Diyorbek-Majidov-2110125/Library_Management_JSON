package repo

type BookData interface {
	GetBooks() ([]Book)
	AddBook(Book)
	RemoveBook(Book)
	UpdateBook(Book)
	GetBookByID(Book)
	GetBookByCategory(Book)
}

type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Year     int    `json:"year"`
	Status   string `json:"status"`
	Price    int    `json:"price"`
	Period   int    `json:"period"`
	Category string `json:"category"`
	Page     int    `json:"page"`
}

type BookLibrary struct {
	Books []Book
}