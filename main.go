package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"simple-rest-api/component"
	"simple-rest-api/middleware"
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
	appCtx := component.NewAppContext(db)
	r := gin.Default()
	r.Use(middleware.Recover(appCtx))

	restaurants := r.Group("/restaurant")

	restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
	restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
	restaurants.GET("/:id", ginrestaurant.GetByIdRestaurant(appCtx))
	restaurants.PATCH("/:id", ginrestaurant.UpdateByIdRestaurant(appCtx))
	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurantById(appCtx))

	return r.Run()
}
