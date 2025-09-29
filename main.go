package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"
)

// Estructura del Hero Section
type HeroSection struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Desc     string `json:"desc"`
}

// Estructuras para Features
type FeatureItem struct {
	Icon        string `json:"icon"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Feature struct {
	ID       int           `json:"id"`
	Badge    string        `json:"badge"`
	Title    string        `json:"title"`
	Subtitle string        `json:"subtitle"`
	Image    string        `json:"image"`
	Features []FeatureItem `json:"features"`
}

// Almacenamiento en memoria
var (
	heroes   = make(map[int]HeroSection)
	features = make(map[int]Feature)
	heroID   = 1
	featureID = 1
	mu       sync.Mutex
)

func main() {
	// Datos de ejemplo para Hero Section
	heroes[1] = HeroSection{
		ID:       1,
		Title:    "Bienvenido a Mi Sitio Web",
		Subtitle: "Creando experiencias increíbles",
		Desc:     "Desarrollamos soluciones innovadoras para hacer crecer tu negocio en el mundo digital.",
	}
	heroID = 2

	// Datos de ejemplo para Features
	features[1] = Feature{
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
	featureID = 2

	// Rutas de la API con middleware CORS
	http.HandleFunc("/hero-sections", corsMiddleware(heroSectionsHandler))
	http.HandleFunc("/hero-sections/", corsMiddleware(heroSectionHandler))
	http.HandleFunc("/features", corsMiddleware(featuresHandler))
	http.HandleFunc("/features/", corsMiddleware(featureHandler))

	log.Println("🚀 Servidor iniciado en http://localhost:8080")
	log.Println("✅ CORS habilitado")
	log.Println("📍 Endpoints disponibles:")
	log.Println("   - /hero-sections")
	log.Println("   - /features")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Middleware CORS
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allowedOrigins := map[string]bool{
			"http://localhost:4200": true,
			"http://localhost:80":   true,
			"http://localhost":      true,
		}
		
		origin := r.Header.Get("Origin")
		if allowedOrigins[origin] {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}
		
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next(w, r)
	}
}

// ========== HANDLERS PARA HERO SECTIONS ==========

func heroSectionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		mu.Lock()
		heroList := make([]HeroSection, 0, len(heroes))
		for _, hero := range heroes {
			heroList = append(heroList, hero)
		}
		mu.Unlock()
		json.NewEncoder(w).Encode(heroList)

	case "POST":
		var hero HeroSection
		if err := json.NewDecoder(r.Body).Decode(&hero); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if hero.Title == "" {
			http.Error(w, "El título es requerido", http.StatusBadRequest)
			return
		}

		mu.Lock()
		hero.ID = heroID
		heroID++
		heroes[hero.ID] = hero
		mu.Unlock()

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(hero)

	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func heroSectionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := r.URL.Path[len("/hero-sections/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "GET":
		mu.Lock()
		hero, exists := heroes[id]
		mu.Unlock()

		if !exists {
			http.Error(w, "Hero section no encontrado", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(hero)

	case "PUT":
		mu.Lock()
		_, exists := heroes[id]
		mu.Unlock()

		if !exists {
			http.Error(w, "Hero section no encontrado", http.StatusNotFound)
			return
		}

		var hero HeroSection
		if err := json.NewDecoder(r.Body).Decode(&hero); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if hero.Title == "" {
			http.Error(w, "El título es requerido", http.StatusBadRequest)
			return
		}

		hero.ID = id
		mu.Lock()
		heroes[id] = hero
		mu.Unlock()

		json.NewEncoder(w).Encode(hero)

	case "DELETE":
		mu.Lock()
		_, exists := heroes[id]
		if exists {
			delete(heroes, id)
		}
		mu.Unlock()

		if !exists {
			http.Error(w, "Hero section no encontrado", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

// ========== HANDLERS PARA FEATURES ==========

func featuresHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		mu.Lock()
		featureList := make([]Feature, 0, len(features))
		for _, feature := range features {
			featureList = append(featureList, feature)
		}
		mu.Unlock()
		json.NewEncoder(w).Encode(featureList)

	case "POST":
		var feature Feature
		if err := json.NewDecoder(r.Body).Decode(&feature); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if feature.Title == "" {
			http.Error(w, "El título es requerido", http.StatusBadRequest)
			return
		}

		mu.Lock()
		feature.ID = featureID
		featureID++
		features[feature.ID] = feature
		mu.Unlock()

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(feature)

	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func featureHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := r.URL.Path[len("/features/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "GET":
		mu.Lock()
		feature, exists := features[id]
		mu.Unlock()

		if !exists {
			http.Error(w, "Feature no encontrado", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(feature)

	case "PUT":
		mu.Lock()
		_, exists := features[id]
		mu.Unlock()

		if !exists {
			http.Error(w, "Feature no encontrado", http.StatusNotFound)
			return
		}

		var feature Feature
		if err := json.NewDecoder(r.Body).Decode(&feature); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if feature.Title == "" {
			http.Error(w, "El título es requerido", http.StatusBadRequest)
			return
		}

		feature.ID = id
		mu.Lock()
		features[id] = feature
		mu.Unlock()

		json.NewEncoder(w).Encode(feature)

	case "DELETE":
		mu.Lock()
		_, exists := features[id]
		if exists {
			delete(features, id)
		}
		mu.Unlock()

		if !exists {
			http.Error(w, "Feature no encontrado", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}