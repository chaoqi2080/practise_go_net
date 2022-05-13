package userlso

import (
	"fmt"
	"practise_go_net/bz_server/mod/user/userdao"
	"practise_go_net/bz_server/mod/user/userdata"
	"practise_go_net/common/async_op"
)

type UserLso struct {
	UserRef *userdata.User
}

func (lso *UserLso) GetLsoId() string {
	return fmt.Sprintf("User_%v", lso.UserRef.UserId)
}

func (lso *UserLso) SaveOrUpdate() {
	//循环引用问题，引入一个第三者，通过他分别调用 userdata, userdao
	async_op.Process(
		int(lso.UserRef.UserId),
		func() {
			userdao.SaveOrUpdate(lso.UserRef)
		},
		nil,
	)
}
