package config
import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/fsnotify/fsnotify"
)

var Conf = new(AppConfig)

type AppConfig struct {
	AliYunConfig	`mapstructure:"aliyun"`
	Contact map[string]string `yaml:"contact"`
}

type AliYunConfig struct {
	AccessKeyId     string	`mapstructure:"access_key_id"`
	AccessKeySecret string	`mapstructure:"access_key_secret"`
	Endpoint        string 	`mapstructure:"endpoint"`
	TtsCode         string	`mapstructure:"tts_code"`
}


// Init 解析配置文件
func Init(configPath string) (err error) {
	viper.SetConfigFile(configPath) // 指定配置文件路径
	err = viper.ReadInConfig()      // 读取配置信息
	if err != nil {                 // 读取配置信息失败
		fmt.Printf("viper.ReadInConfig() failed, err: %v\n", err)
		return
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		fmt.Printf("ConfigFile Unmarshal failed, err: %v\n", err)
	}
	// 监控配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件发生了变化")

		if err := viper.Unmarshal(&Conf); err != nil {
			fmt.Printf("ConfigFile Unmarshal failed, err: %v\n", err)
		}
	})
	return
}