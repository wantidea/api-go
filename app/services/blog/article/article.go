package article

import (
	"api-go/app/models/blog"
	"api-go/app/services/redis"
	"api-go/lib/orm"
)

type Article struct {
}

// Total 读取文章数量
func (t *Article) Total() int64 {
	var total int64
	article := &blog.Article{}
	orm.DB().Model(article).Count(&total)
	return total
}

// TotalOnRedis 返回文章数量 读取 Redis
func (t *Article) TotalOnRedis() int64 {
	total, err := redis.ArticleTotal()
	if err != nil {
		total = t.Total()
		err = redis.SetArticleTotal(total)
	}
	return total
}
