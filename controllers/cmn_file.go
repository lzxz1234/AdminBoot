package controllers

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"strings"
	"time"

	"github.com/nfnt/resize"

	"git.lzxz1234.cn/lzxz1234/AdminBoot/utils"
	"github.com/astaxie/beego"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

// CmnFileController 通用文件处理器
type CmnFileController struct {
	BaseController
}

var accessKey = beego.AppConfig.String("qiNiuAccesskey")
var secretKey = beego.AppConfig.String("qiNiuSecretKey")
var bucket = beego.AppConfig.String("qiNiuBucket")
var domain = beego.AppConfig.String("qiNiuDomain")

// Upload 通用文件上传
// @router /upload [post]
func (c *CmnFileController) Upload() {

	file, header, err := c.GetFile("file")
	var reader io.Reader = file
	var size = header.Size

	defer file.Close()

	if err != nil {
		c.Data["json"] = utils.NewResult(1, "上传失败", err)
		c.ServeJSON()
		return
	}
	suffix := "tmp"
	if header.Filename != "" && strings.Contains(header.Filename, ".") {
		suffix = header.Filename[strings.LastIndex(header.Filename, ".")+1:]
		suffix = strings.ToLower(suffix)
	}
	if suffix == "jpg" || suffix == "png" {
		img, _, err := image.Decode(file)
		if err == nil {
			img = resize.Resize(0, 0, img, resize.Lanczos3)
			buf := new(bytes.Buffer)
			jpeg.Encode(buf, img, &jpeg.Options{Quality: 40})
			reader = buf
			size = int64(buf.Len())
		}
	}

	key := fmt.Sprintf("%d.%s", time.Now().UnixNano(), suffix)
	putPolicy := storage.PutPolicy{
		Scope: fmt.Sprintf("%s:%s", bucket, key),
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	uploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": header.Filename,
		},
	}

	err = uploader.Put(context.Background(), &ret, upToken, key, reader, size, &putExtra)

	if err != nil {
		c.Data["json"] = utils.NewResult(1, "上传失败", err)
		c.ServeJSON()
		return
	}
	c.Data["json"] = utils.NewResult(0, "上传成功", map[string]string{
		"src": storage.MakePublicURL(domain, ret.Key),
	})
	c.ServeJSON()
}
