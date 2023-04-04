package models

type SpringProject struct {
	ArtifactId  string `json:"artifactId" binding:"required"`
	BootVersion string `json:"bootVersion" binding:"required"`
	Description string `json:"description"`
	GroupId     string `json:"groupId" binding:"required"`
	JavaVersion string `json:"javaVersion" binding:"required"`
	Language    string `json:"language" binding:"required"`
	Name        string `json:"name" binding:"required"`
	PackageName string `json:"packageName" binding:"required"`
	Packaging   string `json:"packaging" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Version     string `json:"version"`
}
