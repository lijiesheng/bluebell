package mysql

import (
	"bluebell/models"
	"database/sql"
	"go.uber.org/zap"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id, community_name from community"
	if err := DB.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}
	return
}

// 根据 ID 查询社区详情
func GetCommunityDetailByID(id int64) (community *models.CommunityDetail, err error) {
	community = new(models.CommunityDetail)
	sqlStr := `select community_id, community_name, introduction, 
		create_time, from community where community_id = ?`
	if err := DB.Get(community, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Error("there is no community detail in db ")
			err = ErrorInvalidID
		}
	}
	return community, err
}

