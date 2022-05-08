package loginsrv

import (
	"practise_go_net/bz_server/base"
	"practise_go_net/bz_server/mod/user/userdao"
	"practise_go_net/bz_server/mod/user/userdata"
	"practise_go_net/common/async_op"
	"time"
)

//1. 一般用 callback 作为一个回调函数
//func LoginByPasswordAsync(userName string, password string, callback func(*userdata.User)) {
//2. 参考使用 async await 方式

func LoginByPasswordAsync(userName string, password string) *base.AsyncBizResult {
	if len(userName) <= 0 || len(password) <= 0 {
		return nil
	}

	bizResult := &base.AsyncBizResult{}

	async_op.Process(
		async_op.Str2BindId(userName),
		func() {
			//通过 dao 获取用户数据
			user := userdao.GetUserByName(userName)

			tNow := time.Now().UnixMilli()

			if user == nil {
				user = &userdata.User{
					UserName:   userName,
					Password:   password,
					CreateTime: tNow,
					HeroAvatar: "Hero_Hammer",
				}
			}

			//更新最后的登录时间
			user.LastLoginTime = tNow
			userdao.SaveOrUpdate(user)

			bizResult.SetReturnedObj(user)
		},
		nil,
	)

	return bizResult
}
