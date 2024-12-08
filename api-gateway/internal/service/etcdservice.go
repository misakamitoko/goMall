package service

import (
	"api-gateway/internal/svc"
	"context"
	"log"
	"sync"

	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
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
	mu     sync.Mutex
	logger *zap.Logger
)

func InitLogger() {
	logger, _ = zap.NewProduction()
}

func NewEtcdService(c context.Context, l *svc.ServiceContext, key string) *EtcdService {
	if serviceMap[key] != nil {
		return serviceMap[key]
	}
	eService := &EtcdService{
		l:        l,
		ctx:      c,
		Key:      key,
		Services: make(map[string]string),
	}
	eService.DisCoveryService()
	serviceMap[key] = eService
	return eService
}

func GetOneNodeByParent(serviceName string) string {
	if serviceMap[serviceName] == nil {
		logger.Error(
			"Discovrey service first",
		)
		return ""
	}
	// TODO loadbalance
	e := serviceMap[serviceName]
	if e.Services == nil {
		logger.Error(
			"no service available",
			zap.String("serviceName", serviceName),
		)
	}
	var endpoint string
	for _, v := range e.Services {
		endpoint = v
	}
	return endpoint
}

func (e *EtcdService) Watch() {
	logger.Info(
		"start watch service at",
		zap.Int("port", e.l.Config.Port),
	)
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
					logger.Info(
						"etcd watch update service",
					)
					e.updateService(childKey, childValue)
				case mvccpb.DELETE:
					logger.Info(
						"etcd watch delete service",
					)
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
	InitLogger()
	if serviceMap[e.Key] != nil {
		logger.Info(
			"service start at",
			zap.Int("port", e.l.Config.Port),
		)
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
	if serviceMap == nil {
		serviceMap = make(map[string]*EtcdService)
	}
	serviceMap[e.Key] = e
	return nil
}
