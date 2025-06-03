package main

import (
	"net/http"
	"quote/src/handlers"
	"quote/src/repository"
	"quote/src/service"
)

func main() {

	r := http.NewServeMux()

	quoteRepository := repository.NewQuoteRepository()
	quoteService := service.NewQuoteService(quoteRepository)
	quoteHandler := handlers.NewQuoteHandler(quoteService)

	r.HandleFunc("GET /quotes", quoteHandler.GetQuotes)

	r.HandleFunc("POST /quotes", quoteHandler.AddQuote)

	r.HandleFunc("GET /quotes/random", quoteHandler.GetRandomQuote)

	r.HandleFunc("DELETE /quotes/{id}", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(r.PathValue("id")))
		if err != nil {
			return
		}
	})

	err := http.ListenAndServe(":80", r)
	if err != nil {
		return
	}

}
