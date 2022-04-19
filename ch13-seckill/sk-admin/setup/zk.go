package setup

import (
	"fmt"
	conf "github.com/pp553933054/micro-go-book/ch13-seckill/pkg/config"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

//初始化Etcd
func InitZk() {
	var hosts = []string{"localhost:2181"}
	conn, _, err := zk.Connect(hosts, time.Second*5)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Zookeeper Register Service Success")
	conf.Zk.ZkConn = conn
	conf.Zk.SecProductKey = "/product"
}
