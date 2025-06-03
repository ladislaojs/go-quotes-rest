package repository

import (
	"math/big"
	"math/rand"
	"quote/src/model"
)

type QuoteRepository struct {
	quotes []model.Quote
}

func NewQuoteRepository() *QuoteRepository {
	return &QuoteRepository{}
}

func (r *QuoteRepository) Add(quote model.Quote) *model.Quote {
	quote.ID = *big.NewInt(int64(len(r.quotes)) + 1)
	r.quotes = append(r.quotes, quote)
	return &quote
}

func (r *QuoteRepository) FindAll() []model.Quote {
	if len(r.quotes) == 0 {
		return []model.Quote{}
	}

	return r.quotes
}

func (r *QuoteRepository) FindRandom() *model.Quote {
	return &r.quotes[rand.Intn(len(r.quotes))]
}

func (r *QuoteRepository) FindAllByAuthor(author string) []model.Quote {
	var quotes []model.Quote

	for _, quote := range r.quotes {
		if quote.Author == author {
			quotes = append(quotes, quote)
		}
	}

	if len(quotes) == 0 {
		return []model.Quote{}
	}

	return quotes
}

func (r *QuoteRepository) Delete(id *big.Int) {
	for i, quote := range r.quotes {
		if quote.ID.Int64() == id.Int64() {
			r.quotes = append(r.quotes[:i], r.quotes[i+1:]...)
		}
	}
}
