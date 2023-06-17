package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	// 查数据库 查找到所有的community 并返回
	return mysql.GetCommunityList()
}

func GetCommunityDetailByID(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetailByID(id)
}
