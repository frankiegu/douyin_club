package main 

//流控
type ConnLimiter struct {
	concurrentConn int //当前连接
	bucket chan struct{} //bucket理解为票池
}

func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter {
		concurrentConn: cc,
		bucket: make(chan struct{}, cc),
	}
}

//业界通用的流控解决方案
//使用channel来模拟流控，bucket是一个带缓存的通道
//得到一张票，如果没有空位就阻塞
func (cl *ConnLimiter) GetConn() bool {
	if len(cl.bucket) >= cl.concurrentConn {
		log.Printf("Reached the rate limitation.")
		return false
	}
	//struct{}{}空结构体在内存中只会存在一份，所以用来做channel通信是最佳选择
	cl.bucket <- struct{}{}
	return true
}

//归还一张票
func (cl *ConnLimiter) ReleaseConn() {
	c :=<- cl.bucket
	log.Printf("New connction coming: %d", c)
}
