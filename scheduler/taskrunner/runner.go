package taskrunner

import (
	"time"
)

type Runner struct {
	Controller controlChan //控制通道
	Error controlChan //错误通道
	Data dataChan  //真正的数据通道
	dataSize int
	longLived bool  //是否是长期运行的通道，不是则会被回收
	Dispatcher fn //下达命令函数
	Executor fn  //执行任务函数
}

func NewRunner(size int, longlived bool, d fn, e fn) *Runner {
	return &Runner {
		Controller: make(chan string, 1), //使用非阻塞通道
		Error: make(chan string, 1), //使用非阻塞通道
		Data: make(chan interface{}, size),  //数据通道自定义大小
		longLived: longlived,
		dataSize: size,
		Dispatcher: d, //闭包
		Executor: e, //闭包
	}
}

func (r *Runner) startDispatch() {
	defer func() {
		//根据是否需要回收通道资源，选择是否要关闭channel
		if !r.longLived {
			close(r.Controller)
			close(r.Data)
			close(r.Error)
		}
	}()

	//READY_TO_DISPATCH 和 READY_TO_EXECUTE两种指令交替发送
	for {
		//select模拟了操作系统底层epoll
		select {
		case c :=<- r.Controller: //控制器通道
			if c == READY_TO_DISPATCH { //如果收到分发指令
				err := r.Dispatcher(r.Data)
				if err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- READY_TO_EXECUTE
				}
			}

			if c == READY_TO_EXECUTE {
				err := r.Executor(r.Data)
				if err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- READY_TO_DISPATCH
				}
			}
		case e :=<- r.Error:
			if e == CLOSE {
				//结束本次select
				return
			}
		default:
			time.Sleep(time.Second * 2)
		}
	}
}

func (r *Runner) StartAll() {
	//预置一个消息过去，否则任务会卡主
	r.Controller <- READY_TO_DISPATCH
	//开始轮训
	r.startDispatch()
}