package server

import (
	"business/sms/common"
	"business/sms/yunzhixun"
)

type Server struct {
	Yunzhixun *yunzhixun.YunzhixunApp
}

func NewServer() *Server {
	return &Server{
	}
}

func (s *Server) PatchYunzhixun(config common.YunzhixunConfig) {
	yunzhixun := yunzhixun.NewYunzhixunApp(
		config.AppId,
		config.Sid,
		config.Token,
		config.Timeout.D(),
	)
	yunzhixun.SetHttpClientTimeout()
	s.Yunzhixun = yunzhixun
}
