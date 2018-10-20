package v1_0

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	commonErrors "github.com/chenjun-git/umbrella-common/errors"

	"github.com/chenjun-git/umbrella-sms/common"
	"github.com/chenjun-git/umbrella-sms/pb"
	"github.com/chenjun-git/umbrella-sms/server"
	"github.com/chenjun-git/umbrella-sms/utils"
)

func sendSignupVerifyCode(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.JSON(w, r, commonErrors.NewError(common.SmsGatewayRequestIOErr, err.Error()))
		return
	}
	var req pb.SmsSendSingleReq
	if err = json.Unmarshal(body, &req); err != nil {
		utils.JSON(w, r, commonErrors.NewError(common.SmsGatewayJsonUnmarshalErr, err.Error()))
		return
	}

	s := server.NewServer()
	s.PatchYunzhixun(*common.Config.Yunzhixun)

	resp, err := s.SendSignupVerifyCode(context.Background(), &req)
	if err != nil {
		utils.JSON(w, r, common.ConvertError(err))
		return
	}

	utils.JSON(w, r, wrapCode(resp))
}
