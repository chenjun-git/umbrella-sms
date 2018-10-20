package server

import (
	"context"
	"fmt"
	"strconv"

	"business/sms/common"
	"business/sms/pb"
	"business/sms/yunzhixun"
)

func (s *Server) SendSignupVerifyCode(ctx context.Context, req *pb.SmsSendSingleReq) (*pb.SmsSendSingleResp, error) {

	sendReq := yunzhixun.SmsSendSingleReq{
		TemplateId: common.Config.Yunzhixun.SignupTemplateId,
		Param:      req.VerifyCode,
		Mobile:     req.Mobile,
		Uid:        "0",
	}
	resp := s.Yunzhixun.SendSingleVerifyCode(sendReq)
	code, err := strconv.Atoi(resp.Code)
	if err != nil{
		return nil, common.NewRPCError(common.SmsServiceInternalErr, fmt.Sprintf("code: %v, msg: %v", resp.Code, resp.Msg))
	}

	if code != yunzhixun.YUNZHIXUN_RESULT_CODE_OK {
		return nil, common.NewRPCError(common.SmsServiceInternalErr, fmt.Sprintf("code: %v, msg: %v", resp.Code, resp.Msg))
	}

	return nil, nil
}
