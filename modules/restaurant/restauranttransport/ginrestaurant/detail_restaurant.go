package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple_golang/common"
	"simple_golang/component"
	"simple_golang/modules/restaurant/restaurantbiz"
	"simple_golang/modules/restaurant/restaurantstorage"
)

func DetailRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		store := restaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())

		biz := restaurantbiz.NewDetailRestaurantBiz(store)
		result, err := biz.DetailRestaurant(c.Request.Context(), id)
		if err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
