package main

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	conf "github.com/pp553933054/micro-go-book/ch13-seckill/pkg/config"
	"log"
	"testing"
	"time"
)

func TestMain(m *testing.M) {

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"}, // conf.Etcd.Host
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Printf("Connect etcd failed. Error : %v", err)
	}
	conf.Etcd.EtcdSecProductKey = "product"
	conf.Etcd.EtcdConn = cli

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	rsp, err := conf.Etcd.EtcdConn.Get(ctx, "product")
	log.Printf("get from etcd %v", rsp)
}
