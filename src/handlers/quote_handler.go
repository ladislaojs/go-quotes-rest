package handlers

import (
	"encoding/json"
	"net/http"
	"quote/src/model"
	"quote/src/service"
	"strconv"
)

type QuoteHandler struct {
	QuoteService *service.QuoteService
}

func NewQuoteHandler(s *service.QuoteService) *QuoteHandler {
	return &QuoteHandler{QuoteService: s}
}

func handleError(w http.ResponseWriter, err error, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if encodeErr := json.NewEncoder(w).Encode(map[string]string{"error": err.Error()}); encodeErr != nil {
		return
	}
}

func (h *QuoteHandler) AddQuote(w http.ResponseWriter, r *http.Request) {
	var quote model.Quote

	if err := json.NewDecoder(r.Body).Decode(&quote); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if quote.Quote == "" {
		http.Error(w, "quote can't be empty", http.StatusBadRequest)
		return
	}

	if quote.Author == "" {
		http.Error(w, "author can't be empty", http.StatusBadRequest)
		return
	}

	quote = h.QuoteService.AddQuote(quote.Author, quote.Quote)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(quote); err != nil {
		return
	}
}

func (h *QuoteHandler) GetQuotes(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")
	quotes := h.QuoteService.GetAllQuotes(author)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(quotes); err != nil {
		return
	}
}

func (h *QuoteHandler) GetRandomQuote(w http.ResponseWriter, r *http.Request) {
	quotes, err := h.QuoteService.GetRandomQuote()

	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(quotes); err != nil {
		return
	}
}

func (h *QuoteHandler) DeleteQuote(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	quoteId, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err = h.QuoteService.DeleteQuote(quoteId); err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
}
