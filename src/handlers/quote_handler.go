package handlers

import (
	"encoding/json"
	"math/big"
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

func (h *QuoteHandler) AddQuote(w http.ResponseWriter, r *http.Request) {
	var quote model.Quote

	if err := json.NewDecoder(r.Body).Decode(&quote); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if quote.Quote == "" || quote.Author == "" {
		http.Error(w, `{"error": "Both 'quote' and 'author' fields are required"}`, http.StatusBadRequest)
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
	quotes := h.QuoteService.GetRandomQuote()
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(quotes); err != nil {
		return
	}
}

func (h *QuoteHandler) DeleteQuote(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	quoteId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	h.QuoteService.DeleteQuote(big.NewInt(quoteId))
}
