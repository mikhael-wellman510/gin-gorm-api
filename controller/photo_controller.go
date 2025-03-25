package controller

import (
	"GinGonicGorm/constants"
	"GinGonicGorm/entity"
	"GinGonicGorm/utils"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type (
	PhotoController interface {
		UploadPhoto(ctx *gin.Context)
	}

	photoController struct {
	}
)

func NewPhotoController() PhotoController {

	return &photoController{}
}

func (pc *photoController) UploadPhoto(ctx *gin.Context) {

	file, err := ctx.FormFile("file")

	if err != nil {
		log.Println("Error kenapa  ? ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	log.Println("Hasil File : ", file)

	filePath := filepath.Join("uploads", file.Filename)

	log.Println("File Path : ", filePath)

	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		fail := utils.BuildResponseFailed(err.Error())
		ctx.JSON(http.StatusBadRequest, fail)
		return
	}

	nameUrl := constants.SERVER + "/uploads/" + file.Filename

	photo := entity.Photo{
		FileName: file.Filename,
		Url:      nameUrl,
	}

	log.Println("Hasil photo : ", photo)

	res := utils.BuildResponseSucces("Succes", photo)
	ctx.JSON(http.StatusOK, res)
}
