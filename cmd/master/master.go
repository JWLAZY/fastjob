package main

import (
	"context"
	"fmt"

	"gitlab.xuanke.com/live/fastjob/internal/etcd"
	"gitlab.xuanke.com/live/fastjob/internal/fhttp"
)

func main() {
	// 1. 启动 api服务
	// 2. 注册 master(etcd)
	// 3. 接收任务注册
	// 4. 维护任务队列
	https := fhttp.NewHttpServer(1010)
	etcdClient, err := etcd.NewEtcdClient("xxx")
	if err != nil {
		fmt.Println("error:", err)
	}
	etcdClient.PutWithAlive(context.Background(), "job_master", "xxx", 6)
	etcdClient.PullItemFromQueue("master")
	https.Start()
}
