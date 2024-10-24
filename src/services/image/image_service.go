package image

import (
	"context"
	"fmt"
	"go-pix-api/src/config"
	"go-pix-api/src/entity"
	"go-pix-api/src/models/image"
	"go-pix-api/src/services"
	"go-pix-api/src/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

type ImageService struct {
	services.BaseService
}

func NewImageService() *ImageService {
	service := new(ImageService)
	service.Collection = config.DB.Collection("images")
	return service
}

func (service *ImageService) InsertManyImage(saveFolder string, images []*multipart.FileHeader) ([]int64, error) {
	var imageEntities []entity.Image
	var nextId = service.GetNextID()
	var insertImageIds []int64
	for idx, image := range images {
		path, file, err := service.SaveImage(saveFolder, image)
		width, height, err := utils.GetImageDimensions(file)
		if err != nil {
			return nil, err
		}

		var id = nextId + int64(idx)
		insertImageIds = append(insertImageIds, id)
		imageEntities = append(imageEntities, entity.Image{
			ID:        id,
			Url:       path,
			Width:     width,
			Height:    height,
			CreatedAt: time.Now().String(),
		})
	}

	var _, err = service.Collection.InsertMany(context.TODO(), imageEntities)

	return insertImageIds, err
}

func (service *ImageService) FindImageById(id int64) (*image.ImageResponse, error) {
	var data entity.Image
	err := service.Collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&data)
	var res = image.ImageResponse{
		ID:        data.ID,
		Url:       data.Url,
		Width:     data.Width,
		Height:    data.Height,
		CreatedAt: data.CreatedAt,
	}
	return &res, err
}

func (service *ImageService) FindManyImageById(ids []int64) ([]*image.ImageResponse, error) {
	filter := bson.M{"_id": bson.M{"$in": ids}}

	var ctx = context.TODO()
	cursor, err := service.Collection.Find(ctx, filter, options.Find())
	if err != nil {
		return nil, err
	}

	var images []*image.ImageResponse
	for cursor.Next(ctx) {
		var data entity.Image
		if err := cursor.Decode(&data); err != nil {
			return nil, err
		}
		images = append(images, &image.ImageResponse{
			ID:        data.ID,
			Url:       config.AppConfig.ServerURL + data.Url,
			Width:     data.Width,
			Height:    data.Height,
			CreatedAt: data.CreatedAt,
		})
	}
	return images, err
}

func (service *ImageService) SaveImage(saveFolder string, fileHeader *multipart.FileHeader) (string, multipart.File, error) {
	uploadPath := saveFolder + "/"
	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		os.Mkdir(uploadPath, os.ModePerm)
	}

	filename := fmt.Sprintf("%d%s", time.Now().Unix(), filepath.Ext(fileHeader.Filename))
	filepath := filepath.Join(uploadPath, filename)

	multipartFile, _ := fileHeader.Open()

	newfile, _ := os.Create(filepath)

	_, err := io.Copy(newfile, multipartFile)
	if err != nil {
		return "", nil, err
	}

	return filepath, multipartFile, nil
}
