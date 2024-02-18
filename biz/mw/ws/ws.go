package ws

import "qnc/biz/servers"

func Init() {
	go servers.NewClientManager().Start()

	// go servers.WriteMessage()

	go servers.PingTimer()
}
