package main

import (
	"fmt"
	"rpc"
	"time"
)

func main() {
	//服务器启动时，初始化一个redis连接池
	rpc.MyUserDao = rpc.NewUserDao("127.0.0.1:6379", 16, 0, 300*time.Second)

	data, err := rpc.MyUserDao.GetData("")
	fmt.Println(data, err)
}

