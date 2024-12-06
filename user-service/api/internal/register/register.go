package register

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"net"
	"strconv"
	"time"
	"user-service/api/internal/config"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type Register struct {
	Client  *clientv3.Client
	LocalIp string
	Port    string
	Key     string
	ID      string
}

func generateId() (int64, error) {
	// 定义 10^16 的最大值，用于生成随机数
	max := new(big.Int).Exp(big.NewInt(10), big.NewInt(16), nil)

	// 使用安全随机数生成器生成一个随机数
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0, err
	}
	return n.Int64(), nil
}

func getLocalIP() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, iface := range interfaces {
		// 跳过没有启用或被禁用的接口
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		// 获取接口的地址
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			// 过滤 IPv4 地址
			if ip != nil && ip.To4() != nil {
				return ip.String(), nil
			}
		}
	}
	return "", fmt.Errorf("no valid IP found")
}

func newEtcdClient(c *config.Config) *clientv3.Client {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   c.EtcdConfig.Hosts,
		DialTimeout: time.Duration(c.EtcdConfig.DialTimeout) * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func NewRegister(c *config.Config) *Register {
	client := newEtcdClient(c)
	port := c.Port
	ip, err := getLocalIP()
	if err != nil {
		log.Fatal(err)
	}
	return &Register{
		Client:  client,
		LocalIp: ip,
		Port:    strconv.Itoa(port),
		Key:     c.EtcdConfig.Key,
		ID:      strconv.FormatFloat(float64(time.Now().UnixNano()), 'f', -1, 64),
	}
}

func (r *Register) CreateKeyWithParent() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// check if parentKey exists
	_, err := r.Client.Put(ctx, r.Key+"/"+r.ID, r.LocalIp+":"+r.Port)
	if err != nil {
		return fmt.Errorf("failed to create child key: %v", err)
	}
	return nil
}

// UnRegister deletes all keys under the key with the prefix of the IP and port of the service.
// It is used to clean up the keys created by the service when it is stopped or deleted.
func (r *Register) UnRegister() error {
	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client := r.Client
	defer client.Close()
	_, err := client.Delete(context, r.Key+"/"+r.ID)
	return err
}
