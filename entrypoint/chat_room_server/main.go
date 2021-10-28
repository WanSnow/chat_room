package main

import "chatRooms/server/chat_room_server"

func main() {
	crs := chat_room_server.NewChatRoomServer()
	crs.Start()
}
