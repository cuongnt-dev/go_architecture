package account

import "context"

type UseCaseI interface {
	CreateTransaction(ctx context.Context, ts Transaction) error
}

type UserCase struct {
	Repository RepositoryI
}

func NewUserCase(r *Repository) *UserCase {
	return &UserCase{
		Repository: r,
	}
}

func (u *UserCase) CreateTransaction(ctx context.Context, ts Transaction) error {
	return nil
}
