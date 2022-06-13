package db

import (
	"douyin/common"
)

func SelectUserById(userId int64) (user common.User) {
	row := DB.QueryRow("select user_id, name, follow_count, follower_count, is_follow, password from tb_user where user_id = ?", userId)
	row.Scan(&user.Id, &user.Name, &user.FollowCount, &user.FollowerCount, &user.IsFollow, &user.Password)
	return user
}

func SelectUserIdBy(username string) (userId int64, err error) {
	row := DB.QueryRow("select user_id from tb_user where name = ?", username)
	err = row.Scan(&userId)
	if err != nil {
		return -1, err
	}
	return userId, err
}

func InsertUser(user common.User) (insertId int64, err error) {
	result, err := DB.Exec(
		"insert into tb_user(name,password) values(?,?)",
		user.Name, user.Password)
	if err != nil {
		return -1, err
	}
	user.Id, err = result.LastInsertId() //新增数据的ID
	if err != nil {
		return -1, err
	}
	return user.Id, err
}

func SelectUserIdNamePasswordBy(username string) (user common.User) {
	rows := DB.QueryRow(
		"select user_id,name,password from tb_user where name = ?",
		username)
	rows.Scan(&user.Id, &user.Name, &user.Password)
	return user
}

func SelectUserInfoBy(userId int64) (user common.User) {
	row := DB.QueryRow(
		"select user_id,name,follow_count,follower_count,is_follow from tb_user where user_id = ?", userId)
	row.Scan(&user.Id, &user.Name, &user.FollowCount, &user.FollowerCount, &user.IsFollow)
	return user
}
