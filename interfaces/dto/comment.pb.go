package dto

import "time"

type H struct{}

type Comment struct {
	Id         int64     `json:"id"`
	Pid        int64     `json:"pid"`
	ArticleId  int64     `json:"article_id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Url        string    `json:"url"`
	Content    string    `json:"content"`
	CreateTime time.Time `json:"create_time"`
}
