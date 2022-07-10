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

func DeleteRestaurantById(ctx component.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		var id restaurantmodel.RestaurantId
		err := context.ShouldBindUri(&id)
		if err != nil {
			panic(err)

		}

		store := restaurantstorage.NewSQLStore(ctx.GetMainDbConnection())
		biz := restaurantbiz.NewDeleteRestaurantBiz(store)
		if err := biz.DeleteRestaurantById(context.Request.Context(), &id); err != nil {
			panic(err)

		}

		context.JSON(http.StatusOK, common.SimpleSuccessResponse("ok"))
	}
}
