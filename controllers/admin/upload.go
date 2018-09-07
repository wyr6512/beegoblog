package admin

import (
	"path"

	"github.com/astaxie/beego"
)

type UploadController struct {
	beego.Controller
}

const (
	Protocol = "http://"
)

func (this *UploadController) Upload() {
	f, h, _ := this.GetFile("upload_file") //获取上传的文件
	filepath := Protocol + this.Ctx.Request.Host + path.Join("/static/upload/images", h.Filename)
	f.Close()                                //关闭上传的文件，不然的话会出现临时文件不能清除的情况
	this.SaveToFile("upload_file", filepath) //存文件
	var images = []string{filepath}
	this.Data["json"] = map[string]interface{}{"errno": 0, "data": images}
	this.ServeJSON()
	this.StopRun()
}
