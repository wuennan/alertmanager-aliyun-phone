package aliyun

import (
	"log"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dyvmsapi20170525 "github.com/alibabacloud-go/dyvmsapi-20170525/v4/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
)

type aliYun struct {}


func NewAliCloud() *aliYun {
	return &aliYun{}
}

// CreateClient 创建阿里云Client
func (a *aliYun)CreateClient(accessKeyId,accessKeySecret,endpoint string) (client *dyvmsapi20170525.Client, err error) {
	config := &openapi.Config{
		AccessKeyId:     &accessKeyId,
		AccessKeySecret: &accessKeySecret,
		Endpoint:        &endpoint,
	}
	client, err = dyvmsapi20170525.NewClient(config)
	return
}

// Call 调用阿里云语音服务
func (a *aliYun)Call(client *dyvmsapi20170525.Client,calledNumber,ttsCode,alertMsg string) (err error) {
	// 构建请求
	singleCallByTtsRequest := &dyvmsapi20170525.SingleCallByTtsRequest{
		CalledNumber: &calledNumber,
		TtsCode:      &ttsCode,
		TtsParam:     &alertMsg,
	}
	runtime := &util.RuntimeOptions{}
	callResult, err := client.SingleCallByTtsWithOptions(singleCallByTtsRequest, runtime)

	log.Println(callResult.Body)
	return err
}