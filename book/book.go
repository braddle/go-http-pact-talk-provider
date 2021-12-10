package book

type Book struct {
	ISBN            string   `json:"isbn"`
	Title           string   `json:"title"`
	Author          string   `json:"author"`
	Genre           string   `json:"genre"`
	PublicationData JsonDate `json:"publication_data"`
	InPrint         bool     `json:"in_print"`
	NumberOfPages   int64    `json:"number_of_pages"`
}

type JsonDate struct {
}
