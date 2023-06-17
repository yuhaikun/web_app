package models

import "time"

type Community struct {
	ID   int64  `json:"id" db:"community_id"`
	Name string `json:"name" db:"community_name"`
}

type CommunityDetail struct {
	ID          int64     `json:"id" db:"community_id"`
	Name        string    `json:"name" db:"community_name"`
	Instruction string    `json:"introduction,omitempty" db:"introduction"`
	CreateTime  time.Time `json:"createTime" db:"create_time"`
}
