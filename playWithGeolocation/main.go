package main

import (
	"log"
	"net/http"

	"github.com/elastic/go-elasticsearch"
	"github.com/gin-gonic/gin"
	handlers "github.com/venemarcos/go_study/playWithGeolocation/Handlers"
)

func main() {

	// Use default settings from gi and elasticsearch client
	router := gin.Default()
	es, _ := elasticsearch.NewDefaultClient()

	routeGroup := router.Group("/api")
	{
		routeGroup.GET("/", func(ctx *gin.Context) {
			handlers.HomeHandler(es, ctx)
		})

		routeGroup.GET("/search", func(ctx *gin.Context) {
			order := ctx.DefaultQuery("order", "asc")
			unit := ctx.DefaultQuery("unit", "km")
			limit := ctx.DefaultQuery("limit", "10")
			latLon := ctx.Query("latlon")

			if latLon == "" {
				ctx.JSON(
					http.StatusBadRequest,
					gin.H{
						"error": "You must specify latlon in query.",
					},
				)
				return
			}

			handlers.SearchHandler(order, unit, limit, latLon, es, ctx)
		})
	}

	err := router.Run()
	if err != nil {
		log.Fatalf("Error creating routes: %s", err)
	}
}
