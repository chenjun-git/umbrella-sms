package common

import (
	"time"

	"github.com/BurntSushi/toml"
)

type Configs struct {
	Listen    string
	Yunzhixun *YunzhixunConfig
	Monitor   *MonitorConfig
	HTTP      *httpConfig
}

type YunzhixunConfig struct {
	AppId                  string
	Sid                    string
	Token                  string
	Timeout                Duration
	SignupTemplateId       string
	ResetPasswdTemplateId  string
	ChangePasswdTemplateId string
}

type MonitorConfig struct {
	Namespace string
	Subsystem string
}

type httpConfig struct {
	Listen string
}

// Config 全局配置信息
var Config *Configs

// InitConfig 加载配置
func InitConfig(path string) {
	config, err := loadConfig(path)
	if err != nil {
		panic(err)
	}
	Config = config
}

func loadConfig(path string) (*Configs, error) {
	config := new(Configs)
	if _, err := toml.DecodeFile(path, config); err != nil {
		return nil, err
	}
	return config, nil
}

type Duration struct {
	time.Duration
}

func (d *Duration) UnmarshalText(text []byte) (err error) {
	d.Duration, err = time.ParseDuration(string(text))
	return err
}

func (d *Duration) D() time.Duration {
	return d.Duration
}
