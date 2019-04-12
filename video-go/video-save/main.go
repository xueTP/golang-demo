package main

import (
	"config"
	"golang-demo/video-go/video-save/handle"
	"log"
	"net/http"
)

func main() {
	r := handle.RegisterRouter()
	log.Printf("video-go save to upload or download is runing in %v", config.VideoConf.SaveServerAddress)
	err := http.ListenAndServe(config.VideoConf.SaveServerAddress, handle.NewMiddleHandle(r))
	log.Println(err)
}
