package models

// UberEat GetFeed return type

type GetFeedResponse struct {
	Data GetFeedData `json:"data"`
}

type GetFeedData struct {
	FeedItems []GetFeedFeedItem `json:"feedItems"`
}

type GetFeedFeedItem struct {
	Store GetFeedStore `json:"store"`
}

type GetFeedStore struct {
	StoreUuid string       `json:"storeUuid"`
	Title     GetFeedTitle `json:"title"`
}

type GetFeedTitle struct {
	Text string `json:"text"`
}

// UberEat GetStore return type
type GetStoreResponse struct {
	Data GetStoreData `json:"data"`
}

type GetStoreData struct {
	CatalogSectionsMap map[string][]GetStoreCatalogSection `json:"catalogSectionsMap"`
}

type GetStoreCatalogSection struct {
	Payload GetStorePayload `json:"payload"`
}

type GetStorePayload struct {
	StandardItemsPayload GetStoreStandardItemsPayload `json:"standardItemsPayload"`
}

type GetStoreStandardItemsPayload struct {
	CatalogItems []GetStoreCatalogItem `json:"catalogItems"`
}

type GetStoreCatalogItem struct {
	UUID            string               `json:"uuid"`
	Title           string               `json:"title"`
	ItemDescription string               `json:"itemDescription"`
	PriceTagline    GetStorePriceTagline `json:"priceTagline"`
}

type GetStorePriceTagline struct {
	Text string `json:"text"`
}

// recommendations api return type

type MenuItem struct {
	Name  string `json:"name"`
	ID    string `json:"id"`
	Price string `json:"price"`
}

type RecommendStore struct {
	Name     string     `json:"name"`
	ID       string     `json:"id"`
	MenuItem []MenuItem `json:"menuItem"`
}

type Recommendation struct {
	RecommendationID string         `json:"recommendationId"`
	Store            RecommendStore `json:"store"`
}
