package main

import (
	_ "Golang-Redis-Gin/cmd/docs"
	"Golang-Redis-Gin/internal/bootstrap"
	"Golang-Redis-Gin/internal/models"
)
func main() {
	// setup router
	router := bootstrap.SetupRouter()
	// Start RPC
	go models.StartRpcServer()
	// run server
	bootstrap.StartServer(router)
}





