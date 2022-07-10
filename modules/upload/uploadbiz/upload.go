package uploadbiz

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"path/filepath"
	"simple-rest-api/common"
	"simple-rest-api/component/uploadprovider"
	"simple-rest-api/modules/upload/uploadmodel"
	"strings"
	"time"
)

type CreateImageStorage interface {
	CreateImage(context context.Context, data *common.Image) error
}

type uploadBiz struct {
	store    CreateImageStorage
	provider uploadprovider.UploadProvider
}

func NewUploadBiz(provider uploadprovider.UploadProvider, store CreateImageStorage) *uploadBiz {
	return &uploadBiz{store: store, provider: provider}
}

func (biz *uploadBiz) Upload(ctx context.Context, data []byte,
	folder, fileName string,
) (*common.Image, error) {
	fileBytes := bytes.NewBuffer(data)
	w, h, err := getImageDimension(fileBytes)

	if err != nil {
		return nil, uploadmodel.ErrFileIsNotImage(err)
	}

	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	fileExt := filepath.Ext(fileName)
	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt)

	img, err := biz.provider.SaveFileUploaded(ctx, data,
		fmt.Sprintf("%s/%s", folder, fileName))
	if err != nil {
		return nil, uploadmodel.ErrCannotSaveFile(err)
	}
	img.Width = w
	img.Height = h
	img.CloudName = "s3"
	img.Extension = fileExt

	if err := biz.store.CreateImage(ctx, img); err != nil {
		//delete image on S3
		return nil, uploadmodel.ErrCannotSaveFile(err)
	}

	return img, nil

}

func getImageDimension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		img, err = jpeg.DecodeConfig(reader)
		if err != nil {
			return 0, 0, err
		}
		return img.Width, img.Height, nil

	}
	return img.Width, img.Height, nil
}
