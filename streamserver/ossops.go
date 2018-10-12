package main 

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"Yq2/config"
)

var EP string
var AK string
var SK string
//accesskey
//AK SK相当于使用阿里云的用户名和密码

func init() {
	AK = "LTAI4ijU2FCPBLl4" //accessKeyID
	SK = "xZe8OE1ertC5LTg5KkIbRCT00kwMRt" //accessKeySecret
	EP = config.GetOssAddr()  //oss_addr
}

//上传文件到阿里云存储
func UploadToOss(filename string, path string, bn string) bool {
	client, err := oss.New(EP, AK, SK)
	if err != nil {
		log.Printf("Init oss service error: %s", err)
		return false
	}

	bucket, err := client.Bucket(bn) //bn == bucketName
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


