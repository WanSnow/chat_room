package heart_beat_server

import (
	"chatRooms/model/containers/server_map"
	"chatRooms/service/heart_beat"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc/peer"

	"google.golang.org/grpc"
)

var (
	heartBeatPort = flag.Int("port", 9696, "The server port")
)

type HeartBeatServer struct {
	heart_beat.HeartBeatServer
}

func (hbs *HeartBeatServer) Start() {
	go server_map.HeartBeatServerMap.Monitor()
	flag.Parsed()
	grpcServer := grpc.NewServer()

	// 控制台输出serverMap
	go printChatServerMap()

	heart_beat.RegisterHeartBeatServer(grpcServer, hbs)

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *heartBeatPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalln(err)
		return
	}
}

func printChatServerMap() {
	for {
		fmt.Println("heart beat")
		server_map.HeartBeatServerMap.RLock()
		for k, v := range server_map.HeartBeatServerMap.ChatServerMap {
			fmt.Println(k, v)
		}
		server_map.HeartBeatServerMap.RUnlock()
		time.Sleep(time.Second)
	}
}

func NewHeartBeatServer() *HeartBeatServer {
	return &HeartBeatServer{}
}

func (hbs *HeartBeatServer) HeartBeat(ctx context.Context, request *heart_beat.HeartBeatMesRequest) (*heart_beat.HeartBeatMesResponse, error) {
	if ctx.Err() == context.Canceled {
		log.Println("SearchService.Search canceled")
		return nil, status.Errorf(codes.Canceled, "SearchService.Search canceled")
	}
	p, ok := peer.FromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("connect error")
	}
	addr := strings.Split(p.Addr.String(), ":")
	server_map.HeartBeatServerMap.Lock()
	info, ok := server_map.HeartBeatServerMap.ChatServerMap[addr[0]]
	if ok && info.Status {
		info.LastHeartBeatTimestamp = time.Now()
	} else {
		server_map.HeartBeatServerMap.ChatServerMap[addr[0]] = &server_map.ServerInfo{
			LastHeartBeatTimestamp: time.Now(),
			RequestCount:           0,
			Status:                 true,
			Port:                   request.Port,
		}
	}
	server_map.HeartBeatServerMap.Unlock()
	return &heart_beat.HeartBeatMesResponse{
		Status: "ok",
	}, nil
}
