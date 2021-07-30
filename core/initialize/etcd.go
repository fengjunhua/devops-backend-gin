package initialize

import (
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

var EtcdClient *clientv3.Client

func InitEtcdClient(){


	config := clientv3.Config{
		//Endpoints: []string{"192.168.0.18:2379","192.168.0.83:2379","192.168.0.105:2379"},
		Endpoints: []string{"192.168.0.87:2379","192.168.0.172:2379","192.168.0.34:2379"},
		DialTimeout: 5 * time.Second,
	}
	var err error
	EtcdClient, err= clientv3.New(config)
	if err != nil{
		log.Println(err)
	} else{
		log.Println("connect etcd success!")
	}

	//log.Println(EtcdClient)
}

func CloseEtcdClient(){
    EtcdClient.Close()
}