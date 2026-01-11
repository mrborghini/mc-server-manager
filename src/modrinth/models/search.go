/*
 * search.go - Model for Modrinth search endpoint
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

package models

type SearchHit struct {
	ProjectId         string      `json:"project_id"`
	ProjectType       string      `json:"project_type"`
	Slug              string      `json:"slug"`
	Author            string      `json:"author"`
	Title             string      `json:"title"`
	Description       string      `json:"description"`
	Categories        []string    `json:"categories"`
	DisplayCategories []string    `json:"display_categories"`
	Versions          []string    `json:"versions"`
	Downloads         int64       `json:"downloads"`
	Follows           int64       `json:"follows"`
	IconURL           string      `json:"icon_url"`
	DateCreated       string      `json:"date_created"`
	DateModified      string      `json:"date_modified"`
	LatestVersion     string      `json:"latest_version"`
	License           string      `json:"license"`
	ClientSide        ProjectSide `json:"client_side"`
	ServerSide        ProjectSide `json:"server_side"`
	Gallery           []string    `json:"gallery"`
	FeaturedGallery   *string     `json:"featured_gallery"`
}

type SearchResponse struct {
	Hits []SearchHit `json:"hits"`
}
