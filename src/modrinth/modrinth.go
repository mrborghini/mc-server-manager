package modrinth

import (
	"mrborghini/mc-server-manager/src/modrinth/client"
	"mrborghini/mc-server-manager/src/modrinth/models"
	"mrborghini/mc-server-manager/src/modrinth/services"
)

const BASE_URL = "https://api.modrinth.com"

type Modrinth struct {
	client *client.Client
	info   *services.InfoService
	search *services.SearchService
}

func New() *Modrinth {
	m := new(Modrinth)
	m.client = client.New(BASE_URL)
	m.info = services.NewInfoService(m.client)
	m.search = services.NewSearchService(m.client)
	return m
}

func (m *Modrinth) GetApiInfo() (*models.APIInfo, error) {
	return m.info.Get()
}

func (m *Modrinth) SearchMod(query string) (*models.SearchResponse, error) {
	return m.search.Get(query)
}
