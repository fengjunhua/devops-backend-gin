package etcd

import (
	"context"
	"fmt"
	"reflect"
)

func Watch()  {
	watcher := EtcdClient.Watch(context.Background(),"test")
	fmt.Println(watcher)
	for changer := range watcher{
		fmt.Println(changer)
		for ek,ev := range changer.Events{

			fmt.Println(ek)
			fmt.Println("========================")
			fmt.Println(reflect.TypeOf(ev))
			fmt.Println("========================")
			fmt.Println(reflect.ValueOf(ev))
			fmt.Println("========================")
			fmt.Printf("%+v",ev)
			fmt.Printf("Type: %s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			fmt.Println(ev.Kv.Version,ev.Kv.CreateRevision,ev.Kv.Lease)
		}
	}
}




func watch2()  {
	watcher := EtcdClient.Watch(context.Background(),"test")
	fmt.Println(watcher)
	for changer := range watcher{
		fmt.Println(changer)
		for ek,ev := range changer.Events{

			fmt.Println(ek)
			fmt.Println("========================")
			fmt.Println(reflect.TypeOf(ev))
			fmt.Println("========================")
			fmt.Println(reflect.ValueOf(ev))
			fmt.Println("========================")
			fmt.Printf("%+v",ev)
			fmt.Printf("Type: %s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			fmt.Println(ev.Kv.Version,ev.Kv.CreateRevision,ev.Kv.Lease)
		}
	}
}
