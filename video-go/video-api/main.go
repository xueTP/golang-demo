package main

import (
	"golang-demo/video-go/video-api/config"
	"golang-demo/video-go/video-api/handle"
	"log"
	"net/http"
)

func main() {
	// 注册 handle function
	r := handle.RegisterRouter()
	log.Printf("video-go api server is running in %s", config.VideoConf.ServerAddress)
	http.ListenAndServe(config.VideoConf.ServerAddress, handle.NewMiddleHandle(r))
}
