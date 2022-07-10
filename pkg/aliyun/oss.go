package aliyun

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"sync"
)

type OssServer struct {

}

var once sync.Once
var ossServerInstance *OssServer

func NewOssServer() *OssServer {
	once.Do(func() {
		ossServerInstance = &OssServer{}
	})
	return ossServerInstance
}

func (that *OssServer) Init (accessKeyId, accessKeySecret, bucketName, endPoint string) (client *oss.Client, bucket *oss.Bucket, err error) {
	client, err = oss.New(endPoint, accessKeyId, accessKeySecret)
	if err != nil {
		return
	}

	if bucketName != "" {
		bucket, err = client.Bucket(bucketName)
		if err != nil {
			return
		}
	}
	return
}

//func main() {
//	accessKeyId := ""
//	accessKeySecret := ""
//	bucketName := ""
//	endPoint := "https://oss-cn-hangzhou.aliyuncs.com"
//	_, bucket, err := NewOssServer().Init(accessKeyId, accessKeySecret, bucketName, endPoint)
//	if err != nil {
//		fmt.Println(err)
//	}
//	_ = bucket.PutObjectFromFile("a/b.txt", "/Users/go/Downloads/1.txt")
//}