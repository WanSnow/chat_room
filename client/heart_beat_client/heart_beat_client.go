package heart_beat_client

import (
	"chatRooms/model/containers/server_map"
	"chatRooms/service/heart_beat"
	"context"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

type HeartBeatClient struct {
	heart_beat.HeartBeatClient
}

func (c *HeartBeatClient) Start() {
	for {
		ctx, _ := context.WithTimeout(context.Background(), time.Second)
		// Contact the server and print out its response.
		r, err := c.HeartBeat(ctx, &heart_beat.HeartBeatMesRequest{
			Mes:  "doki",
			Port: "9000",
		})
		if err != nil {
			statusErr, ok := status.FromError(err)
			if ok {
				if statusErr.Code() == codes.DeadlineExceeded {
					log.Fatalln("client.HeartBeat err: deadline")
				}
			}

			log.Fatalf("client.HeartBeat err: %v", err)
		}

		log.Printf("status: %s", r.GetStatus())
		time.Sleep(server_map.HeartBeatTimeSecond * time.Second)
	}

}

func NewClient(conn grpc.ClientConnInterface) *HeartBeatClient {
	c := &HeartBeatClient{
		heart_beat.NewHeartBeatClient(conn),
	}
	return c
}
