package userdata

import "fmt"

type User struct {
	UserId         int64  `db:"user_id"`
	UserName       string `db:"user_name"`
	Password       string `db:"password"`
	HeroAvatar     string `db:"hero_avatar"`
	CurrHp         int32  `db:"curr_hp"`
	CreateTime     int64  `db:"create_time"`
	LastLoginTime  int64  `db:"last_login_time"`
	MoveState      *MoveState
	lastUpdateTime int64
}

func (user *User) GetLsoId() string {
	return fmt.Sprintf("User_%v", user.UserId)
}

func (user *User) GetLastUpdateTime() int64 {
	return user.lastUpdateTime
}

func (user *User) SetLastUpdateTime(val int64) {
	user.lastUpdateTime = val
}
