package models

import "sync"

// FeatureItem representa un item individual dentro de un feature
type FeatureItem struct {
	Icon        string `json:"icon"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Feature representa la estructura completa de un feature
type Feature struct {
	ID       int           `json:"id"`
	Badge    string        `json:"badge"`
	Title    string        `json:"title"`
	Subtitle string        `json:"subtitle"`
	Image    string        `json:"image"`
	Features []FeatureItem `json:"features"`
}

// Almacenamiento en memoria para features
var (
	Features   = make(map[int]Feature)
	FeatureID  = 1
	FeatureMu  sync.Mutex
)

// InitFeatureData inicializa los datos de ejemplo para features
func InitFeatureData() {
	Features[1] = Feature{
		ID:       1,
		Badge:    "Deploy faster",
		Title:    "A better workflow",
		Subtitle: "Lorem ipsum, dolor sit amet consectetur adipisicing elit. Maiores impedit perferendis suscipit eaque, iste dolor cupiditate blanditiis ratione.",
		Image:    "https://tailwindcss.com/plus-assets/img/component-images/project-app-screenshot.png",
		Features: []FeatureItem{
			{
				Icon:        "M5.5 17a4.5 4.5 0 0 1-1.44-8.765 4.5 4.5 0 0 1 8.302-3.046 3.5 3.5 0 0 1 4.504 4.272A4 4 0 0 1 15 17H5.5Zm3.75-2.75a.75.75 0 0 0 1.5 0V9.66l1.95 2.1a.75.75 0 1 0 1.1-1.02l-3.25-3.5a.75.75 0 0 0-1.1 0l-3.25 3.5a.75.75 0 1 0 1.1 1.02l1.95-2.1v4.59Z",
				Title:       "Push to deploy.",
				Description: "Lorem ipsum, dolor sit amet consectetur adipisicing elit. Maiores impedit perferendis suscipit eaque, iste dolor cupiditate blanditiis ratione.",
			},
			{
				Icon:        "M10 1a4.5 4.5 0 0 0-4.5 4.5V9H5a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-6a2 2 0 0 0-2-2h-.5V5.5A4.5 4.5 0 0 0 10 1Zm3 8V5.5a3 3 0 1 0-6 0V9h6Z",
				Title:       "SSL certificates.",
				Description: "Anim aute id magna aliqua ad ad non deserunt sunt. Qui irure qui lorem cupidatat commodo.",
			},
			{
				Icon:        "M4.632 3.533A2 2 0 0 1 6.577 2h6.846a2 2 0 0 1 1.945 1.533l1.976 8.234A3.489 3.489 0 0 0 16 11.5H4c-.476 0-.93.095-1.344.267l1.976-8.234ZM4 13a2 2 0 1 0 0 4h12a2 2 0 1 0 0-4H4Zm11.24 2a.75.75 0 0 1 .75-.75H16a.75.75 0 0 1 .75.75v.01a.75.75 0 0 1-.75.75h-.01a.75.75 0 0 1-.75-.75V15Zm-2.25-.75a.75.75 0 0 0-.75.75v.01c0 .414.336.75.75.75H13a.75.75 0 0 0 .75-.75V15a.75.75 0 0 0-.75-.75h-.01Z",
				Title:       "Database backups.",
				Description: "Ac tincidunt sapien vehicula erat auctor pellentesque rhoncus. Et magna sit morbi lobortis.",
			},
		},
	}
	FeatureID = 2
}