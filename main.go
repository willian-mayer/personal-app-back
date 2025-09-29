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

// Almacenamiento en memoria
var (
	heroes = make(map[int]HeroSection)
	nextID = 1
	mu     sync.Mutex
)

func main() {
	// Agregar un hero section de ejemplo
	heroes[1] = HeroSection{
		ID:       1,
		Title:    "Bienvenido a Mi Sitio Web",
		Subtitle: "Creando experiencias increíbles",
		Desc:     "Desarrollamos soluciones innovadoras para hacer crecer tu negocio en el mundo digital.",
	}
	nextID = 2

	// Rutas de la API con middleware CORS
	http.HandleFunc("/hero-sections", corsMiddleware(heroSectionsHandler))
	http.HandleFunc("/hero-sections/", corsMiddleware(heroSectionHandler))

	log.Println("🚀 Servidor iniciado en http://localhost:8080")
	log.Println("✅ CORS habilitado para http://localhost:4200")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Middleware CORS para permitir peticiones desde Angular
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Permitir solicitudes desde tu app Angular
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		// Manejar solicitudes preflight (OPTIONS)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		// Continuar con el handler normal
		next(w, r)
	}
}

// Handler para /hero-sections - GET (listar) y POST (crear)
func heroSectionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		// Listar todos los hero sections
		mu.Lock()
		heroList := make([]HeroSection, 0, len(heroes))
		for _, hero := range heroes {
			heroList = append(heroList, hero)
		}
		mu.Unlock()

		json.NewEncoder(w).Encode(heroList)

	case "POST":
		// Crear nuevo hero section
		var hero HeroSection
		if err := json.NewDecoder(r.Body).Decode(&hero); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Validación básica
		if hero.Title == "" {
			http.Error(w, "El título es requerido", http.StatusBadRequest)
			return
		}

		mu.Lock()
		hero.ID = nextID
		nextID++
		heroes[hero.ID] = hero
		mu.Unlock()

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(hero)

	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

// Handler para /hero-sections/{id} - GET (obtener), PUT (actualizar), DELETE (eliminar)
func heroSectionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extraer ID de la URL
	idStr := r.URL.Path[len("/hero-sections/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "GET":
		// Obtener hero section específico
		mu.Lock()
		hero, exists := heroes[id]
		mu.Unlock()

		if !exists {
			http.Error(w, "Hero section no encontrado", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(hero)

	case "PUT":
		// Actualizar hero section
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

		// Validación básica
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
		// Eliminar hero section
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