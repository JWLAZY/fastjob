package etcd

import (
	"context"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"

	// "go.etcd.io/etcd/clientv3"
	recipe "go.etcd.io/etcd/contrib/recipes"
)

type EtcdClient struct {
	cli     *clientv3.Client
	ctx     context.Context
	Timeout time.Duration
}

func NewEtcdClient(entpoint string) (*EtcdClient, error) {
	var err error
	client := &EtcdClient{}
	client.cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{entpoint},
		DialTimeout: 1,
	})

	if err != nil {
		fmt.Println("etcd error", err)
		return nil, err
	}

	return client, nil
}

func (client *EtcdClient) PutWithAlive(ctx context.Context, key string, value string, interval int64) (clientv3.LeaseID, error) {
	cctx, cancel := context.WithTimeout(ctx, client.Timeout)
	grant, err := client.cli.Grant(cctx, interval)
	cancel()
	fmt.Println(err)
	if err != nil {
		return 0, err
	}
	cctx, cancel = context.WithTimeout(ctx, client.Timeout)
	_, err = client.cli.Put(cctx, key, value, clientv3.WithLease(grant.ID))
	cancel()
	if err != nil {
		return 0, err
	}

	alive, err := client.cli.KeepAlive(ctx, grant.ID)
	if err != nil {
		return 0, err
	}

	go func() {
		for {
			select {
			case _, ok := <-alive:
				if !ok {
					fmt.Println("keep alive not ok", key)
					return
				}
			case <-ctx.Done():
				fmt.Println("ctx over")
				return
			}
		}
	}()
	return grant.ID, nil
}

func (client *EtcdClient) PutItemToQueue(ctx context.Context, queue string, value string) {
	q := recipe.NewQueue(client.cli, queue)
	err := q.Enqueue(value)
	if err != nil {
		fmt.Println(err)
	}
}

func (client *EtcdClient) PullItemFromQueue(queue string) (string, error) {
	q := recipe.NewQueue(client.cli, queue)
	str, err := q.Dequeue()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return str, nil
}
