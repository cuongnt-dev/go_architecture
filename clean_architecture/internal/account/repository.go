package account

import "context"

type RepositoryI interface {
	CreateTransaction(ctx context.Context, ts *Transaction) error
}

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) CreateTransaction(ctx context.Context, ts *Transaction) error {
	return nil
}
