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
//
//	// 字符串上传下载
//	_ = bucket.PutObject("my-object-1", strings.NewReader("Hello Oss"))
//	rd, _ := bucket.GetObject("my-object-1")
//	data, _ := ioutil.ReadAll(rd)
//	rd.Close()
//	fmt.Println(string(data))
//
//	// 文件上传下载
//	_ = bucket.PutObjectFromFile("a/b.txt", "/Users/go/Downloads/1.txt")
//	_ = bucket.GetObjectToFile("my-object-2", "mynewpic.jpg")
//
//	// 分片并发，断电续传上传/下载
//	_ = bucket.UploadFile("my-object-3", "mypic.jpg", 100*1024, oss.Routines(3), oss.Checkpoint(true, ""))
//	_ = bucket.DownloadFile("my-object-3", "mynewpic.jpg", 100*1024, oss.Routines(3), oss.Checkpoint(true, ""))
//
//	// 查看object
//	lsRes, _ := bucket.ListObjects()
//	fmt.Println("my objects", lsRes.Objects)
//
//}