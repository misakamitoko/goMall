package svc

import (
	"api-gateway/internal/config"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type ServiceContext struct {
	Config     config.Config
	EtcdClient *clientv3.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	EtcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   c.EtcdConfig.Hosts,
		DialTimeout: time.Duration(c.EtcdConfig.DialTimeout) * time.Second,
	})
	if err != nil {
		log.Fatalf("connect etcd failed, err:%v\n", err)
	}
	return &ServiceContext{
		Config:     c,
		EtcdClient: EtcdClient,
	}
}
