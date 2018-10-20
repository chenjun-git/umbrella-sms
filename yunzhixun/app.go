package yunzhixun

import (
	"time"
)

type YunzhixunApp struct {
	AppId   string
	Sid     string
	Token   string
	Timeout time.Duration
}

func NewYunzhixunApp(appId, sid, token string, timeout time.Duration) *YunzhixunApp {
	return &YunzhixunApp{
		AppId:   appId,
		Sid:     sid,
		Token:   token,
		Timeout: timeout,
	}
}

func (yzx *YunzhixunApp) SetHttpClientTimeout() {
	if yzx.Timeout != 0 {
		SetClientTimeout(yzx.Timeout)
	}
}
