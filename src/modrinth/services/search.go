package services

import (
	"mrborghini/mc-server-manager/src/modrinth/client"
	"mrborghini/mc-server-manager/src/modrinth/models"
	"net/url"
)

type SearchService struct {
	client *client.Client
}

func NewSearchService(c *client.Client) *SearchService {
	return &SearchService{client: c}
}

func (s *SearchService) Get(query string) (*models.SearchResponse, error) {
	var u, err = url.Parse("/v2/search")
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Set("query", query)
	u.RawQuery = q.Encode()
	var searchResponse models.SearchResponse
	err = s.client.Get(u.String(), &searchResponse)
	if err != nil {
		return nil, err
	}
	return &searchResponse, nil
}
