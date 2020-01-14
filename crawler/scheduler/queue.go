package scheduler

import "github.com/yihuaiyuan/learngo/crawler/engine"

type QueueScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (s *QueueScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueueScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueueScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueueScheduler) Run() {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)

	go func() {
		//两个队列
		var workerQ []chan engine.Request
		var requestQ []engine.Request

		for {
			var activeWorker chan engine.Request
			var activeRequest engine.Request
			if len(workerQ) > 0 && len(requestQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case activeWorker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
