package ginupload

import (
	"g05-food-delivery/common"
	"g05-food-delivery/component/appctx"
	bizupload "g05-food-delivery/module/upload/biz"
	"github.com/gin-gonic/gin"
	"net/http"
)

//func UploadImage(appCtx appctx.AppContext) func(c *gin.Context) {
//	return func(c *gin.Context) {
//		fileHeader, err := c.FormFile("image")
//
//		if err != nil {
//			panic(err)
//		}
//
//		if err := c.SaveUploadedFile(fileHeader, fmt.Sprintf("static/%s", strings.TrimSpace(fileHeader.Filename))); err != nil {
//			panic(err)
//		}
//
//		c.JSON(http.StatusOK, common.SimpleSuccessResponse(common.Image{
//			Id:        0,
//			Url:       "localhost:8080/static/" + strings.TrimSpace(fileHeader.Filename),
//			Width:     0,
//			Height:    0,
//			CloudName: "local",
//			Extension: "png",
//		},
//		))
//	}
//}

func UploadImage(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		folder := c.DefaultQuery("folder", "img")

		file, err := fileHeader.Open()

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		defer file.Close() //we can close here

		dataBytes := make([]byte, fileHeader.Size)

		if _, err = file.Read(dataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		biz := bizupload.NewUploadBiz(appCtx.UploadProvider(), nil)

		img, err := biz.Upload(c.Request.Context(), dataBytes, folder, fileHeader.Filename)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(img))
	}
}
