/*
 * search.go - Modrinth search endpoint
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
	print(u.String())
	err = s.client.Get(u.String(), &searchResponse)
	if err != nil {
		return nil, err
	}
	return &searchResponse, nil
}
