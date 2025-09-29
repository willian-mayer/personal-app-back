package models

import "sync"

// HeroSection representa la estructura de un hero section
type HeroSection struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Desc     string `json:"desc"`
}

// Almacenamiento en memoria para heroes
var (
	Heroes   = make(map[int]HeroSection)
	HeroID   = 1
	HeroMu   sync.Mutex
)

// InitHeroData inicializa los datos de ejemplo para heroes
func InitHeroData() {
	Heroes[1] = HeroSection{
		ID:       1,
		Title:    "Bienvenido a Mi Sitio Web",
		Subtitle: "Creando experiencias increíbles",
		Desc:     "Desarrollamos soluciones innovadoras para hacer crecer tu negocio en el mundo digital.",
	}
	HeroID = 2
}