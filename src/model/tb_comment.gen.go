// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameComment = "tb_comment"

// Comment mapped from table <tb_comment>
type Comment struct {
	CommentID int64  `gorm:"column:comment_id;primaryKey;autoIncrement:true" json:"comment_id"`
	Username  string `gorm:"column:username" json:"username"`
	Content   string `gorm:"column:content" json:"content"`
	Isdeleted int32  `gorm:"column:isdeleted" json:"isdeleted"`
}

// TableName Comment's table name
func (*Comment) TableName() string {
	return TableNameComment
}