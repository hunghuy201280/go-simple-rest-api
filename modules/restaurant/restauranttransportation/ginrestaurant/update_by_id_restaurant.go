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

func UpdateByIdRestaurant(ctx component.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		var data restaurantmodel.RestaurantUpdate
		var id restaurantmodel.RestaurantId
		if err := context.ShouldBindUri(&id); err != nil {
			panic(err)

		}

		err := context.ShouldBind(&data)
		if err != nil {
			panic(err)

		}

		store := restaurantstorage.NewSQLStore(ctx.GetMainDbConnection())
		biz := restaurantbiz.NewUpdateRestaurantByIdBiz(store)
		if err := biz.UpdateByIdRestaurant(context.Request.Context(), &id, &data); err != nil {
			panic(err)

		}

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
		return
	}
}
