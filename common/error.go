package common

import "strings"

// 错误码
const (
	OK                         = 0
	SmsServiceInternalErr      = 100001
	SmsGatewayRequestIOErr     = 100002
	SmsGatewayJsonUnmarshalErr = 100003
	SmsConvertErr              = 100004
)

// ErrorMap 错误码与错误信息map
var ErrorMap = map[int]map[string]string{
	0: map[string]string{
		"EN-US": "OK",
		"ZH-CN": "操作成功",
	},
	100001: map[string]string{
		"EN-US": "Sms ServiceInternal Err",
		"ZH-CN": "操作失败",
	},
	100002: map[string]string{
		"EN-US": "Sms Gateway Request IO Err",
		"ZH-CN": "网关请求io 失败",
	},
	100003: map[string]string{
		"EN-US": "Sms Gateway Json Unmarshal Err",
		"ZH-CN": "Json 解析失败",
	},
	100004: map[string]string{
		"EN-US": "Sms Convert Err",
		"ZH-CN": "错误转换 失败",
	},
}

// GetMsg 错误码转各个语言的错误信息
func GetMsg(code int, languages []string) string {
	msgMap, ok := ErrorMap[code]
	if !ok {
		return "Unknown error"
	}
	for _, lang := range languages {
		if msg, ok := msgMap[strings.ToUpper(lang)]; ok {
			if msg != "" {
				return msg
			}
		}
	}
	return "Unknown error"
}
