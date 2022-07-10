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
			panic(common.ErrInvalidRequest(err))

		}

		store := restaurantstorage.NewSQLStore(ctx.GetMainDbConnection())
		biz := restaurantbiz.NewGetByIdRestaurantBiz(store)
		result, err := biz.GetByIdRestaurant(context.Request.Context(), &data)
		if err != nil {

			panic(err)
		}

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(result))

	}
}
