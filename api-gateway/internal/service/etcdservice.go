package service

import (
	"api-gateway/internal/svc"
	"context"
	"fmt"
	"log"
	"sync"

	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type EtcdService struct {
	l        *svc.ServiceContext
	ctx      context.Context
	Key      string
	Services map[string]string
}

var (
	serviceMap map[string]*EtcdService
	// lock Services before write or read
	mu sync.Mutex
)

func NewEtcdService(c context.Context, l *svc.ServiceContext, key string) *EtcdService {
	if serviceMap[key] != nil {
		return serviceMap[key]
	}
	return &EtcdService{
		l:        l,
		ctx:      c,
		Key:      key,
		Services: make(map[string]string),
	}
}

func (e *EtcdService) GetOneNodeByParent() string {
	mu.Lock()
	defer mu.Unlock()
	if len(e.Services) == 0 {
		log.Fatalf("Discovery service first!")
	}
	// TODO loadbalance
	var endpoint string
	for _, v := range e.Services {
		endpoint = v
	}
	return endpoint
}

func (e *EtcdService) Watch() {
	fmt.Println("start watch service")
	watchChan := e.l.EtcdClient.Watch(e.ctx, e.Key, clientv3.WithPrefix())
	for {
		select {
		case <-e.ctx.Done():
			log.Println("stop watch service")
			return
		case resp, ok := <-watchChan:
			if !ok {
				log.Println("etcd watch closed")
				return
			}
			if resp.Canceled {
				log.Println("etcd watch cancel")
				return
			}
			if resp.Err() != nil {
				log.Println(resp.Err())
			}
			for _, ev := range resp.Events {
				childKey := string(ev.Kv.Key)
				childValue := string(ev.Kv.Value)
				switch ev.Type {
				case mvccpb.PUT:
					log.Printf("etcd watch update service %s", childKey)
					e.updateService(childKey, childValue)
				case mvccpb.DELETE:
					log.Printf("etcd watch delete service %s", childKey)
					e.deleteService(childKey)
				}
			}
		}
	}
}

func (e *EtcdService) updateService(childKey, childValue string) {
	mu.Lock()
	defer mu.Unlock()
	e.Services[childKey] = childValue
}

func (e *EtcdService) deleteService(childKey string) {
	mu.Lock()
	defer mu.Unlock()
	delete(e.Services, childKey)
}

func (e *EtcdService) Close() {
	e.l.EtcdClient.Close()
}

func (e *EtcdService) DisCoveryService() error {
	if serviceMap[e.Key] != nil {
		log.Println("service already exist")
		return nil
	}
	resp, err := e.l.EtcdClient.Get(e.ctx, e.Key, clientv3.WithPrefix())
	if err != nil {
		return err
	}
	for _, kv := range resp.Kvs {
		e.Services[string(kv.Key)] = string(kv.Value)
	}
	go e.Watch()
	return nil
}
