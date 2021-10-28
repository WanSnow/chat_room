package main

import (
	"chatRooms/service/heart_beat"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9696", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	for {
		c := heart_beat.NewHeartBeatClient(conn)
		// Contact the server and print out its response.
		ctx, _ := context.WithTimeout(context.Background(), time.Second)
		r, err := c.HeartBeat(ctx, &heart_beat.HeartBeatMesRequest{
			Mes: "doki",
		})
		if err != nil {
			log.Fatalf("could not send heart beat: %v", err)
		}
		log.Printf("status: %s", r.GetStatus())
		time.Sleep(3 * time.Second)
	}
}
