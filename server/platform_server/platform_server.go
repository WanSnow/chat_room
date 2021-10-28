package platform_server

import (
	"chatRooms/server/heart_beat_server"
	"time"
)

type PlatformServer struct {
	HBServer *heart_beat_server.HeartBeatServer
}

func (ps *PlatformServer) Start() {
	go ps.HBServer.Start()
	for {
		time.Sleep(1 * time.Second)
	}
}

func NewPlatformServer() *PlatformServer {
	return &PlatformServer{
		HBServer: heart_beat_server.NewHeartBeatServer(),
	}
}
