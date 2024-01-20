package handler

import (
	pbMessage "qeim.com/testv1/pb/generate"
)

var (
	//uid 到 userName 的映射
	uidToUserNameDict map[uint64]string = make(map[uint64]string, 0)

	//userName 到 uid 的映射
	userNameToUidDict map[string]uint64 = make(map[string]uint64, 0)

	//uid 到密码的映射
	uidToPassWordDict map[uint64]string = make(map[uint64]string, 0)
)

func handleRegister(registerRequest *pbMessage.RegisterRequest) *pbMessage.RegisterResponse {
	registerResponse := pbMessage.RegisterResponse{}

	//1、前置检查参数是否合法
	if ok, res := preCheck(registerRequest); ok ==false {
		return res
	}

	//2、用户是否已经注册过
	localUid := userNameToUidDict[string(registerRequest.UserName)]


	return endCheck(&registerResponse)
}



//前置检查（语言、系统级别）：参数是否合法
func preCheck(registerRequest *pbMessage.RegisterRequest) (bool, *pbMessage.RegisterResponse) {
	ok := (registerRequest.Password != "") && (registerRequest.UserName != nil && len(registerRequest.UserName) != 0)
	
	if (registerRequest.UserName == nil || len(registerRequest.UserName) == 0) {
		return ok, &pbMessage.RegisterResponse{
			Ok: false,
			CodeState: 0,
			StringState: "userName can not be emperty!!!",
			Uid: 0,
		}
	}

	if (registerRequest.Password == "") {
		return ok, &pbMessage.RegisterResponse{
			Ok: false,
			CodeState: 0,
			StringState: "password can not be emperty!!!",
			Uid: 0,
		}

	}
	
	return ok, nil
}

// 后置检查主要检查输出参数是否合法，不合法的地方自动补全
func endCheck(response *pbMessage.RegisterResponse) *pbMessage.RegisterResponse {
	if response.StringState == "" {
		response.StringState = "hello world"
	}

	return response
}