package file

import (
	"api-go/lib/config"
	"api-go/lib/md5"
	"github.com/unknwon/com"
	"strings"
	"time"
)

// ImageStaticDir 图片资源保存路径
func ImageStaticDir() string {
	return config.AppConfig.RuntimeRootPath + "/" + config.ImageConfig.SavePath + "/"
}

// ImageStaticDirAdd 图片资源创建文件夹并返回路径
func ImageStaticDirAdd(dir string) string {
	if dir == "" {
		imageStaticDir := ImageStaticDir()
		return imageStaticDir
	}
	imageStaticDir := ImageStaticDir() + dir
	_ = IsNotExistMkDir(imageStaticDir)
	return imageStaticDir
}

// ImageUrlDir 图片资源 URL 路径
func ImageUrlDir() string {
	return config.ImageConfig.PrefixUrl + "/" + config.ImageConfig.StaticUrl + "/"
}

// ImageUrl 根据图片名称生成 url
func ImageUrl(fileName string) string {
	return ImageUrlDir() + fileName
}

// ImagePathToUrl 根据图片静态路径转为URL
func ImagePathToUrl(path string) string {
	return strings.Replace(path, ImageStaticDir(), ImageUrlDir(), 1)
}

// 检查图片后缀名合法性
func CheckImageExt(fileName string) bool {
	allowExtList := strings.Split(config.ImageConfig.AllowExt, ",")
	imageExt := Ext(fileName)
	for _, allowExt := range allowExtList {
		if allowExt == imageExt {
			return true
		}
	}
	return false
}

// ImageSaveName 生成图片名称
func ImageSaveName(fileName string) string {
	return md5.ToMD5(com.ToStr(time.Now().Unix())+fileName) + Ext(fileName)
}
