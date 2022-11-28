package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple_golang/common"
	"simple_golang/component"
	"simple_golang/modules/restaurant/restaurantbiz"
	"simple_golang/modules/restaurant/restaurantmodel"
	"simple_golang/modules/restaurant/restaurantstorage"
)

func CreateRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := restaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())

		biz := restaurantbiz.NewCreateRestaurantBiz(store)
		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
