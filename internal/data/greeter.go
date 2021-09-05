package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	v1 "scrm/api/auth/v1"
	"scrm/internal/biz"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *greeterRepo) CreateGreeter(ctx context.Context, g *biz.Greeter) error {
	msg, err := r.data.ac.GetToken(ctx, &v1.LoginRequest{
		Username: "ekin",
		Password: "test",
	})
	fmt.Println("++++++", msg)
	return err
}

func (r *greeterRepo) UpdateGreeter(ctx context.Context, g *biz.Greeter) error {
	return nil
}
