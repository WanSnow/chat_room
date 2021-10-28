package main

import (
	"chatRooms/server/platform_server"
	"log"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalln(r)
		}
	}()

	ps := platform_server.NewPlatformServer()
	ps.Start()
}
