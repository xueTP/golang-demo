package handle

import "config"

type ConnLimit struct {
	connNum  int
	backChan chan int
}

func NewConnLimit() ConnLimit {
	return ConnLimit{
		connNum:  config.VideoConf.VideoSaveMaxConn,
		backChan: make(chan int, config.VideoConf.VideoSaveMaxConn),
	}
}

func (this ConnLimit) GetConn() {
	for {
		if len(this.backChan) <= this.connNum {
			this.backChan <- 1
			return
		}
	}
}

func (this ConnLimit) GcConn() {
	<-this.backChan
	return
}
