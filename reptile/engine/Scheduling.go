package engine

type Scheduling struct {
	RequestChan chan Request
	WorkChan chan chan Request
}

func (this *Scheduling) SubmitRequest(req Request) {
	this.RequestChan <- req
}

func (this *Scheduling) WorkReady(workChan chan Request) {
	this.WorkChan <- workChan
}

func (this *Scheduling) Controller() {
	this.WorkChan = make(chan chan Request)
	this.RequestChan = make(chan Request)
	go func() {
		requestQueue := []Request{}
		workChanQueue := []chan Request{}
		for {
			var activeReq Request
			var activeWork chan Request
			if len(requestQueue) > 0 && len(workChanQueue) > 0 {
				activeWork = workChanQueue[0]
				activeReq = requestQueue[0]
			}
			select {
			case r := <-this.RequestChan:
				requestQueue = append(requestQueue, r)
			case w := <-this.WorkChan:
				workChanQueue = append(workChanQueue, w)
			case activeWork <- activeReq:
				requestQueue = requestQueue[1:]
				workChanQueue = workChanQueue[1:]
			}
		}
	}()
}