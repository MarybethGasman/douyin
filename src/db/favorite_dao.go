package db

//查询tb_favorite表判断是否存在一条记录
func RecordExists(userId, videoId int64) bool {
	row := DB.QueryRow("select favorite_id from tb_favorite where user_id = ? and video_id = ?", userId, videoId)
	favoriteId := -1
	row.Scan(&favoriteId)
	if favoriteId > 0 {
		return true
	}
	return false
}
