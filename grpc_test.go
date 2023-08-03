package main

import (
	"comment/infrastructure/util/consul"
	"comment/interfaces/proto"
	"context"
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	setup()

	instances, err := consul.Client.GetInstances("comment", "secure=false")
	if err != nil {
		log.Fatalf("consul.GetServices err: %v", err)
	}
	if len(instances) == 0 {
		log.Fatalf("consul.GetServices")
	}
	instance := instances[rand.Intn(len(instances))]

	// 连接服务器
	conn, err := consul.Client.GetGRPCInstanceConn(instance)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer conn.Close()

	// 建立gRPC连接
	var grpcClient proto.CommentClient
	grpcClient = proto.NewCommentClient(conn)

	insert(grpcClient)
}

func insert(client proto.CommentClient) {
	// 初始化上下文，设置请求超时时间为1秒
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	// 延迟关闭请求会话
	defer cancel()

	// 调用我们的服务
	r, err := client.Insert(ctx, &proto.CommentBaseInfo{Pid: 0, Name: "test", Content: "ttt"})
	if err != nil {
		log.Fatalf("Call ListStr err: %v", err)
	}
	log.Printf("resp: %v", r)
}
