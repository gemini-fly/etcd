package utils

import (
	"context"
	etcd "go.etcd.io/etcd/client/v3"
	"log"
	"os"
	"strings"
	"time"
)

func Endpoints() (endpoints []string) {
	v, ok := os.LookupEnv("ETCD_ADDR")
	if ok != true {
		log.Fatalln(ok)
	}
	return strings.Split(strings.TrimSpace(v), ";")
}

// NewEtcdClient 创建并返回一个新的 EtcdClient 实例
func NewEtcdClient() (*etcd.Client, error) {
	cli, err := etcd.New(etcd.Config{
		Endpoints:   Endpoints(),
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return nil, err
	}

	return cli, nil
}

func Client() (*etcd.Client, error) {
	clt, err := NewEtcdClient()
	if err != nil {
		panic(err)
	}
	return clt, nil
}

func Get(key string) (value string, err error) {

	clt, err := Client()
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	resp, err := clt.KV.Get(ctx, key)
	if err != nil {
		return "", err
	}
	if len(resp.Kvs) == 0 {
		return "", err
	}
	return string(resp.Kvs[0].Value), nil
}
