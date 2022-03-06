package loginsrv

import (
	"practise_go_net/bz_server/mod/user/userdao"
	"practise_go_net/bz_server/mod/user/userdata"
)

func LoginByPasswordAsync(userName string, password string) *userdata.User {
	if len(userName) <= 0 || len(password) <= 0 {
		return nil
	}

	user := userdao.GetUserByName(userName)

	//需要检查密码

	return user
}
