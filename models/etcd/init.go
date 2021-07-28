package etcd

import (
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

/*
使用方法:在main函数中调用该InitEtcdClient()函数,
并且在main函数中调用defer CloseEtcdClient()即可。
*/
var (
	EtcdClient *clientv3.Client
)

func InitEtcdClient()  {
	config := clientv3.Config{
		Endpoints: []string{"182.92.236.162:2379"},
		DialTimeout: 5 * time.Second,
	}
	var err error
	EtcdClient,err = clientv3.New(config)

	if err != nil{
		fmt.Println(err)
	} else{
		fmt.Println("connect etcd success!")
	}

	fmt.Println(EtcdClient)

}

func CloseEtcdClient(){

	EtcdClient.Close()

}


