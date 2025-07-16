package repository

import (
	"test-naga-exchange/model"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindByUserID(userID string) ([]model.Transaction, error)
	Create(transaction *model.Transaction) error
	Update(transaction *model.Transaction) error
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) FindByUserID(userID string) ([]model.Transaction, error) {
	var transactions []model.Transaction
	err := r.db.Where("user_id = ?", userID).Find(&transactions).Error
	return transactions, err
}

func (r *transactionRepository) Create(tx *model.Transaction) error {
	return r.db.Create(tx).Error
}

func (r *transactionRepository) Update(tx *model.Transaction) error {
	return r.db.Save(tx).Error
}
