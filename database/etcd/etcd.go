package etcd

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func get() {
	config := clientv3.Config{
		Endpoints: []string{"182.92.236.162:2379"},
		DialTimeout: 10 * time.Second,
	}
	client,_ := clientv3.New(config)
	defer client.Close()
	/*
	kv := clientv3.NewKV(client)
	ctx := context.Background()
	kv.Put(ctx,"/services/user","user1")

	 */
	ctx := context.Background()
	put, err := client.Put(ctx, "/services/test", "test")
	if err != nil{
		panic(err)
	}
    fmt.Println(put)
}