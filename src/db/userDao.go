package db

import "douyin/src/common"

func SelectUserById(userId int64) (user common.User) {
	row := DB.QueryRow("select user_id, name, follow_count, follower_count, is_follow, password from tb_user where user_id = ?", userId)
	row.Scan(&user.Id, &user.Name, &user.FollowCount, &user.FollowerCount, &user.IsFollow, &user.Password)
	return user
}
