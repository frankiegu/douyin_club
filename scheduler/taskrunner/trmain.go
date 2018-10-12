package taskrunner

import (
	"time"
)

type Worker struct {
	ticker *time.Ticker
	runner *Runner
}

func NewWorker(interval time.Duration, r *Runner) *Worker {
	return &Worker {
		ticker: time.NewTicker(interval * time.Second),
		runner: r,
	}
}

func (w *Worker) startWorker() {
	//定时器
	for {
		select {
		case <- w.ticker.C:
			//开始通道select
			go w.runner.StartAll()
		}
	}
}

func Start() {
	// Start video file cleaning
	r := NewRunner(10, true, VideoClearDispatcher, VideoClearExecutor)
	//定时器
	w := NewWorker(10, r)
	go w.startWorker()
}


