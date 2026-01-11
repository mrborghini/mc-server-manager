/*
 * info.go - Modrinth root endpoint
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
