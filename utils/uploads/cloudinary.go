package uploads

import (
	"BE-REPO-20/app/configs"
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type CloudinaryInterface interface {
	Upload(file multipart.File, nameFile string, folderName string) (*uploader.UploadResult, error)
	Destroy(publicId string) (*uploader.DestroyResult, error)
}

type cloud struct{}

func NewCloudService() CloudinaryInterface {
	return &cloud{}
}

// Upload implements CloudinaryInterface.
func (c *cloud) Upload(file multipart.File, nameFile string, folderName string) (*uploader.UploadResult, error) {
	ctx := context.Background()

	cldUrl := fmt.Sprintf("cloudinary://%s:%s@%s", configs.CLOUDINARY_KEY, configs.CLOUDINARY_SECRET, configs.CLOUDINARY_CLOUD_NAME)

	cldService, _ := cloudinary.NewFromURL(cldUrl)
	url, errUpl := cldService.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID: string(time.Now().Format("2006-01-02-15:04:05")) + nameFile,
		// Folder:      "img/items",
		Folder:      folderName,
		DisplayName: nameFile,
	})
	if errUpl != nil {
		return nil, errUpl
	}

	return url, nil
}

func (c *cloud) Destroy(publicId string) (*uploader.DestroyResult, error) {
	ctx := context.Background()

	cldUrl := fmt.Sprintf("cloudinary://%s:%s@%s", configs.CLOUDINARY_KEY, configs.CLOUDINARY_SECRET, configs.CLOUDINARY_CLOUD_NAME)
	cldService, _ := cloudinary.NewFromURL(cldUrl)
	url, errDest := cldService.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID:     publicId,
		Type:         "",
		ResourceType: "",
		Invalidate:   new(bool),
	})
	if errDest != nil {
		return nil, errDest
	}

	return url, nil
}

func (c *cloud) Remove(file multipart.File, nameFile string) error {
	ctx := context.Background()

	cld, _ := cloudinary.NewFromParams(configs.CLOUDINARY_CLOUD_NAME, configs.CLOUDINARY_KEY, configs.CLOUDINARY_SECRET)
	result, _ := cld.Admin.DeleteAssets(ctx, admin.DeleteAssetsParams{
		AssetType:       "image",
		DeliveryType:    "upload",
		PublicIDs:       []string{"vams9gc4ao9vsrljzwm7"},
		KeepOriginal:    new(bool),
		Invalidate:      new(bool),
		Transformations: "",
		NextCursor:      "",
	})
	log.Println(result)
	return nil
}
