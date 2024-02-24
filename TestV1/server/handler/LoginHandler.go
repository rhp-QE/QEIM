package handler

import (
	"github.com/cloudwego/netpoll"
	pbMessage "qeim.com/testv1/protobuf/generate"
	"qeim.com/testv1/storage/local"
)

func handleLogin(req *pbMessage.LoginRequest, con netpoll.Connection) *pbMessage.LoginResponse {
	// 前置检查
	if check, str := preCheckLoign(req); check == false {
		return &pbMessage.LoginResponse{
			Ok:          false,
			CodeState:   0,
			StringState: str,
		}
	}

	// 是否注册
	uid := userNameToUidDict[req.UserName]
	if uid == 0 {
		return &pbMessage.LoginResponse{
			Ok:          false,
			CodeState:   0,
			StringState: "[Login Error]:not regist",
		}
	}

	// 是否已经登录
	_, logined := uidToConnection[uid]
	if logined {
		return &pbMessage.LoginResponse{
			Ok:          false,
			CodeState:   0,
			StringState: "[Login Error]:logined in other place",
		}
	}

	_ = loginUser(uid, con)

	return &pbMessage.LoginResponse{
		Ok:          true,
		CodeState:   1,
		StringState: "[Login] : Login success",
	}
}

func preCheckLoign(req *pbMessage.LoginRequest) (bool, string) {
	if req == nil {
		return false, "[Login Error]:LoginRequest is empty"
	}

	if req.UserName == "" {
		return false, "[Login Error]:UserName is Emperty"
	}

	if req.PassWord == "" {
		return false, "[Login error]:PassWord is Emperty"
	}

	return true, ""
}

func loginUser(uid uint64, con netpoll.Connection) bool {
	local.StoreConnection(uid, con)
	return true
}

func isLogin(uid uint64) (bool, netpoll.Connection) {
	con := local.ConecntionForUid(uid)
	return con != nil, con
}
