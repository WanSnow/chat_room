package server_map

import (
	"fmt"
	"sync"
	"time"
)

const (
	HeartBeatTimeSecond = 1
	LifeTimeSecond      = 3
)

var HeartBeatServerMap *ChatServerMap

func init() {
	HeartBeatServerMap = &ChatServerMap{
		ChatServerMap: make(map[string]*ServerInfo),
	}
}

type ChatServerMap struct {
	sync.RWMutex
	ChatServerMap map[string]*ServerInfo
}

type ServerInfo struct {
	LastHeartBeatTimestamp time.Time
	RequestCount           int64
	Port                   string
	Status                 bool
}

func (csm *ChatServerMap) Monitor() {
	for {
		csm.Lock()
		for _, serverInfo := range csm.ChatServerMap {
			now := time.Now()
			last := serverInfo.LastHeartBeatTimestamp
			if serverInfo.Status {
				if now.Sub(last) > LifeTimeSecond*time.Second {
					serverInfo.Status = false
				}
			}
		}
		csm.Unlock()
		//printChatServerMap()
		time.Sleep(HeartBeatTimeSecond * time.Second)
	}
}

func printChatServerMap() {
	for {
		fmt.Println("heart beat")
		HeartBeatServerMap.RLock()
		for k, v := range HeartBeatServerMap.ChatServerMap {
			fmt.Println(k, v)
		}
		HeartBeatServerMap.RUnlock()
		time.Sleep(time.Second)
	}
}
