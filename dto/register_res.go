package dto

type RegisterRes struct {
	ID       int64  `json:"id"`
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
	Username string `json:"username"`
}
