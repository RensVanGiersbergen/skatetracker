package models

type AddBoard struct {
	UserId   string `json:"userId"`
	Nickname string `json:"nickname" binding:"required"`
	Brand    string `json:"brand" binding:"required"`
}
