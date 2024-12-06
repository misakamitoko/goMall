package service

import (
	"api-gateway/internal/svc"
	"context"
	"fmt"
	"log"
	"sync"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type EtcdService struct {
	l   *svc.ServiceContext
	ctx context.Context
}

var (
	etcdCient *EtcdService
	once      sync.Once
)

func NewEtcdService(c context.Context, l *svc.ServiceContext) *EtcdService {
	once.Do(
		func() {
			etcdCient = &EtcdService{
				l:   l,
				ctx: c,
			}
		},
	)
	return etcdCient
}

func (e *EtcdService) GetByParent(key string) []string {
	resp, err := e.l.EtcdClient.Get(e.ctx, key, clientv3.WithPrefix())
	var endpoints []string
	if err != nil {
		log.Fatalf("get from etcd failed, err:%v\n", err)
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
		endpoints = append(endpoints, string(ev.Value))
	}
	return endpoints
}
