package services

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.Asseter/internal/env"
	"go.Asseter/internal/util"
)

type StorageService struct {
	Config *env.Config
}

func NewStorageService(config *env.Config) *StorageService {
	return &StorageService{
		Config: config,
	}
}

func (e *StorageService) WriteFile(ctx *gin.Context) {
	fileData, err := ctx.FormFile("file")
	if err != nil {
		util.GenerateResponse(ctx, http.StatusConflict, "Failed to Upload file.", true, nil)
		return
	}
	isFileVaild := util.CheckFile(fileData.Filename)
	if !isFileVaild {
		util.GenerateResponse(ctx, http.StatusInternalServerError, "Uploaded invalid File.", true, nil)
		return
	}
	err = ctx.SaveUploadedFile(fileData, "/app/storage/"+fileData.Filename)
	if err != nil {
		util.GenerateResponse(ctx, http.StatusConflict, "Failed to Write file.", true, nil)
		return
	} else {
		util.GenerateResponse(ctx, http.StatusOK, "", false, nil)
		return
	}
}

func (e *StorageService) GetFile(ctx *gin.Context, name string) {
	isFileVaild := util.CheckFile(name)
	if !isFileVaild {
		util.GenerateResponse(ctx, http.StatusInternalServerError, "Uploaded invalid File.", true, nil)
		return
	}
	ctx.File("/app/storage/" + name)
}

func (e *StorageService) DeleteFile(ctx *gin.Context, name string) {
	isFileVaild := util.CheckFile(name)
	if !isFileVaild {
		util.GenerateResponse(ctx, http.StatusInternalServerError, "Uploaded invalid File.", true, nil)
		return
	}
	err := os.Remove("/app/storage/" + name)
	if err != nil {
		util.GenerateResponse(ctx, http.StatusConflict, "Failed to Delete file.", true, nil)
		return
	} else {
		util.GenerateResponse(ctx, http.StatusOK, "Successfull deleted file!", false, nil)
		return
	}
}
