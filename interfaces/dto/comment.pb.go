package dto

type H struct{}

type Comment struct {
	Pid       int64  `json:"pid"`
	ArticleId int64  `json:"article_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Url       string `json:"url"`
	Content   string `json:"content"`
}
