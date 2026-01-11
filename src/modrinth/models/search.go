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
