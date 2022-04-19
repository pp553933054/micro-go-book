package setup

import (
	"github.com/coreos/etcd/clientv3"
	conf "github.com/pp553933054/micro-go-book/ch13-seckill/pkg/config"
	"log"
	"time"
)

//初始化Etcd
func InitEtcd() {

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"}, // conf.Etcd.Host
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Printf("Connect etcd failed. Error : %v", err)
	}
	conf.Etcd.EtcdSecProductKey = "product"
	conf.Etcd.EtcdConn = cli
}
