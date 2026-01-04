package domain

type Book struct {
	id         int
	title      string
	year       int
	author     string
	price      int
	stock      int
	categoryID int
}

type NewBookData struct {
	ID         int
	Title      string
	Year       int
	Author     string
	Price      int
	Stock      int
	CategoryID int
}

func NewBook(data NewBookData) (Book, error) {
	return Book{
		id:         data.ID,
		title:      data.Title,
		year:       data.Year,
		author:     data.Author,
		price:      data.Price,
		stock:      data.Stock,
		categoryID: data.CategoryID,
	}, nil
}

func (b Book) ID() int {
	return b.id
}

func (b Book) Title() string {
	return b.title
}

func (b Book) Year() int {
	return b.year
}

func (b Book) Author() string {
	return b.author
}

func (b Book) Price() int {
	return b.price
}

func (b Book) Stock() int {
	return b.stock
}

func (b Book) CategoryID() int {
	return b.categoryID
}
