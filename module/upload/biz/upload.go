package bizupload

import (
	"bytes"
	"context"
	"fmt"
	"g05-food-delivery/common"
	"g05-food-delivery/component/uploadprovider"
	uploadmodel "g05-food-delivery/module/upload/model"
	"image"
	_ "image/gif"  // Register GIF format
	_ "image/jpeg" // Register JPEG format
	_ "image/png"  // Register PNG format
	"io"
	"log"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type CreateImageStorage interface {
	CreateImage(context context.Context, data *common.Image)
}
type UploadBiz struct {
	provider uploadprovider.UploadProvider
	imgStore CreateImageStorage
}

func NewUploadBiz(provider uploadprovider.UploadProvider, imgStore CreateImageStorage) *UploadBiz {
	return &UploadBiz{provider: provider, imgStore: imgStore}
}

func (biz *UploadBiz) Upload(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error) {
	// Create a copy of the data for dimension checking
	fileBytes := bytes.NewReader(data)

	w, h, err := getImageDimension(fileBytes)

	if err != nil {
		return nil, uploadmodel.ErrFileIsNotImage(err)
	}

	// Rest of the code remains the same
	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	fileExt := filepath.Ext(fileName)
	// If extension is empty or not recognized, force it to .png
	if fileExt == "" || !isValidImageExt(fileExt) {
		fileExt = ".png"
	}

	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt)

	img, err := biz.provider.SaveFileUploaded(ctx, data, path.Join(folder, fileName))

	if err != nil {
		return nil, uploadmodel.ErrCannotSaveFile(err)
	}

	img.Width = w
	img.Height = h
	img.Extension = fileExt

	return img, nil
}

func getImageDimension(fileBytes io.Reader) (int, int, error) {
	img, format, err := image.DecodeConfig(fileBytes)

	if err != nil {
		log.Printf("Error decoding image: %v, format: %s", err, format)
		return 0, 0, err
	}

	log.Printf("Successfully decoded image format: %s, dimensions: %dx%d", format, img.Width, img.Height)
	return img.Width, img.Height, nil
}

// Helper function to check if file extension is valid for images
func isValidImageExt(ext string) bool {
	ext = strings.ToLower(ext)
	validExts := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
	}
	return validExts[ext]
}
