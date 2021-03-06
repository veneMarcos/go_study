package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/elastic/go-elasticsearch"
	"github.com/gin-gonic/gin"
)

// HomeHandler is ...
func HomeHandler(es *elasticsearch.Client, context *gin.Context) {
	var r map[string]interface{}

	result, err := es.Info()
	if err != nil {
		log.Fatalf("Error when getting elasticsearch info. %s", err)
	}

	if err := json.NewDecoder(result.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	headerStatus := http.StatusOK
	body := gin.H{
		"elasticsearch_version": r["version"].(map[string]interface{})["number"],
	}
	context.JSON(
		headerStatus,
		body,
	)
}
