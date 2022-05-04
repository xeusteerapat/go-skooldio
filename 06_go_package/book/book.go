package book

type book struct {
	Name   string
	Author string
}

func New() book {
	return book{}
}
