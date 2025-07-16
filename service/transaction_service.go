package service

import (
	"test-naga-exchange/model"
	"test-naga-exchange/repository"
	"time"

	"github.com/google/uuid"
)

type TransactionService interface {
	GetUserTransactions(userID uuid.UUID) ([]model.Transaction, error)
	ProcessTransaction(tx *model.Transaction, userID uuid.UUID) error
}

type transactionService struct {
	repo repository.TransactionRepository
}

func NewTransactionService(r repository.TransactionRepository) TransactionService {
	return &transactionService{r}
}

func (s *transactionService) GetUserTransactions(userID uuid.UUID) ([]model.Transaction, error) {
	return s.repo.FindByUserID(userID.String())
}

func (s *transactionService) ProcessTransaction(tx *model.Transaction, userID uuid.UUID) error {
	tx.UserID = userID
	tx.UpdatedAt = time.Now()
	if tx.ID != uuid.Nil {
		return s.repo.Update(tx)
	}
	tx.ID = uuid.New()
	tx.CreatedAt = time.Now()
	return s.repo.Create(tx)
}
