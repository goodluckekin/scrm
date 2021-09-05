package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"scrm/app/auth/service/internal/biz"
)

var _ biz.AuthRepo = (*authRepo)(nil)

type authRepo struct {
	data *Data
	log  *log.Helper
}

func NewAuthRepo(data *Data, logger log.Logger) biz.AuthRepo {
	return &authRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/server-service")),
	}
}

func (r *authRepo) LoginByUsername(ctx context.Context, username string, password string) (*biz.Auth, error) {
	return &biz.Auth{Token: username + password}, nil
}
