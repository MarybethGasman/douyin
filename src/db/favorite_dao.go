package db

import "database/sql"

// RecordExists 查询tb_favorite表判断是否存在一条记录
func RecordExists(userId, videoId int64) bool {
	row := DB.QueryRow("select favorite_id from tb_favorite where user_id = ? and video_id = ?", userId, videoId)
	favoriteId := -1
	row.Scan(&favoriteId)
	if favoriteId > 0 {
		return true
	}
	return false
}

//更新tb_favorite某条记录的is_deleted字段
func updateIsDeletedBy(tx sql.Tx, userId, videoId int64, isDeleted bool) (sql.Tx, error) {
	_, err := tx.Exec(
		"update tb_favorite set is_deleted = ? where user_id = ? and video_id = ?",
		isDeleted, userId, videoId)
	if err != nil {
		return tx, err
	}
	return tx, err
}
