package emptySymbols

const (
	QueryCacheEnabled = true
)

func CallGPT(p string) string {
	return ""
}

// / NearEnough check if two locations are near enough to share a recommendation
func NearEnough(loc1, loc2 string) bool {
	return true
}

// / EmbedQuery embed the query into a vector and store in vec db.
func EmbedQuery(query string, location string) {
	// embed it as vectors
	// store to Faiss, e.g..
	return
}

// / GetMostSimilarExistingQuery get the most similar existing query from faiss to the new query.
func GetMostSimilarExistingQuery(newquery string) (old_q string, old_location string) {
	return "", ""
}

func PreviousQueryHasntExpired(old_query_id string) bool {
	return true
}
