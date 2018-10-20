package yunzhixun

import (
	"encoding/json"
)

type YunzhixunReq struct {
	Sid   string
	Token string
	AppId string
}

func (r *YunzhixunReq) setAppInfo(sid, token, appId string) {
	r.Sid = sid
	r.Token = token
	r.AppId = appId
}

type YunzhixunResp struct {
	Code string
	Msg  string
}

type SmsSendSingleReq struct {
	YunzhixunReq
	TemplateId string
	Param      string
	Mobile     string
	Uid        string
}

type SmsSendSingleResp struct {
	YunzhixunResp
	Count      string
	CreateDate string
	SmsId      string
	Mobile     string
	Uid        string
}

type SmsSendGroupReq struct {
	YunzhixunReq
	TemplateId string
	Param      string
	Mobile     string
	Uid        string
}

type SmsSendGroupItemResp struct {
	Code   string
	Msg    string
	SmsId  string
	Mobile string
	Uid    string
	Count  string
}

type SmsSendGroupResp struct {
	YunzhixunResp
	CountSum   string
	CreateDate string
	Report     []SmsSendGroupItemResp
}

type SmsSendStatusResp struct {
	Code       string
	Msg        string
	SmsId      string
	Mobile     string
	ReportDate string
	Uid        string
}

func (yzx *YunzhixunApp) SendSingleVerifyCode(req SmsSendSingleReq) *SmsSendSingleResp {
	url := apiURL(SMS_SEND_SINGLE)

	req.setAppInfo(yzx.Sid, yzx.Token, yzx.AppId)

	resp, err := httpReqWithParams("POST", url, req)
	if err != nil {
		return &SmsSendSingleResp{
			YunzhixunResp: YunzhixunResp{
				Code: string(YUNZHIXUN_RESULT_CODE_UNSUPPORTED_RESP),
				Msg:  err.Error(),
			},
		}
	}

	var result SmsSendSingleResp
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return &SmsSendSingleResp{
			YunzhixunResp: YunzhixunResp{
				Code: string(YUNZHIXUN_RESULT_CODE_UNSUPPORTED_RESP),
				Msg:  err.Error(),
			},
		}
	}

	return &result
}
