package category

import (
	ArticleModels "api-go/app/models/blog/article"
	"api-go/lib/orm"
)

func NameById(id int64) string {
	category := &ArticleModels.Category{}
	res := orm.DB().Model(category).Where("id = ?", id).Select("name").First(&category)
	if res.Error != nil {
		return ""
	}
	return category.Name
}

func NameListByIds(ids []int64) map[int64]string {
	result := map[int64]string{}
	categoryList := []*ArticleModels.Category{}
	category := &ArticleModels.Category{}

	res := orm.DB().Model(category).Where("id in (?)", ids).Select(
		"id",
		"name",
	).Find(&categoryList)
	if res.Error != nil {
		return result
	}

	for _, item := range categoryList {
		result[item.ID] = item.Name
	}
	return result
}
