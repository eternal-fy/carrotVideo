package bosService

import "github.com/baidubce/bce-sdk-go/services/bos"

const (
	bucketName = "199935"
	objectName = "firstObject"
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
func BosUpload() {
	byteArr := []byte("test put object")
	bosClient := InitClient()

	etag, err := bosClient.PutObjectFromBytes(bucketName, objectName, byteArr, nil)
	if err != nil {
		panic("添加对象失败")
	}
	println(etag)

}
