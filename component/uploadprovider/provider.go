package uploadprovider

import (
	"context"
	"g05-food-delivery/common"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, dataBytes []byte, dst string) (*common.Image, error)
}
