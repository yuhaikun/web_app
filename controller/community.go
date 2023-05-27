package controller

import (
	"strconv"
	"web_app/logic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 参数处理，路由跳转
func CommunityHandler(c *gin.Context) {
	//	查询到所有的社区（community_id，community_name）以列表的形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不轻易把服务端报错暴露给外面
	}
	ResponseSuccess(c, data)
}

// CommunityDetailHandler 社区分类详情
func CommunityDetailHandler(c *gin.Context) {
	//	1.获取社区id
	communityID := c.Param("id")
	id, err := strconv.ParseInt(communityID, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
	}

	//	查询到所有的社区（community_id，community_name）以列表的形式返回
	data, err := logic.GetCommunityDetailByID(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityByID() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不轻易把服务端报错暴露给外面
	}
	ResponseSuccess(c, data)
}
