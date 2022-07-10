package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"simple-rest-api/component"
	"simple-rest-api/component/uploadprovider"
	"simple-rest-api/middleware"
	"simple-rest-api/modules/restaurant/restauranttransportation/ginrestaurant"
	"simple-rest-api/modules/upload/uploadtransportation/ginupload"
)

func main() {
	dsn := os.Getenv("DBConnectionStr")

	s3BuckerName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3APIKey := os.Getenv("S3APIKey")
	s3Secret := os.Getenv("S3Secret")
	s3Domain := os.Getenv("S3Domain")

	uploadProvider := uploadprovider.NewS3Provider(
		s3BuckerName,
		s3Region,
		s3APIKey,
		s3Secret,
		s3Domain,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	if err := runService(db, uploadProvider); err != nil {
		log.Fatalln(err)
	}

}
func runService(db *gorm.DB, uploadProvider uploadprovider.UploadProvider) error {

	appCtx := component.NewAppContext(db, uploadProvider)
	r := gin.Default()
	r.Use(middleware.Recover(appCtx))

	r.POST("/upload", ginupload.Upload(appCtx))

	restaurants := r.Group("/restaurant")

	restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
	restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
	restaurants.GET("/:id", ginrestaurant.GetByIdRestaurant(appCtx))
	restaurants.PATCH("/:id", ginrestaurant.UpdateByIdRestaurant(appCtx))
	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurantById(appCtx))

	return r.Run()
}
