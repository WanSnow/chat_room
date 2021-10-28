package chat_room_server

import (
	"chatRooms/client/heart_beat_client"
	"log"
	"time"

	"google.golang.org/grpc"
)

type ChatRoomServer struct {
	HBClient *heart_beat_client.HeartBeatClient
}

func (crs *ChatRoomServer) Start() {
	go crs.HBClient.Start()
	for {
		time.Sleep(1 * time.Second)
	}
}

func NewChatRoomServer() *ChatRoomServer {
	conn, err := grpc.Dial("127.0.0.1:9696", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	crs := &ChatRoomServer{
		HBClient: heart_beat_client.NewClient(conn),
	}
	return crs
}
