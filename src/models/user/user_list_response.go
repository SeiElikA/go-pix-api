package user

type UserListResponse struct {
	TotalCount int             `json:"total_count"`
	Posts      []*UserResponse `json:"users"`
}
