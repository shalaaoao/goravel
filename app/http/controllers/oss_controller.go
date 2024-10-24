package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"goravel/app/http/response"
	"io"
	"os"
	"path/filepath"
)

type OssController struct {
	//Dependent services
}

func NewOssController() *OssController {
	return &OssController{
		//Inject services
	}
}

func (r *OssController) Show(ctx http.Context) http.Response {
	return ctx.Response().Success().Json(http.Json{
		"Hello": "Goravel",
	})
}

func (r *OssController) Update(ctx http.Context) http.Response {

	file, err := ctx.Request().File("name")
	if err != nil || file == nil {
		return ctx.Response().Status(500).Json(http.Json{
			"err": err.Error(),
		})
	}

	// 把file转移到 /mnt/Fanxiang2T/oss/ 目录下
	dst := "/mnt/Fanxiang2T/oss/" + filepath.Base(file.GetClientOriginalName()) // use the original filename
	sourceFile, err := os.Open(file.File())
	if err != nil {
		return ctx.Response().Status(500).Json(http.Json{
			"err": err.Error(),
		})
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		return ctx.Response().Status(500).Json(http.Json{
			"err": err.Error(),
		})
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return ctx.Response().Status(500).Json(http.Json{
			"err": err.Error(),
		})
	}

	err = os.Remove(file.File())
	if err != nil {
		return ctx.Response().Status(500).Json(http.Json{
			"err": err.Error(),
		})
	}

	// 返回公网、内网的下载链接
	ossPublicUrl := "http://nas.julyaoao.top:8090/" + file.GetClientOriginalName()
	ossPrivateUrl := "http://192.168.100.15:8090/" + file.GetClientOriginalName()

	return response.Success(ctx, http.Json{
		"ossPublicUrl":  ossPublicUrl,
		"ossPrivateUrl": ossPrivateUrl,
	})
}
