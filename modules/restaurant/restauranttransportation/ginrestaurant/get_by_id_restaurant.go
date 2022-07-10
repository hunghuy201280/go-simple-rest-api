package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-rest-api/common"
	"simple-rest-api/component"
	"simple-rest-api/modules/restaurant/restaurantbiz"
	"simple-rest-api/modules/restaurant/restaurantmodel"
	"simple-rest-api/modules/restaurant/restaurantstorage"
)

func GetByIdRestaurant(ctx component.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		var data restaurantmodel.RestaurantId

		if err := context.ShouldBindUri(&data); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return

		}

		store := restaurantstorage.NewSQLStore(ctx.GetMainDbConnection())
		biz := restaurantbiz.NewGetByIdRestaurantBiz(store)
		result, err := biz.GetByIdRestaurant(context, &data)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(result))

	}
}
