package upload

import (
	UploadModels "api-go/app/models/upload"
	"api-go/lib/file"
	"api-go/lib/jwt"
	"api-go/lib/orm"
	"api-go/lib/page"
	"api-go/lib/response"
	"github.com/gin-gonic/gin"
)

type Image struct {
}

// Upload 上传图片
func (t *Image) Upload(c *gin.Context) {
	appG := &response.Gin{C: c}
	image, err := c.FormFile("image")
	if err != nil {
		appG.ErrorResponse(response.CodeErrorUploadImage)
		return
	}

	if !file.CheckImageExt(image.Filename) {
		appG.ErrorResponse(response.CodeErrorUploadImage)
		return
	}

	imageSaveDir := file.ImageStaticDirAdd("")
	imageSaveName := file.ImageSaveName(image.Filename)
	imageSavePath := imageSaveDir + imageSaveName

	err = c.SaveUploadedFile(image, imageSavePath)
	if err != nil {
		appG.ErrorResponse(response.CodeErrorUploadImage)
		return
	}

	imageSaveUrl := file.ImagePathToUrl(imageSavePath)

	imageModel := &UploadModels.Image{
		Url:           imageSaveUrl,
		Path:          imageSavePath,
		CreatedUserId: jwt.AdminId(c),
	}

	res := orm.DB().Create(imageModel)
	if res.Error != nil {
		appG.ErrorResponse(response.CodeErrorUploadImage)
		return
	}

	data := map[string]interface{}{
		"image_id": imageModel.ID,
		"path":     imageSavePath,
		"url":      imageSaveUrl,
	}
	appG.SuccessResponse(response.CodeSuccess, data)
}

// SecretUpload 密钥上传图片
func (t *Image) SecretUpload(c *gin.Context) {
	appG := &response.Gin{C: c}

	secret := c.Query("secret")
	if secret == "" {
		appG.ErrorMsgResponse("密钥不为空")
		return
	}

	if secret != "123456789" {
		appG.ErrorMsgResponse("密钥错误")
		return
	}

	image, err := c.FormFile("image")
	if err != nil {
		appG.ErrorResponse(response.CodeErrorUploadImage)
		return
	}

	if !file.CheckImageExt(image.Filename) {
		appG.ErrorResponse(response.CodeErrorUploadImage)
		return
	}

	imageSaveDir := file.ImageStaticDirAdd("")
	imageSaveName := file.ImageSaveName(image.Filename)
	imageSavePath := imageSaveDir + imageSaveName

	err = c.SaveUploadedFile(image, imageSavePath)
	if err != nil {
		appG.ErrorResponse(response.CodeErrorUploadImage)
		return
	}

	imageSaveUrl := file.ImagePathToUrl(imageSavePath)

	imageModel := &UploadModels.Image{
		Url:           imageSaveUrl,
		Path:          imageSavePath,
		CreatedUserId: 0,
	}

	res := orm.DB().Create(imageModel)
	if res.Error != nil {
		appG.ErrorResponse(response.CodeErrorUploadImage)
		return
	}

	data := map[string]interface{}{
		"image_id": imageModel.ID,
		"path":     imageSavePath,
		"url":      imageSaveUrl,
	}
	appG.SuccessResponse(response.CodeSuccess, data)
}

// 图床
func (t *Image) List(c *gin.Context) {
	appG := &response.Gin{C: c}
	var result []map[string]interface{}
	orm.DB().Model(&UploadModels.Image{}).
		Select(
			"id",
			"url",
		).
		Order("id desc").
		Scopes(page.Paginate(c)).
		Find(&result)

	appG.SuccessResponse(response.CodeSuccess, result)
}

// 更新所有图片链接
func (t *Image) UpdAllUrl(c *gin.Context) {
	var imageList []*UploadModels.Image
	orm.DB().Model(&UploadModels.Image{}).
		Select(
			"id",
			"path",
		).
		Order("id asc").
		Find(&imageList)

	for i := 1; i < len(imageList); i++ {
		id := imageList[i].ID
		url := file.ImagePathToUrl(imageList[i].Path)
		orm.DB().Model(&UploadModels.Image{}).Where("id = ?", id).Update("url", url)
	}

	appG := &response.Gin{C: c}
	appG.SuccessResponse(response.CodeSuccess)
}
