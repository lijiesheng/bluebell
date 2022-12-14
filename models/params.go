package models

// 定义请求的参数结构体

const (
	OrderTime  = "time"
	OrderScore = "score"
)

// ParamSignUp 注册请求参数
type ParamSignUp struct {
	Username string `json:"username" binding:"required"`
	// binding:"required" 不能传入零值，比如字符串的不能传入空串，int类型的不能传入0，bool类型的不能传入false
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}

// ParamLogin 登录请求参数  POST 请求
type ParamLogin struct {
	// Get 请求
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`

	// POST 请求
	//Username string `json:"username" binding:"required" form:"username"`
	//Password string `json:"password" binding:"required" form:"password"`
}

// ParamVoteData 投票数据
type ParamVoteData struct {
	PostID    string `json:"post_id" binding:"required"`
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1"` // 赞成票(1)还是反对票(-1)取消投票(0)
}

// ParamPostList 获取帖子列表query string参数
type ParamPostList struct {
	CommunityID int64  `json:"community_id" form:"community_id"`   // 可以为空
	Page        int64  `json:"page" form:"page" example:"1"`       // 页码
	Size        int64  `json:"size" form:"size" example:"10"`      // 每页数据量
	Order       string `json:"order" form:"order" example:"score"` // 排序依据
}
