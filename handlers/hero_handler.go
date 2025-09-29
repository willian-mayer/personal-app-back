package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"personal-app-back/models"
)

// HeroSectionsHandler maneja las peticiones GET (listar) y POST (crear) para hero sections
func HeroSectionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		// Listar todos los hero sections
		models.HeroMu.Lock()
		heroList := make([]models.HeroSection, 0, len(models.Heroes))
		for _, hero := range models.Heroes {
			heroList = append(heroList, hero)
		}
		models.HeroMu.Unlock()

		json.NewEncoder(w).Encode(heroList)

	case "POST":
		// Crear nuevo hero section
		var hero models.HeroSection
		if err := json.NewDecoder(r.Body).Decode(&hero); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Validación básica
		if hero.Title == "" {
			http.Error(w, "El título es requerido", http.StatusBadRequest)
			return
		}

		models.HeroMu.Lock()
		hero.ID = models.HeroID
		models.HeroID++
		models.Heroes[hero.ID] = hero
		models.HeroMu.Unlock()

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(hero)

	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

// HeroSectionHandler maneja las peticiones GET (obtener), PUT (actualizar) y DELETE (eliminar) para un hero section específico
func HeroSectionHandler(w http.ResponseWriter, r *http.Request) {
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
		models.HeroMu.Lock()
		hero, exists := models.Heroes[id]
		models.HeroMu.Unlock()

		if !exists {
			http.Error(w, "Hero section no encontrado", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(hero)

	case "PUT":
		// Actualizar hero section
		models.HeroMu.Lock()
		_, exists := models.Heroes[id]
		models.HeroMu.Unlock()

		if !exists {
			http.Error(w, "Hero section no encontrado", http.StatusNotFound)
			return
		}

		var hero models.HeroSection
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
		models.HeroMu.Lock()
		models.Heroes[id] = hero
		models.HeroMu.Unlock()

		json.NewEncoder(w).Encode(hero)

	case "DELETE":
		// Eliminar hero section
		models.HeroMu.Lock()
		_, exists := models.Heroes[id]
		if exists {
			delete(models.Heroes, id)
		}
		models.HeroMu.Unlock()

		if !exists {
			http.Error(w, "Hero section no encontrado", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}