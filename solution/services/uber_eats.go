package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"main/models"
)

func (s *UberEatsService) GetRecommendationsFromUberEats(food string) ([]models.Recommendation, error) {
	url := "https://www.ubereats.com/_p/api/getSearchFeedV1?localeCode=ca"
	payload := map[string]interface{}{
		"userQuery":      food,
		"date":           "",
		"startTime":      0,
		"endTime":        0,
		"sortAndFilters": []interface{}{},
		"vertical":       "ALL",
		"searchSource":   "SEARCH_SUGGESTION",
		"displayType":    "SEARCH_RESULTS",
		"searchType":     "GLOBAL_SEARCH",
		"keyName":        "",
		"cacheKey":       "",
		"recaptchaToken": "",
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Csrf-Token", "x")

	httpResp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer httpResp.Body.Close()

	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var response models.GetFeedResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	var resultList []map[string]string
	for i, item := range response.Data.FeedItems {
		if i >= 3 {
			break
		}
		result := map[string]string{
			"storeUuid": item.Store.StoreUuid,
			"title":     item.Store.Title.Text,
		}
		resultList = append(resultList, result)
	}

	var recommendations []models.Recommendation
	for _, item := range resultList {
		recommendation := models.Recommendation{
			RecommendationID: item["storeUuid"],
			Store: models.RecommendStore{
				Name:     item["title"],
				ID:       item["storeUuid"],
				MenuItem: []models.MenuItem{},
			},
		}
		recommendations = append(recommendations, recommendation)
	}

	for i, recommendation := range recommendations {
		menuItems, err := s.getMenuItemFromStore(recommendation.Store.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get menu items from store: %w", err)
		}
		recommendations[i].Store.MenuItem = menuItems
	}

	return recommendations, nil
}

func (s *UberEatsService) getMenuItemFromStore(storeID string) ([]models.MenuItem, error) {
	url := "https://www.ubereats.com/_p/api/getStoreV1?localeCode=ca"
	payload := map[string]interface{}{
		"storeUuid":  storeID,
		"diningMode": "DELIVERY",
		"time": map[string]bool{
			"asap": true,
		},
		"cbType": "EATER_ENDORSED",
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Csrf-Token", "x")

	httpResp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer httpResp.Body.Close()

	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var res models.GetStoreResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	sections := res.Data.CatalogSectionsMap
	var menuItems []models.MenuItem

	for _, section := range sections {
		for _, item := range section {
			for i, dish := range item.Payload.StandardItemsPayload.CatalogItems {
				if i >= 2 {
					break
				}
				menuItem := models.MenuItem{
					ID:    dish.UUID,
					Name:  dish.Title,
					Price: dish.PriceTagline.Text,
				}
				menuItems = append(menuItems, menuItem)
			}
		}
		break
	}

	return menuItems, nil
}
