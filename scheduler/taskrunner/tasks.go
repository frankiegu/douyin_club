package taskrunner

import (
	"errors"
	"sync"
	"github.com/Yq2/douyin_club/scheduler/dbops"
	"github.com/Yq2/douyin_club/scheduler/ossops"
	Logger "github.com/Yq2/douyin_club/scheduler/logs"
	"time"
)
var log = Logger.Log

func deleteVideo(vid string) error {
	ossfn := "videos/" + vid
	bn := "yq2-videos"
	//在阿里云OSS里面删除指定视频
	//ossfn 对象文件名
	//bn BUCKET存储桶名
	ok := ossops.DeleteObject(ossfn, bn)
	if !ok {
		log.Printf("Deleting video error, oss operation failed")
		return errors.New("Deleting video error")
	}

	return nil
}
//发送删除视频指令
func VideoClearDispatcher(dc dataChan) error {
	//从视频回收站里面随机找3个视频来删除
	res, err := dbops.ReadVideoDeletionRecord(3) //这里查询3条记录
	if err != nil {
		log.Printf("Video clear dispatcher error: %v", err)
		return err
	}

	if len(res) == 0 {
		return errors.New("All tasks finished")
	}
	//把待删除的视频ID发送到dataChan通道里面，可能会阻塞
	for _, id := range res {
		dc <- id //发送视频ID
	}

	return nil
}

//执行清除视频任务
func VideoClearExecutor(dc dataChan) error {
	errMap := &sync.Map{}
	var err error

	forloop:
		for {
			select {
			case vid :=<- dc: //接收视频ID
				//每接收到一个ID开启一个go，所以说多个删除视频任务之间是并发的
				go func(id interface{}) {
					//从OSS里面删除视频，id.(string)对id做类型断言  因为id是interface{}类型所以可以直接做类型断言
					if err := deleteVideo(id.(string)); err != nil {
						//如果删除视频出错，则将ID和ERR存储到map
						errMap.Store(id, err)
						return
					}
					//从视频回收站里面删除视频，id.(string)对id做类型断言  因为id是interface{}类型所以可以直接做类型断言
					if err := dbops.DelVideoDeletionRecord(id.(string)); err != nil {
						//如果删除视频出错，则将ID和ERR存储到map
						errMap.Store(id, err)
						return 
					}
				}(vid)
			default:
				time.Sleep(time.Second * 2)
				break forloop
			}
		}

	errMap.Range(func(k, v interface{}) bool {
		err = v.(error)
		if err != nil {
			return false
		}
		return true
	})

	return err
}
