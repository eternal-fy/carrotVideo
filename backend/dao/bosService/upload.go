package bosService

import (
	"github.com/baidubce/bce-sdk-go/services/bos"
	"github.com/baidubce/bce-sdk-go/services/bos/api"
)

const (
	bucketName = "199935"
)

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
	bosClient := InitClient()
	meta := SetMeta(userName)

	etag, err := bosClient.PutObjectFromBytes(bucketName, objectName, byteArr, meta)
	if err != nil {
		panic("添加对象失败")
	}
	println(etag)

}

//设置用户自定义元数据
func SetMeta(userName string) *api.PutObjectArgs {
	args := new(api.PutObjectArgs)

	// 设置上传内容的长度
	args.ContentLength = 8 * 1024 * 1024 * 100
	userMap := make(map[string]string)
	userMap["userName"] = userName
	return args
}
