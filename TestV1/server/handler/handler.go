package handler

import (
	"github.com/cloudwego/netpoll"
	pbMessage "qeim.com/testv1/pb/generate"
)

func Handle(imCloudPbMesage *pbMessage.IMCloudPbMessage, con netpoll.Connection) *pbMessage.IMCloudPbMessage {
	res := &pbMessage.IMCloudPbMessage{
		CmdType: imCloudPbMesage.CmdType,
		IsRequest: !imCloudPbMesage.IsRequest,
		RequestBody: nil,
		ResponseBody: nil,
	}

	if imCloudPbMesage.CmdType == pbMessage.Cmd_RegisterCmd {
		rep := handleRegister(imCloudPbMesage.GetRequestBody().RegisterRequestBody)
		res.ResponseBody = &pbMessage.IMCloudPbMessageResponseBody{
			RegisterResponseBody: rep,
		}
	}

	if imCloudPbMesage.CmdType == pbMessage.Cmd_LoginCmd {
		rep := handleLogin(imCloudPbMesage.GetRequestBody().LoginRequestBody, con) 
		res.ResponseBody = &pbMessage.IMCloudPbMessageResponseBody{
			LoginResponseBody: rep,
		}
	}

	return res
}