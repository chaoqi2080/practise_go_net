package userdata

type userGroup struct {
	innerMap map[int64]*User
}

var userGroupInstance = &userGroup{
	innerMap: make(map[int64]*User),
}

//GetUserGroup 获取用户组
func GetUserGroup() *userGroup {
	return userGroupInstance
}

func (group *userGroup) Add(user *User) {
	if nil == user {
		return
	}

	group.innerMap[user.UserId] = user
}

func (group *userGroup) RemoveByUserId(userId int64) {
	if userId <= 0 {
		return
	}

	delete(group.innerMap, userId)
}

func (group *userGroup) GetByUserId(userId int64) *User {
	if userId <= 0 {
		return nil
	}

	return group.innerMap[userId]
}

// GetUserAll 获取所有用户
func (group *userGroup) GetUserAll() map[int64]*User {
	return group.innerMap
}
