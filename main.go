package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tidwall/gjson"
	"github.com/wuennan/alertmanager-aliyun-phone/aliyun"
	"github.com/wuennan/alertmanager-aliyun-phone/config"
)

func Alert(w http.ResponseWriter, r *http.Request){
	var calledNumber string
	// 2. 接受告警
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("read body err, %v\n", err)
		return 
	}

	jsonStr := string(body)
	log.Println(jsonStr)
	summary := gjson.Get(jsonStr, "groupLabels.alertname")

	// 组装告警消息
	alertMsg := fmt.Sprintf("{\"summary\":\"%s\"}", summary)
	

	// 3. 初始化阿里云客户端
	ali := aliyun.NewAliCloud()
	client ,err := ali.CreateClient(config.Conf.AccessKeyId, config.Conf.AccessKeySecret,config.Conf.Endpoint)
	if err!= nil {
        log.Printf("create aliyun client failed, err: %v", err)
        return
    }

	// 获取告警电话
	for c,n := range config.Conf.Contact {
		if r.URL.Path == "/"+c {
			calledNumber = n
		}
	}
	
	// 4. 打电话
	if err := ali.Call(client, calledNumber,config.Conf.TtsCode,alertMsg); err!= nil {
        log.Printf("call aliyun failed, err: %v", err)
        return
    }
    log.Println("告警已转发到阿里云")
}




func main() {
	var configPath string
	// 1. 加载配置
	flag.StringVar(&configPath, "f", "./config.yaml", "配置文件路径")
	flag.Parse()
		
	if err := config.Init(configPath); err != nil {
		log.Printf("init config failed, err: %v", err)
		return
	}

	http.HandleFunc("/", Alert)               // 设置路由处理函数
	err := http.ListenAndServe(":8890", nil) // 监听端口并启动服务器
	if err != nil {
		log.Println("服务器启动失败：", err)
	}
}