package api

import "EatRecmd/emptySymbols"

func GetEatsApiPayloadThruGPT(query string) string {
	prompt := "I need to extract users' takeout preference from the query I'm giving you. " +
		"Return a JSON object and nothing else. The JSON could have the following fields: " +
		"food_type, restaurant_name(if specified),food_name. " +
		"User's query: " + query
	res := emptySymbols.CallGPT(prompt)
	return res
}
