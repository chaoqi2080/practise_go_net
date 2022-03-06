package userdao

import (
	"practise_go_net/bz_server/base"
	"practise_go_net/bz_server/mod/user/userdata"
	"practise_go_net/common/log"
)

const sqlSaveOrUpdate = `
insert into t_user(
    user_name, password, hero_avatar, create_time, last_login_time
) value (
    ?, ?, ?, ?, ?
)
on duplicate key update last_login_time = ?
`

func SaveOrUpdate(user *userdata.User) {
	if user == nil {
		return
	}

	stmt, err := base.MysqlDB.Prepare(sqlSaveOrUpdate)
	if err != nil {
		log.Error("%+v", err)
		return
	}

	result, err := stmt.Exec(
		user.UserName,
		user.Password,
		user.HeroAvatar,
		user.CreateTime,
		user.LastLoginTime,
		user.LastLoginTime,
	)

	if err != nil {
		log.Error("%+v", err)
		return
	}

	rowId, err := result.LastInsertId()
	
	if err != nil {
		log.Error("%+v", err)
		return
	}

	user.UserId = rowId
}
