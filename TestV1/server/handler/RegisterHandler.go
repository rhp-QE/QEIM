package handler

import (
	"github.com/cloudwego/netpoll"
	uuidUtil "github.com/google/uuid"
	pbMessage "qeim.com/testv1/pb/generate"
)

var (
	//uid 到 userName 的映射
	uidToUserNameDict map[uint64]string = make(map[uint64]string, 0)

	//userName 到 uid 的映射
	userNameToUidDict map[string]uint64 = make(map[string]uint64, 0)

	//uid 到密码的映射
	uidToPassWordDict map[uint64]string = make(map[uint64]string, 0)

	//uid到 con 的映射
	uidToConnection  map[uint64]netpoll.Connection = make(map[uint64]netpoll.Connection)
)

func handleRegister(registerRquest *pbMessage.RegisterRequest) *pbMessage.RegisterResponse {
	//1、前置检查参数是否合法
	if ok, res := preCheck(registerRquest); ok ==false {
		return res
	}

	//2、检查是否注册过
	localUid := userNameToUidDict[registerRquest.UserName]
	if localUid != 0 {
		return &pbMessage.RegisterResponse{
			Ok: false,
			CodeState: 0,
			StringState: "you have registed",
			Uid: localUid,
		}
	}

	localUid = uint64(uuidUtil.New().ID())
	
	// TODO: 检查是否保存到数据库
	saveUser(localUid, registerRquest)

	return &pbMessage.RegisterResponse{
		Ok: true,
		CodeState: 1,
		StringState: "register success",
		Uid: localUid,
	}
}



//前置检查（语言、系统级别）：参数是否合法
func preCheck(registerRquest *pbMessage.RegisterRequest) (bool, *pbMessage.RegisterResponse) {
	ok := (registerRquest.Password != "") && (registerRquest.UserName != "")
	
	if (registerRquest.UserName == "") {
		return ok, &pbMessage.RegisterResponse{
			Ok: false,
			CodeState: 0,
			StringState: "userName can not be emperty!!!",
			Uid: 0,
		}
	}

	if (registerRquest.Password == "") {
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

func saveUser(uid uint64, req *pbMessage.RegisterRequest) bool {
	userNameToUidDict[req.UserName] = uid
	uidToPassWordDict[uid] = req.Password
	uidToUserNameDict[uid] = req.UserName
	
	return true
}


func isRegist(uid uint64) bool {
	_, found := uidToUserNameDict[uid]
	return found
}