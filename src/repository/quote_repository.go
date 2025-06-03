package repository

import (
	"errors"
	"math/rand"
	"quote/src/model"
)

type QuoteRepository struct {
	quotes []model.Quote
}

func NewQuoteRepository() *QuoteRepository {
	return &QuoteRepository{}
}

func (r *QuoteRepository) Add(quote *model.Quote) *model.Quote {
	if len(r.quotes) > 0 {
		quote.ID = r.quotes[len(r.quotes)-1].ID + 1
	} else {
		quote.ID = 1
	}
	r.quotes = append(r.quotes, *quote)
	return quote
}

func (r *QuoteRepository) FindAll() []model.Quote {
	if len(r.quotes) == 0 {
		return []model.Quote{}
	}

	return r.quotes
}

func (r *QuoteRepository) FindRandom() (*model.Quote, error) {
	if len(r.quotes) == 0 {
		return nil, errors.New("no quotes found")
	}

	return &r.quotes[rand.Intn(len(r.quotes))], nil
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

func (r *QuoteRepository) Delete(id int) error {
	if len(r.quotes) == 0 {
		return errors.New("no quotes found")
	}

	for i, quote := range r.quotes {
		if quote.ID == id {
			r.quotes = append(r.quotes[:i], r.quotes[i+1:]...)
			return nil
		}
	}

	return errors.New("quote not found")
}
