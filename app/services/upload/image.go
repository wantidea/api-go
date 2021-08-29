package upload

import (
	"api-go/app/models/upload"
	"api-go/lib/file"
	"api-go/lib/orm"
)

// ImagePath 根据 id 取图片路径
func ImagePath(id int64) string {
	result := map[string]interface{}{}
	res := orm.DB().Model(upload.Image{}).Select("path").Where("id = ?", id).First(&result)
	if res.Error != nil || result["path"] == nil {
		return ""
	}
	return result["path"].(string)
}

// ImageUrl 根据 id 取图片链接
func ImageUrl(id int64) string {
	if id <= 0 {
		return ""
	}

	result := map[string]interface{}{}
	res := orm.DB().Model(upload.Image{}).Select("path").Where("id = ?", id).First(&result)
	if res.Error != nil || result["path"] == nil {
		return ""
	}
	return file.ImagePathToUrl(result["path"].(string))
}
