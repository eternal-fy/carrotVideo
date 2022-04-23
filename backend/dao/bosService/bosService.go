package bosService

import (
	"github.com/baidubce/bce-sdk-go/services/bos"
	"github.com/baidubce/bce-sdk-go/services/bos/api"
	"io/ioutil"
)

const (
	bucketName = "199935"
)

var bosClient *bos.Client

func init() {
	bosClient = InitClient()
}
func InitClient() (bosClient *bos.Client) {
	AK := "6a4695427c1c45218515753d08de7ebc"
	SK := "405082d5d35b477c8a5390b31f792d11"
	ENDPOINT := "fwh.bcebos.com"
	clientConfig := bos.BosClientConfiguration{
		Ak:               AK,
		Sk:               SK,
		Endpoint:         ENDPOINT,
		RedirectDisabled: false,
	}

	// 初始化一个BosClient
	bosClient, err := bos.NewClientWithConfig(&clientConfig)
	if err != nil {
		panic("创建bosclient失败")
	}
	return bosClient
}

/*
BosUpload
Params-
 bytes: 所需要上传文件的字节
 userName: 上传文件的用户名
 objectName: 上传的文件名
*/
func BosUpload(bytes []byte, userName, objectName string) {
	byteArr := bytes
	meta := SetMeta(userName)
	etag, err := bosClient.PutObjectFromBytes(bucketName, objectName, byteArr, meta)
	if err != nil {
		panic("添加对象失败")
	}
	println(etag)

}

/*
BosGetUrl
Params-
 objectName: 上传的文件名
*/
func BosGetUrl(objectName string) string {
	expirationInSeconds := -1
	url := bosClient.BasicGeneratePresignedUrl(bucketName, objectName, expirationInSeconds)
	return url
}

/*
BosGetObject
Params-

*/
func BosGetObject(objectName string) *api.GetObjectResult {
	res, err := bosClient.BasicGetObject(bucketName, objectName)
	if err != nil {
		panic(err)
	}
	return res
}

func GetStringByName(objectName string) string {
	object := BosGetObject(objectName)
	stream := object.Body
	defer stream.Close()
	bytes, err := ioutil.ReadAll(stream)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

//设置用户自定义元数据
func SetMeta(userName string) *api.PutObjectArgs {
	args := new(api.PutObjectArgs)
	userMap := make(map[string]string)
	userMap["userName"] = userName
	return args
}
