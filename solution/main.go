package main

import (
	"EatRecmd/api"
	"EatRecmd/emptySymbols"
	"github.com/gin-gonic/gin"
)

func main() {
	s := gin.Default()

	s.GET("/recommendations", getRecmd)
}

// / getRecmd is the handler func to process user query.
func getRecmd(c *gin.Context) {
	query := c.Query("query")
	userLocation := c.Query("location")
	// First check if similar queries already exist.
	if emptySymbols.QueryCacheEnabled {
		old_q, old_location := emptySymbols.GetMostSimilarExistingQuery(query)
		if emptySymbols.NearEnough(userLocation, old_location) {
			if emptySymbols.PreviousQueryHasntExpired(old_q) {
				c.JSON(200, gin.H{
					"recommendation": emptySymbols.CallGPT(old_q),
				})
				return
			}
		}
	}
	reqBody := api.GetEatsApiPayloadThruGPT(query)
	emptySymbols.EmbedQuery(query, userLocation)
	// SaveQuery to DB(query) with faiss
	res := api.GetEats(reqBody)
	// save res(recommend results) to Mongo
	c.JSON(200, gin.H{
		"result": res,
	})

}
