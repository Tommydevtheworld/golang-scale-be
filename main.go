package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"simple_golang/component"
	"simple_golang/modules/restaurant/restaurantmodel"
	"simple_golang/modules/restaurant/restauranttransport/ginrestaurant"
)

//type Note struct {
//	Id      uint   `json:"id,omitempty" gorm:"column:id"`
//	Title   string `json:"title" gorm:"column:title"`
//	Content string `json:"content" gorm:"column:content"`
//}

//func (Note) TableName() string {
//	return "notes"
//}

func main() {

	dsn := os.Getenv("DBConnectionStr")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	if err := db.AutoMigrate(&restaurantmodel.Restaurant{}); err != nil {
		log.Fatalln(err)
	}

	if err := runService(db); err != nil {
		log.Fatalln(err)
	}

}

func runService(db *gorm.DB) error {
	r := gin.Default()
	gin.ForceConsoleColor()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	appCtx := component.NewAppContext(db)
	restaurants := r.Group("/restaurants")
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
		restaurants.GET("/:id", ginrestaurant.DetailRestaurant(appCtx))
	}

	return r.Run(":8080")
}
