package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"simple-rest-api/component"
	"simple-rest-api/modules/restaurant/restauranttransportation/ginrestaurant"
)

func main() {
	dsn := os.Getenv("DBConnectionStr")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	if err := runService(db); err != nil {
		log.Fatalln(err)
	}

}
func runService(db *gorm.DB) error {
	r := gin.Default()
	appCtx := component.NewAppContext(db)
	restaurants := r.Group("/restaurant")
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
		restaurants.GET("/:id", ginrestaurant.GetByIdRestaurant(appCtx))
		restaurants.PATCH("/:id", ginrestaurant.UpdateByIdRestaurant(appCtx))
		restaurants.DELETE("/:id", ginrestaurant.DeleteRestarantById(appCtx))
	}

	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	return r.Run()
}
