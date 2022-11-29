package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple_golang/common"
	"simple_golang/component"
	"simple_golang/modules/restaurant/restaurantbiz"
	"simple_golang/modules/restaurant/restaurantmodel"
	"simple_golang/modules/restaurant/restaurantstorage"
	"strconv"
)

func UpdateRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		var data restaurantmodel.RestaurantUpdate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(401, gin.H{
				"err": err.Error(),
			})
		}

		store := restaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewUpdateRestaurantBiz(store)

		if err := biz.UpdateRestaurant(c.Request.Context(), id, &data); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}

}
