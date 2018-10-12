package ossops 

import (
	Logger "github.com/Yq2/video_server/scheduler/logs"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"Yq2/config"
)

var EP string
var AK string
var SK string

var log = Logger.Log

func init() {
	AK = "LTAI4ijU2FCPBLl4" //accessKeyID
	SK = "xZe8OE1ertC5LTg5KkIbRCT00kwMRt" //accessKeySecret
	EP = config.GetOssAddr()  //oss_addr
}

//在这个模块中该函数用不到
func UploadToOss(filename string, path string, bn string) bool {
	client, err := oss.New(EP, AK, SK)
	if err != nil {
		log.Printf("Init oss service error: %s", err)
		return false
	}

	bucket, err := client.Bucket(bn)
	if err != nil {
		log.Printf("Getting bucket error: %s", err)
		return false
	}

	err = bucket.UploadFile(filename, path, 500*1024, oss.Routines(3))
	if err != nil {
		log.Printf("Uploading object error: %s", err)
		return false
	}

	return true
}

func DeleteObject(filename string, bn string) bool {
	client, err := oss.New(EP, AK, SK)
	if err != nil {
		log.Printf("Init oss service error: %s", err)
		return false
	}

	bucket, err := client.Bucket(bn)
	if err != nil {
		log.Printf("Getting bueckt error: %s", err)
		return false
	}

	err = bucket.DeleteObject(filename)
	if err != nil {
		log.Printf("Deleting object error: %s", err)
		return false
	}

	return true
}


