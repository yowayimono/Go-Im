package main

import (
	"im/config"
	routes "im/router"
	"im/server"
	websokcet "im/ws"
)

func init() {
	config.ConfigInit()
	websokcet.Manager = *websokcet.NewClientManager()
	server.InitServer()
}
func main() {
	go websokcet.Manager.Start()
	server.RForm()
	r := routes.InitRoute()
	r.Run(":" + config.ServerPort)
}
