package models

import "time"

// 内存对齐概念
type Post struct {
	ID          int64     `json:"id,string" db:"post_id"`
	AuthorID    int64     `json:"author_id" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"`
	Status      int32     `json:"status" db:"status"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	CreateTime  time.Time `json:"createTime" db:"create_time"`
}

type ApiPostDetail struct {
	AuthorName       string             `json:"authorName"`
	VotesNum         int64              `json:"votes_num"`
	*Post                               // 嵌入帖子结构体
	*CommunityDetail `json:"community"` // 社区信息
}
