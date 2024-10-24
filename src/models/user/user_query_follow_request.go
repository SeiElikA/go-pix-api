package user

type UserQueryFollowRequest struct {
	OrderBy   string `form:"order_by" binding:"omitempty,oneof=created_at like_count"`
	OrderType string `form:"order_type" binding:"omitempty,oneof=asc desc"`
	Page      int64  `form:"page" binding:"omitempty,min=1"`
	PageSize  int64  `form:"page_size" binding:"omitempty,min=1,max=100"`
}
