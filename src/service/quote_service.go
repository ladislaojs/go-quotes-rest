package service

import (
	"math/big"
	"quote/src/model"
	"quote/src/repository"
)

type QuoteService struct {
	QuoteRepository *repository.QuoteRepository
}

func NewQuoteService(r *repository.QuoteRepository) *QuoteService {
	return &QuoteService{QuoteRepository: r}
}

func (s *QuoteService) AddQuote(author string, quote string) model.Quote {
	quoteModel := model.Quote{
		Author: author, Quote: quote,
	}

	return *s.QuoteRepository.Add(quoteModel)
}

func (s *QuoteService) GetAllQuotes(author string) []model.Quote {
	if len(author) > 0 {
		return s.QuoteRepository.FindAllByAuthor(author)
	}

	return s.QuoteRepository.FindAll()
}

func (s *QuoteService) GetRandomQuote() *model.Quote {
	return s.QuoteRepository.FindRandom()
}

func (s *QuoteService) DeleteQuote(id *big.Int) {
	s.QuoteRepository.Delete(id)
}
