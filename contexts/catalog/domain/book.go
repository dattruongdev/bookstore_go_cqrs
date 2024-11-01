package domain

type Book struct {
	Isbn      Isbn    `json:"isbn"`
	Title     string  `json:"title"`
	Edition   string  `json:"edition"`
	Author    string  `json:"author"`
	Publisher string  `json:"publisher"`
	Source    string  `json:"source"`
	Cost      float64 `json:"cost"`
}

func NewBook(isbn Isbn, title, edition, author, publisher, source string, cost float64) *Book {
	return &Book{
		isbn,
		title,
		edition,
		author,
		publisher,
		source,
		cost,
	}
}
