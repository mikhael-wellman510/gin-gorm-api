package controller

import (
	"GinGonicGorm/constants"
	"GinGonicGorm/entity"
	"GinGonicGorm/utils"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type (
	PhotoController interface {
		UploadPhoto(ctx *gin.Context)
		GetPhoto(ctx *gin.Context)
	}

	photoController struct {
	}
)

func NewPhotoController() PhotoController {

	return &photoController{}
}

func (pc *photoController) GetPhoto(ctx *gin.Context) {
	fileName := ctx.Param("filename")

	filePath := filepath.Join("uploads", fileName)

	log.Println("Hasil File path : ", filePath)

	fileInfo, err := os.Stat(filePath)
	log.Println("err : ", err)

	if os.IsNotExist(err) {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	log.Println("file info : ", fileInfo)
	var url string = constants.SERVER + "/uploads/" + fileName
	log.Println("hasil : ", url)
	res := entity.PhotoResponse{
		Url:       url,
		CreatedAt: fileInfo.ModTime(),
	}

	result := utils.BuildResponseSucces("Sukses", res)

	ctx.JSON(http.StatusOK, result)

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
