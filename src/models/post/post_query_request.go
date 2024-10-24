package post

type PostQueryRequest struct {
	OrderBy      string `form:"order_by" binding:"omitempty,oneof=created_at like_count"`
	OrderType    string `form:"order_type" binding:"omitempty,oneof=asc desc"`
	Content      string `form:"content" binding:"-"`
	Tag          string `form:"tag" binding:"-"`
	LocationName string `form:"location_name" binding:"-"`
	Page         int64  `form:"page" binding:"omitempty,min=1"`
	PageSize     int64  `form:"page_size" binding:"omitempty,min=1,max=100"`
}
