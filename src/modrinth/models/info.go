package models

type APIInfo struct {
	About         string `json:"about"`
	Documentation string `json:"documentation"`
	Name          string `json:"name"`
	Version       string `json:"version"`
}
