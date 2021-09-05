package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Auth struct {
	Token string
}

type AuthRepo interface {
	LoginByUsername(ctx context.Context, username string, password string) (*Auth, error)
}

type AuthUseCase struct {
	repo AuthRepo
	log  *log.Helper
}

func NewAuthUseCase(repo AuthRepo, logger log.Logger) *AuthUseCase {
	return &AuthUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/auth"))}
}

func (uc *AuthUseCase) GetToken(ctx context.Context, username string, password string) (*Auth, error) {
	out, err := uc.repo.LoginByUsername(ctx, username, password)
	if err != nil {
		return nil, err
	}
	return out, nil
}
