package etcd

import (
	"context"
	"fmt"
)

func Put() {

	/*
		kv := clientv3.NewKV(client)
		ctx := context.Background()
		kv.Put(ctx,"/services/user","user1")

	*/
	ctx := context.Background()
	put, err := EtcdClient.Put(ctx, "/services/test", "test")
	if err != nil{
		panic(err)
	}
	fmt.Println(put)
}