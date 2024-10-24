package post

type PostListResponse struct {
	TotalCount int             `json:"total_count"`
	Posts      []*PostResponse `json:"posts"`
}
