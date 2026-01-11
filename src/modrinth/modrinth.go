/*
 * modrinth.go - High level modrinth API
 *
 * Copyright (C) 2026 Mrborghini
 *
 * This program is free software; you can redistribute it and/or
 * modify it under the terms of the GNU General Public License
 * as published by the Free Software Foundation; either version 2
 * of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, see <https://www.gnu.org/licenses/>.
 */
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
