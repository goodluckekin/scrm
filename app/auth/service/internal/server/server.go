package server

import (
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"scrm/app/auth/service/internal/conf"
	"time"

	consul "github.com/go-kratos/consul/registry"
	consulAPI "github.com/hashicorp/consul/api"

	etcd "github.com/go-kratos/etcd/registry"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewEtcdRegistrar)

//注册consul
func NewConsulRegistrar(conf *conf.Registry) registry.Registrar {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

//注册etcd
func NewEtcdRegistrar(conf *conf.Registry) registry.Registrar {
	client, err := clientv3.New(clientv3.Config{Endpoints: []string{conf.GetEtcd().Address},
		DialTimeout: time.Second, DialOptions: []grpc.DialOption{grpc.WithBlock()}})
	if err != nil {
		panic(err)
	}
	r := etcd.New(client)
	return r
}
