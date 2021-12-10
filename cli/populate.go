package main

import (
	"book-service/book"
	"bytes"
	"encoding/json"
	"net/http"
)

func main() {
	for _, b := range books {
		c, _ := json.Marshal(b)
		http.Post("http://localhost:8080/book", "application/json", bytes.NewBuffer(c))
	}
}

var books = []book.Book{
	{
		ISBN:            "978-1942788331",
		Title:           "Accelerate: The Science of Lean Software and Devops: Building and Scaling High Performing Technology Organizations",
		Author:          "Nicole Forsgren, Jez Humble & Gene Kim",
		Genre:           "Computing",
		PublicationData: book.JsonDate{}, // 30/04/2018
		InPrint:         true,
		NumberOfPages:   123,
	},
	{
		ISBN:            "978-0321601919",
		Title:           "Continuous Delivery: Reliable Software Releases through Build, Test, and Deployment Automation",
		Author:          "Jez Humble & David Farley",
		Genre:           "Computing",
		PublicationData: book.JsonDate{}, // 05/08/2010
		InPrint:         false,
		NumberOfPages:   456,
	},
	{
		ISBN:            "978-0321146533\n",
		Title:           "Test Driven Development: By Example",
		Author:          "Kent Beck",
		Genre:           "Computing",
		PublicationData: book.JsonDate{}, // 08/10/2002
		InPrint:         false,
		NumberOfPages:   789,
	},
	{
		ISBN:            "978-9123651252",
		Title:           "The Da Vinci Code",
		Author:          "Dan Brown",
		Genre:           "Thriller",
		PublicationData: book.JsonDate{},
		InPrint:         false,
		NumberOfPages:   987,
	},
	{
		ISBN:            "978-0099576877",
		Title:           "Moonraker",
		Author:          "Ian Fleming",
		Genre:           "Thriller",
		PublicationData: book.JsonDate{}, // 06/09/2012
		InPrint:         false,
		NumberOfPages:   654,
	},
}
