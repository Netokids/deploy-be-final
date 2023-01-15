package Articledto

type ArticleRequest struct {
	Title       string `json:"title" form:"title"`
	Image       string `json:"image" form:"image"`
	Description string `json:"description" form:"description"`
	UserID      int    `json:"user_id" form:"user_id"`
}
