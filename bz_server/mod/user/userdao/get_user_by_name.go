package userdao

import (
	"practise_go_net/bz_server/base"
	"practise_go_net/bz_server/mod/user/userdata"
	"practise_go_net/common/log"
)

const sqlGetUserByName = `select user_id, user_name, password, hero_avatar from t_user where user_name = ?`

func GetUserByName(userName string) *userdata.User {
	if len(userName) <= 0 {
		return nil
	}

	row := base.MysqlDB.QueryRow(sqlGetUserByName, userName)
	if row == nil {
		return nil
	}

	user := &userdata.User{}

	err := row.Scan(
		&user.UserId,
		&user.UserName,
		&user.Password,
		&user.HeroAvatar,
	)

	if err != nil {
		log.Error(
			"解析查询出的 user 数据失败 %+v",
			err,
		)
		return nil
	}

	return user
}
