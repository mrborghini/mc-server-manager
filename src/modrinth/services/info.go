package services

import (
	"mrborghini/mc-server-manager/src/modrinth/client"
	"mrborghini/mc-server-manager/src/modrinth/models"
)

type InfoService struct {
	client *client.Client
}

func NewInfoService(c *client.Client) *InfoService {
	return &InfoService{client: c}
}

func (s *InfoService) Get() (*models.APIInfo, error) {
	var info models.APIInfo
	err := s.client.Get("/", &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}
