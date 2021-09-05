package data

import (
	"context"
	"fmt"
	etcd "github.com/go-kratos/etcd/registry"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	clientv3 "go.etcd.io/etcd/client/v3"
	grpc2 "google.golang.org/grpc"
	authv1 "scrm/api/auth/v1"
	"scrm/internal/conf"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewDiscovery, NewAuthServiceClient)

// Data .
type Data struct {
	ac authv1.AuthClient
}

// NewData .
func NewData(
	c *conf.Data,
	logger log.Logger,
	ac authv1.AuthClient,
) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	fmt.Println("===", conf)
	client, err := clientv3.New(clientv3.Config{Endpoints: []string{conf.GetEtcd().Address},
		DialTimeout: time.Second, DialOptions: []grpc2.DialOption{grpc2.WithBlock()}})
	if err != nil {
		panic(err)
	}
	r := etcd.New(client)
	return r
}

func NewAuthServiceClient(r registry.Discovery) authv1.AuthClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery://microservices/scrm.user.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			tracing.Client(),
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return authv1.NewAuthClient(conn)
}
