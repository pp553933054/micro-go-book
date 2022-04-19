package main

import (
	"github.com/pp553933054/micro-go-book/ch13-seckill/sk-core/setup"
)

func main() {

	setup.InitZk()
	setup.InitRedis()
	setup.RunService()

}
