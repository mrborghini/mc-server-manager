package models

type ProjectSide string

const (
	Required    ProjectSide = "required"
	Optional    ProjectSide = "optional"
	Unsupported ProjectSide = "unsupported"
	Unknown     ProjectSide = "unknown"
)
